package registry

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raitonbl/tanuki/internal/context"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func setProvidersRoute(ctx context.Context, apiGroup *gin.RouterGroup) {
	api := apiGroup.Group("providers")
	api.GET("/:namespace/:type/:version", createGetProviderMetadata(ctx))
	api.GET("/:namespace/:type/versions", createGetProviderVersionsHandler(ctx))
	api.GET("/:namespace/:type/:version/download/:os/:arch", createGetProviderDownloadURL(ctx))
}

func createGetProviderVersionsHandler(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		redirectTo := fmt.Sprintf(
			"/v1/providers/%s/%s/versions",
			c.Param("namespace"), c.Param("type"),
		)
		response, err := sendHttpRequestToRegistry(ctx, redirectTo)
		if err != nil {
			ctx.Logger().Error("Error getting provider versions", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		if err = redirectResponse(c, response); err != nil {
			ctx.Logger().Error("Error getting provider versions", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		return
	}
}

func createGetProviderMetadata(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		redirectTo := fmt.Sprintf(
			"/v1/providers/%s/%s/%s",
			c.Param("namespace"), c.Param("type"), c.Param("version"),
		)
		response, err := sendHttpRequestToRegistry(ctx, redirectTo)
		if err != nil {
			ctx.Logger().Error("Error getting provider metadata", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		if err = redirectResponse(c, response); err != nil {
			ctx.Logger().Error("Error getting provider metadata", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		return
	}

}

func createGetProviderDownloadURL(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		redirectTo := fmt.Sprintf(
			"/v1/providers/%s/%s/%s/download/%s/%s",
			c.Param("namespace"), c.Param("type"), c.Param("version"),
			c.Param("os"), c.Param("arch"),
		)
		response, err := sendHttpRequestToRegistry(ctx, redirectTo)
		if err != nil {
			ctx.Logger().Error("Error getting provider metadata", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		if err = redirectResponse(c, response); err != nil {
			ctx.Logger().Error("Error getting provider metadata", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		}
		return

	}
}

func redirectResponse(c *gin.Context, response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	for key, values := range response.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}
	c.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
	return nil
}

func sendHttpRequestToRegistry(ctx context.Context, uri string) (*http.Response, error) {
	url := ctx.Configuration().Targets[0]
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	url += uri
	return http.Get(url) //TODO: USE LOAD-BALANCING & HEALTH CHECKS INSTEAD OF HARDCODING
}
