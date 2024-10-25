#!/bin/bash
path=`pwd`
version=$1

i=$(docker images | grep "basego-server" | grep "$version" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile -t basego-server:$version ../basego
fi

i=$(docker images | grep "basego-server" | grep "$version" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "basego-mysql-$version" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."

docker run -d \
    -p 3306:3306 \
    -v $path/conf/my.cnf:/etc/mysql/mysql.conf.d/my.cnf \
    -v $path/../basego-data:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=basego \
    --name basego-mysql-$version \
    --restart always \
    mysql:8.0
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 basego-mysql-$version
fi

i=$(docker ps -a | grep "basego-server:$version" | awk '{print $1}')
if test ! -z $i; then
echo "the server container already exists, delete..."
docker rm -f basego-server-$version
fi

echo "start the server..."
docker run -d \
    -p 9606:9606 \
    -w /basego \
    -v $path/conf:/basego/conf \
    -v $path/log:/basego/log \
    -v $path/tmp:/basego/tmp \
    -e TZ=Asia/Shanghai \
    -m 1024M \
    --net=host \
    --memory-swap 2048M \
    --cpus 2 \
    --name basego-server-$version \
    --restart always \
    --privileged \
    basego-server:$version \
    bash -c "cd src && ./basego -config ../conf/config.yaml"
sleep 2s
docker logs basego-server-$version
echo "the server has been started!"

