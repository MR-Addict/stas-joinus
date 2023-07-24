# 校大学生科协报名表单

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中。

## 1. 编译项目

编译 Docker 镜像：

```sh
docker build -t joinus .
```

## 2. 部署项目

部署项目需要新建下面的 docker-compose.yaml 文件，请根据要求修改里面的环境变量：

```yaml
version: "3"
services:
  joinus:
    image: joinus
    restart: unless-stopped
    ports:
      - 4000:4000
    environment:
      - USERNAME=dashboard_username
      - PASSWORD=dashboard_password
      - MONGODB_URI=mongodb://admin:password@example.com
```

启动项目：

```sh
docker-compose up -d
```
