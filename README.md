# 校大学生科协报名表单 ![docker](https://github.com/MR-Addict/stas-joinus/actions/workflows/docker.yml/badge.svg)

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。

后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中，最后可以编译成一个很小的 Docker 镜像，你可以直接从 Docker Hub 拉取该镜像。

> 注意，这里的数据库使用的 mongodb，所以你得先有一个 mongodb 的数据库，然后再添加到下面的环境变量中。

## 1. 部署项目

### 1.1 使用 Docker 命令

你可以使用下面的一行 Docker 命令来完成部署，但是不推荐这么做：

```sh
docker run -d --name joinus -p 4000:4000 -e ADMIN_USERNAME=admin_username -e ADMIN_PASSWORD=admin_password -e JWT_SECRET=jwt_secret -e MONGODB_URI=mongodb://admin:password@example.com mraddict063/stas-joinus
```

### 1.2 使用 Docker-Compose（推荐）

如果你有安装 docker-compose，可以新建一个 docker-compose.yaml 文件，根据需求修改里面的环境变量：

```yaml
version: "3"
services:
  joinus:
    image: mraddict063/stas-joinus
    restart: unless-stopped
    ports:
      - 4000:4000
    environment:
      - JWT_SECRET=jwt_secret
      - ADMIN_USERNAME=admin_username
      - ADMIN_PASSWORD=admin_password
      - MONGODB_URI=mongodb://admin:password@example.com
```

然后使用下面的命令启动项目即可：

```sh
docker-compose up -d
```
