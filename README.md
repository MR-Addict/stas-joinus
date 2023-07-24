# 校大学生科协报名表单

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中。

## 1. 部署项目

### 1.1 直接使用 Docker 命令

使用下面的 Docker 命令：

```sh
docker run -d --name joinus -p 4000:4000 -e USERNAME=username -e PASSWORD=password -e MONGODB_URI=mongodb://admin:password@example.com mraddict063/stas-joinus
```

### 1.2 使用 Docker-Compose

部署项目时在根目录下新建下面的 docker-compose.yaml 文件，根据要求修改里面的环境变量：

```yaml
version: "3"
services:
  joinus:
    image: mraddict063/stas-joinus
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
