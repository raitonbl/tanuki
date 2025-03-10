import React, {createContext, useContext, useEffect, useState} from "react";
import {User, UserManager, WebStorageStateStore} from "oidc-client-ts";

const openIdConnectProviderConfiguration = {
    authority: "http://localhost:8080/realms/tanuki", // Replace with your OpenID Provider
    client_id: "test",
    redirect_uri: window.location.origin ,
    response_type: "code",
    scope: "openid profile email",
    post_logout_redirect_uri: window.location.origin,
    userStore: new WebStorageStateStore({store: window.localStorage}),
};

const userManager = new UserManager(openIdConnectProviderConfiguration);

const AuthContext = createContext<{ securityPrincipal: User | null; login: () => void; logout: () => void } | undefined>(undefined);

export const OpenIdConnectProvider = ({children}: { children: React.ReactNode }) => {
    const [user, setUser] = useState<User | null>(null);
    useEffect(() => {
        userManager.getUser().then((loadedUser) => {
            if (loadedUser && !loadedUser.expired) {
                setUser(loadedUser);
            }
        });

        userManager.events.addUserLoaded((newUser) => {
            setUser(newUser);
        });

        userManager.events.addUserSignedOut(() => {
            setUser(null);
        });
    }, []);

    const login = () => userManager.signinRedirect();
    const logout = () => userManager.signoutRedirect();

    // @ts-ignore
    return <AuthContext.Provider value={{user, login, logout}}>{children}</AuthContext.Provider>;
};

export const useOpenIdConnect = () => {
    const context = useContext(AuthContext);
    if (!context) throw new Error("useAuth must be used within an AuthProvider");
    return context;
};
