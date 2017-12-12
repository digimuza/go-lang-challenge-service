.SILENT:
help:
	echo "go-lang-challenge-service"
	echo "	Comands:"
	echo "		help:"
	echo "		watch:"

watch:
	sudo docker-compose -f docker/docker-compose.yml pull
	sudo docker-compose -f docker/docker-compose.yml build
	sudo docker-compose -f docker/docker-compose.yml up