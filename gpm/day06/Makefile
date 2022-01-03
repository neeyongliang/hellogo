all:
	rm -rf generate
	go generate && mkdir generate && mv *_container.go generate
	go run day06.go