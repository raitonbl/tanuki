
run:
	cd backend && go run cmd/main.go serve

run-tls:
	export TANUKI_SERVER_SERVERS_REGISTRY_TLS_KEY=key.pem && \
	 export TANUKI_SERVER_SERVERS_REGISTRY_TLS_CERT=cert.pem && \
	 cd backend && go run cmd/main.go serve