# 1. Install Docker CE

Use curl to download and install script

```shell
curl -sSL https://get.daocloud.io/docker | sh
```

(If user is root, please ignore) give common user permissions to operate docker

```shell
sudo usermod -aG docker xxxx # need to restart after execute
```

# 2. configure Docker self-booting

```shell
sudo systemctl enable docker
sudo systemctl start docker
```

Check docker information

```shell
docker info
```

# 3. Install Docker Compose

Download Docker Compose

```shell
sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

Configure execution rights

```shell
sudo chmod +x /usr/local/bin/docker-compose
```

Create soft link

```shell
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```
Check if it's installed successfully
```shell
docker-compose -v
```
