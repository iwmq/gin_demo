dev:
	DEVELOP=true go run httpd/main.go
prod:
	DEVELOP=false go run httpd/main.go
collect:
	cd ../gin_frontend && npm run build && cd ../gin_backend && cp -r ../gin_frontend/dist/assets/* assets/
