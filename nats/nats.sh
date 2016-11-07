docker network rm natsnet
docker network create natsnet
docker run --rm --name=nats --net=natsnet -p 4222:4222 nats:0.9.4 -DV
