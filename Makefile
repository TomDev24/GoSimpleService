test:
	sudo docker-compose up -d
	go build -v ./cmd/pub
	go build -v ./cmd/app
	nohup nats-streaming-server& sleep 1; ./pub& ./app

stop:
	pgrep pub | xargs kill
	pgrep nats-streaming | xargs kill
