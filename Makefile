run:
	@templ generate
	@bun build client/index.ts --outdir=public
	@bun tailwind -i client/main.css -o public/main.css
	xdg-open http://localhost:9000
	@go run .

