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
curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.4/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
```

Configure execution rights

```shell
sudo chmod +x /usr/local/bin/docker-compose
```

Check if it's installed successfully

```shell
docker-compose -v
```
