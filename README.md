# pod-chef-back-end

API to serve the front-end side

## .env example

How to build a docker image for demo purposes

```bash
DB_USER="<username here>"
DB_PASSWORD="<password here>"
TOKEN_SECRET="<secret here>"
CRYPT_KEY="<secret2 here>"
EMAIL_HOST="<email protocol here>"
EMAIL_PORT="<email port here>"
EMAIL_FROM="<email sender here>"
EMAIL_PASSWORD="<email password>"
TOKEN_CLOUDFLARE="<cloudflare api token here>"
ZONE_ID_CLOUDFLARE="<cloudflare zone id here>"
ZONE_IP_CLOUDFLARE="<cloudflare zone ip here>"

```

## Build demo

How to build a docker image for demo purposes

```bash
sudo docker build -f cmd/demo/Dockerfile -t podchef/app:demo .
docker push podchef/app:demo
```

## Build production image

```bash
sudo docker build -t podchef/backend:latest .
docker push podchef/backend:latest
```

## Run in production

```
sudo docker run -d --name podchef --net=host podchef/backend
```

## Run in dev

```
sudo docker run -d --name podchef --net=host --volume=$HOME/.kube/config:/root/.kube/config  podchef/backend
```

# Cluster init configurations

# Cluster - node - CMD's

Created: May 6, 2021 4:37 PM

# Node init script - [https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/) - [https://docs.docker.com/engine/install/ubuntu/](https://docs.docker.com/engine/install/ubuntu/)

```bash
echo "installing docker engine: -------------------------------------------------------------------\n"

echo "getting updates:\n"
sudo apt-get update
sudo apt-get upgrade -y

curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

echo "docker instalation complete\n"

echo "installing kubernetes: ---------------------------------------------------------------------\n"

echo "adding iptable rules:\n"

cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system

echo "getting updates:\n"
sudo apt-get update

echo "installing dependencies:\n"
sudo apt-get install -y apt-transport-https ca-certificates curl

echo "adding google cloud key:\n"
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

echo "adding kubernetes apt repository:\n"
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

echo "getting updates:\n"
sudo apt-get update

echo "installing kubelet kubeadm kubectl:\n"
sudo apt-get install -y kubelet kubeadm kubectl

echo "setting kubelet kubeadm kubectl on hold:\n"
sudo apt-mark hold kubelet kubeadm kubectl

echo "configuring cgroup driver: ---------------------------------------------------------------------\n"

cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

echo "restart docker en enable on boot:\n"

sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker

#echo "disabling swap:\n"

#sudo swapoff -a
```

# Disable swap

```bash
#abrir o ficheiro e comentar quaisquer linhas que contenham uma partição swap
sudo nano /etc/fstab
sudo reboot 0
```

# Set Static Ip address - /etc/netplan/00-installer-config.yaml

- alterar ips e placa de redes consoante o caso do utilizador

```bash
# This is the network config written by 'subiquity'
network:
  ethernets:
    ens33:
      dhcp4: no
      addresses:
        - 192.168.112.3/24
      gateway4: 192.168.112.2
      nameservers:
          addresses: [8.8.8.8]
  version: 2
```

- Restart network manager

```bash
sudo netplan apply
```

# Master init cluster

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```

# Workers join cluster

```bash
kubeadm join 192.168.112.3:6443 --token ihboah.yw5so0s21ucvjnho \
        --discovery-token-ca-cert-hash sha256:356c080f9771b9dca01c09c25a696f30a349fcb467173f5e5c0cb0edfffafa9a
```

# Master node configuration

```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

# Worker node configuration

```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/kubelet.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

```

# Give authorization to the API to execute kubectl commands - api-auth.yaml

```bash
kubectl apply -f https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/k8/api-auth.yaml
```

# NGINX- node load balancer/proxy

## Generate a self signed SSL certificate for Https

### ./san.cf

```bash
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext
[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
stateOrProvinceName         = State or Province Name (full name)
localityName               = Locality Name (eg, city)
organizationName           = Organization Name (eg, company)
commonName                 = Common Name (e.g. server FQDN or YOUR name)
[ req_ext ]
subjectAltName = @alt_names
[alt_names]
DNS.1   = example.com
```

```bash
sudo apt-get install openssl

sudo openssl req -x509 -nodes -days 1024 -newkey rsa:2048 -keyout selfsigned.key -out selfsigned.crt
```

## NGINX config file - nginx.conf

Este ficheiro contém um ip exemplo, importante mudar!!!!!!

```bash
wget https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/nginx/nginx.conf
```

## NGINX Dockerfile

```bash
wget https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/nginx/Dockerfile
```

## Build dockerfile, test configuration and then run custom NGINX image

```bash
sudo docker build -t podchef-nginx .
sudo docker run -d --name podchef-nginx --net=host podchef-nginx
```

## Changing NGINX config on container runtime (optional)

```bash
sudo docker exec -it podchef-nginx /bin/bash
vim /etc/nginx/nginx.conf
```

Reload container's config

```bash
sudo docker exec -it podchef-nginx nginx -s reload
```

# Create ingress - [https://www.haproxy.com/blog/dissecting-the-haproxy-kubernetes-ingress-controller/](https://www.haproxy.com/blog/dissecting-the-haproxy-kubernetes-ingress-controller/)

## Controller

```bash
kubectl apply -f https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/k8/ingress-controller.yaml
```

## Deploy the API on kubernetes

```yaml
kubectl apply -f https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/k8/api-deploy.yaml
```

## Config map

```yaml
kubectl apply -f https://raw.githubusercontent.com/goncalosn/pod-chef-back-end/main/infrastructure/k8/configmap.yaml
```
