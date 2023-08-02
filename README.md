# 校大学生科协报名表单 ![docker](https://github.com/MR-Addict/stas-joinus/actions/workflows/docker.yml/badge.svg)

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。

后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中，最后可以编译成一个很小的 Docker 镜像，你可以直接从 Docker Hub 拉取该镜像。

## 1. 环境变量

本项目用到了以下三个环境变量：

| 变量名      | 解释                                                    |
| :---------- | :------------------------------------------------------ |
| PASSWORD    | 用来登录后台管理的管理员密码                            |
| JWT_SECRET  | JWT 用来加密使用的，可以随便设置一个安全的字符串        |
| MONGODB_URI | mongodb 数据库的地址，所以你得先有一个 mongodb 的数据库 |

## 2. 部署项目

本项目可以通过编译形成一个单个可执行文件，同时包含前后端，通过优化编译出的可执行文件大小 **5M** 多一点，非常的小！

所以理论上是不需要 Docker 容器的，但是有了 Docker 容器更方便管理，如何编辑可以参考 [Dockerfile](Dockerfile)。

### 2.1 使用 Docker 命令

你可以使用下面的一行 Docker 命令来完成部署，但是不推荐这么做：

```sh
docker run -d --name joinus -p 4000:4000 -e PASSWORD=password -e JWT_SECRET=jwt_secret -e MONGODB_URI=mongodb://admin:password@example.com mraddict063/stas-joinus
```

### 2.2 使用 Docker-Compose（推荐）

如果你有安装 docker-compose，可以新建一个 docker-compose.yaml 文件，根据需要修改里面的环境变量：

```yaml
version: "3"
services:
  joinus:
    image: mraddict063/stas-joinus
    restart: unless-stopped
    ports:
      - 4000:4000
    environment:
      - PASSWORD=password
      - JWT_SECRET=jwt_secret
      - MONGODB_URI=mongodb://admin:password@example.com
```

然后使用下面的命令启动项目即可：

```sh
docker-compose up -d
```
