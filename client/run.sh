if [ "$#" -ne 1 ]; then
  echo "Usage: $0 app" >&2
  exit 1
fi
app=$1

echo "WARNING: This produces non-repeatable builds - use only for demos!"

cd $app

cat > Dockerfile <<EOF
FROM golang:1.7.3-onbuild
ENTRYPOINT ["go-wrapper", "run"]
EOF

docker build -t $app .
rm -f Dockerfile

docker run --rm --net=natsnet $app nats://nats:4222

cd -
