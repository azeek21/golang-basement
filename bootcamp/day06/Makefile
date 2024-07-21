build:
	find cmd/*.go | xargs -n 1 go build -o ./bin/

generate_views:
	templ generate

generate: generate_views build
	./bin/generate_logo

migrate-debug: build
	GIN_MODE=debug ./bin/migrate


migrate-release: build
	GIN_MODE=release ./bin/migrate

dev: generate 
	./bin/start

start: generate
	GIN_MODE=release ./bin/start

clean:
	rm -rf ./bin/
	rm -f ./public/logo.png
	find ./views -name '*_templ.go' -delete

.PHONY: build start dev dev_prepare clean generate_views migrate-debug migrate-release
