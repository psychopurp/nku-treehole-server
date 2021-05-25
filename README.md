# NKU-Treehole-Server

Backend for NKU's treehole.

## 运行

### 1.配置数据库 mysql

docker 配置 mysql (或者本地启动一个 mysql)

```bash
cd docker
docker-compose up -d
```

## 2.建表

> 在数据库里运行 treehole.sql

## 3.配置文件

/conf/config.dev.yml

4.启动

> go run .
