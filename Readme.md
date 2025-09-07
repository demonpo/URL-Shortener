Setup:
<br/>
First install [Taskfile](https://taskfile.dev/installation/)

Build docker image:
```bash
docker build -f Dockerfile.prod --ignorefile .dockerignore.prod -t url-shortener:prod .
```
Run container:
```bash
docker run --rm -p 3000:3000 --env-file .env --name url-shortener url-shortener:prod
```
(Optional) Multi-arch build & push:
```bash
docker buildx build --platform linux/amd64,linux/arm64 -f Dockerfile.prod -t your-repo/url-shortener:latest --push .
```
Explanation:
- The trailing '.' is the build context path that was missing (caused the error).
