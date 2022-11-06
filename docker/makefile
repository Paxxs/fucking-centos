VER = docker.8 docker.7
docker: $(VER)
docker.7: 
	echo ":: build Centos 7"
	docker build --rm -t local/centos:7-ssh 7/ssh/
	docker build --rm -t local/centos:7-ssh-go-1.17 7/go/

docker.8:
	echo ":: build Centos 8"
	docker build --rm -t local/centos:8-ssh 8/ssh/
	docker build --rm -t local/centos:8-ssh-go-1.17 8/go/