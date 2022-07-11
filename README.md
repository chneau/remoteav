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
npm exec --yes -- nodemon@latest --signal SIGTERM --ext go --exec 'sudo env "PATH=$PATH" go run . || false'

# Start frontend
npm --prefix frontend run dev
```
