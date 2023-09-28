#!/bin/bash
path=`pwd`


i=$(docker images | grep "basego-server" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile -t basego-server:v1.0.0 ../basego
fi

i=$(docker images | grep "basego-server" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "basego-mysql" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."

docker run -d \
    -p 33096:3306 \
    -v $path/conf/my.cnf:/etc/mysql/mysql.conf.d/my.cnf \
    -v $path/../basego-data:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=basego \
    --name basego-mysql \
    --restart always \
    mysql:5.7
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 basego-mysql
fi

i=$(docker ps -a | grep "basego-server" | awk '{print $1}')
if test ! -z $i; then
echo "the server container already exists, delete..."
docker rm -f basego-server
fi

echo "start the server..."
docker run -d \
    -p 9606:9606 \
    -w /basego \
    -v $path/conf:/basego/conf \
    -v $path/log:/basego/log \
    -v $path/tmp:/basego/tmp \
    -m 1024M \
    --net=host \
    --memory-swap 2048M \
    --cpus 2 \
    --name basego-server \
    --restart always \
    --privileged \
    basego-server:v1.0.0 \
    bash -c "cd src && ./basego -config ../conf/config.yaml"
sleep 2s
docker logs basego-server
echo "the server has been started!"

