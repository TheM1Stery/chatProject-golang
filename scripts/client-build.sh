#!/bin/sh


bun build client/*.ts --splitting  --outdir=public --define process.env.BACKEND_URL="'localhost:9000'" --minify
bun tailwind -i client/main.css -o public/main.css

