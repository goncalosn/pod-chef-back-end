# pod-chef-back-end

API to serve the front-end side

## Build

How to build a docker image for demo purposes

```bash
sudo docker build -f cmd/demo/Dockerfile -t podchef/app:demo .
docker push podchef/app:demo
```

Service image

```bash
sudo docker build -t podchef/production:latest .
docker push podchef/production:latest
```

## Run in production

```
sudo docker run -d --net=host --volume=$HOME/.kube/config:/root/.kube/config  podchef/production
```
