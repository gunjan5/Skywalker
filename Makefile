build:
	sudo docker build -t skywalker .

run:
	sudo docker run -it --privileged --name sky1 -d skywalker

cli:
	sudo docker exec -it sky1 sh