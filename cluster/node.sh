if [ "$#" -ne 1 ]; then
  echo "Usage: $0 id" >&2
  exit 1
fi
id=$1

docker run --rm --name=nats$id --net=natsnet nats:0.9.4 -DV \
  -routes nats://ruser:T0pS3cr3t@nats:6222
