# remoteav

Remotely access cameras and microphones

## Objectives

### Project

- [ ] Access cameras and microphones remotely
  - [ ] Frontend: Vite, React
  - [ ] Backend: Go
  - [ ] Api: GraphQL
  - [ ] Container: Docker
  - [ ] Container registry: GitHub Container Registry
  - [ ] Continuous delivery: GitHub Actions
  - [ ] Authentication: JWT

### Ease of life

- [ ] Developer mode
- [ ] One container that does both frontend and backend

## Dev logs

```bash
# Init go main then
go mod init

# Init frontend
npm create vite@latest

# Setup some kind of proxy in go
# using gin

# Start backend
sudo npm exec --yes -- nodemon@latest --ignore frontdent --ext go,graphql,html --exec 'sudo fuser -k 7777/tcp; sudo env "PATH=$PATH" go run ./dev || false'

# Automatically regenerate frontend graphql schema
npm --prefix frontend run graphql-codegen

# Start frontend
npm --prefix frontend run dev

# Simulate production build
npm --prefix frontend run build && rm -rf dist/dist && cp -r frontend/dist dist && sudo go run .
```
