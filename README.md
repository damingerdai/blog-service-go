# blog service
---

# setup
## create volume
```bash
docker volume create daming-blog-volume
```
## create network
```bash
docker network create daming-blog-network
```
## create mysql database
```bash
docker-compose build db
```

> if you meet this issue:
```
db use images
```
please run:
```bash
docker-compose up --force-recreate db
```