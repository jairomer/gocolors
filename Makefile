build:
	go build -o gocolors main.go

run:
	go run main.go

clean:
	rm gocolors

install:
	go build -o /bin/gocolors main.go
