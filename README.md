# 校大学生科协报名表单 ![docker](https://github.com/MR-Addict/stas-joinus/actions/workflows/docker.yml/badge.svg)

<img src="preview.png" alt="preview" width="400">

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。

前端是 [Svelte](https://svelte.dev)，后端是 [GO](https://go.dev)，分别放在了 `client` 和 `server` 两个目录中，最后可以编译成一个很小的 Docker 镜像，你可以直接从 Docker Hub 拉取该镜像。

## 1. 环境变量

本项目用到了以下几个环境变量：

| 变量名     | 解释                     | 备注                 |
| :--------- | :----------------------- | :------------------- |
| START_TIME | 报名开始时间             | 必需                 |
| END_TIME   | 报名结束时间             | 必需                 |
| ADMIN_PASS | 用来登录后台的管理员密码 | 必需                 |
| CORS       | 跨域域名设置             | 可选，默认不允许跨域 |
| PORT       | 项目监听端口             | 可选，默认为 4000    |

## 2. 部署项目

本项目可以通过编译成单个可执行文件，同时包含前端、后端和数据库，经优化编译后的可执行文件不到 **7M**！

如果你想了解如何编译的话可以参考本项目的 [Dockerfile](Dockerfile)。

理论上本项目是不需要 Docker 就可以启动的，但是使用 Docker 可以方便管理和部署。

新建一个 docker-compose.yaml 文件，根据需要修改其中的环境变量：

```yaml
services:
  joinus:
    image: mraddict063/stas-joinus
    restart: unless-stopped
    ports:
      - 4000:4000
    environment:
      - START_TIME=2024-09-03T16:00:00.000Z
      - END_TIME=2024-09-15T15:59:59.000Z
      - ADMIN_PASS=password
    volumes:
      - ./data:/data
```

然后使用下面的命令启动项目即可：

```sh
docker-compose up -d
```

默认情况下该项目会在 **4000** 端口启动，并在当前目录下创建一个 **data/data.db** 的数据库文件，你可以通过浏览器访问 `http://localhost:4000` 来查看项目。
