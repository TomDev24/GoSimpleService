pub:
	go build -o ./builds/pub -v ./cmd/pub
	./builds/pub

server:
	go build -o ./builds/pub -v ./cmd/pub
	./builds/app

up:
	sudo docker-compose up -d

test:
	sudo docker-compose up -d
	go build -o ./builds/pub -v ./cmd/pub
	go build -o ./builds/app -v ./cmd/app
	sleep 3; ./builds/pub& ./builds/app

stop:
	sudo docker-compose down
	pgrep pub | xargs kill
	ps aux  |  grep -i nats-streaming  |  awk '{print $2}'  |  xargs sudo kill -9
