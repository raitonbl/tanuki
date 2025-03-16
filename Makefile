# Default target when no argument is provided
.PHONY: run run-tls

run:
	cd backend && go run cmd/main.go serve

run-tls:
	cd backend && go run cmd/main.go serve --server.registry.tls.key=key.pem --server.registry.tls.cert=cert.pem