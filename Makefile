run: build
	@./bin/rio

build: clean
	go build -o ./bin/rio

clean: 
	@rm -rf @./bin/rio