test:
	sudo docker-compose up -d
	go build -o ./builds/pub -v ./cmd/pub
	go build -o ./builds/app -v ./cmd/app
	./builds/pub& ./builds/app

stop:
	pgrep pub | xargs kill
	pgrep nats-streaming | xargs kill
