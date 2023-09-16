clean:
	./clean.sh

docker-central:
ifeq ($(HOST),localhost)
	docker run -d -it --name central -p 50052:50052 --expose 50052 lab1:latest go run Central/main.go
else
	echo "Ejecutar SOLO en dist105"
endif

HOST = $(shell hostname)
docker-regional:
ifeq ($(HOST),dist105)
	docker run -d -it --name regional -p 50052:50052 --expose 50052 lab1:latest go run Regionales/America/regional.go
endif
ifeq ($(HOST),dist106)
	docker run -it --name regional -p 50053:50053 --expose 50053 lab1:latest go run Regionales/Asia/regional.go
endif
ifeq ($(HOST),dist107)
	docker run -it --name regional -p 50054:50054 --expose 50054 lab1:latest go run Regionales/Europa/regional.go
endif
ifeq ($(HOST),dist108)
	docker run  -it --name regional -p 50055:50055 --expose 50055 lab1:latest go run Regionales/Oceania/regional.go
else
	echo "Adios"
endif

docker-rabbit:
ifeq ($(HOST),dist106)
	sudo docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
else
	echo "Ejecutar SOLO en dist106"
endif