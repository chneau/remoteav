# remoteav

[![.github/workflows/publish.yml](https://github.com/chneau/remoteav/actions/workflows/publish.yml/badge.svg)](https://github.com/chneau/remoteav/actions/workflows/publish.yml)

Remotely access cameras and microphones

## Run with docker

```bash
docker run --privileged -p 7777:7777 --rm -it ghcr.io/chneau/remoteav
```

## Objectives

### Project

- [ ] Access cameras and microphones remotely
  - [x] Frontend: Vite, React
  - [x] Backend: Go
  - [x] Api: GraphQL
  - [x] Container: Docker
  - [x] Container registry: GitHub Container Registry
  - [x] Continuous delivery: GitHub Actions
  - [ ] Authentication: JWT

### Ease of life

- [x] Developer mode
- [x] One container that does both frontend and backend

## Dev logs

```bash
# Init go main then
go mod init

# Init frontend
npm create vite@latest

# Add yourself to these groups
sudo adduser $USER video
sudo adduser $USER audio
exec sudo su -l $USER

# Start backend
npm exec --yes -- nodemon@latest --ignore frontend --ext go,graphql,html --exec 'fuser -k 7777/tcp; go run ./dev || false'

# Start frontend
npm --prefix frontend run dev

# Automatically regenerate frontend graphql schema
npm --prefix frontend run graphql-codegen

# Simulate production build
npm --prefix frontend run build && rm -rf dist/dist && cp -r frontend/dist dist && go run .
```
