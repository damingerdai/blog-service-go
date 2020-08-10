# blog service
---

# setup
## create volume
```bash
docker volume create daming-blog-volume
```
## create network
```bash
docker network create daming-blog-volume
```
## create mysql database
```bash
docker-compose build db
```

> if you meet this issue:
```bash
docker-compose up --force-recreate db
```