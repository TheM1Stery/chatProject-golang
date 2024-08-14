run:
	@templ generate
	@bun build client/index.ts --outdir=public --splitting
	@bun tailwind -i client/main.css -o public/main.css
	@go run .


run-browser:
	@templ generate
	@bun build client/index.ts --outdir=public --splitting
	@bun tailwind -i client/main.css -o public/main.css
	xdg-open http://localhost:9000
	@go run .
