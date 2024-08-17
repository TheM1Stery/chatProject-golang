run:
	@templ generate
	@sh ./scripts/client-build.sh
	@go run .


run-browser:
	@templ generate
	@sh ./scripts/client-build.sh
	xdg-open http://localhost:9000
	@go run .
