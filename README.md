# 校大学生科协报名表单 ![docker](https://github.com/MR-Addict/stas-joinus/actions/workflows/docker.yml/badge.svg)

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。

后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中，最后可以编译成一个很小的 Docker 镜像，你可以直接从 Docker Hub 拉取该镜像。

## 1. 环境变量

本项目用到了以下两个环境变量：

| 变量名     | 解释                                             |
| :--------- | :----------------------------------------------- |
| PASSWORD   | 用来登录后台管理的管理员密码                     |
| JWT_SECRET | JWT 用来加密使用的，可以随便设置一个安全的字符串 |

## 2. 部署项目

本项目可以通过编译成单个可执行文件，同时包含前端、后端和数据库，优化编译后的可执行文件仅有 **7M** 多一点，非常的小！

因此理论上本项目是不需要 Docker 就可以部署的，但是使用 Docker 可以方便管理和部署，如果你想了解如何编译的话可以参考本项目的 [Dockerfile](Dockerfile)。

新建一个 docker-compose.yaml 文件，根据需要修改 PASSWORD：

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
    volumes:
      - ./data:/data
```

然后使用下面的命令启动项目即可：

```sh
docker-compose up -d
```
