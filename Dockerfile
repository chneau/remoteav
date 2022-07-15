FROM node:alpine AS node-builder
WORKDIR node
COPY frontend/package.json frontend/package-lock.json .
RUN npm clean-install
COPY frontend .
RUN npm run build

FROM golang:alpine AS golang-builder
WORKDIR golang
COPY go.mod go.sum .
RUN go mod download
COPY . .
COPY --from=node-builder /node/dist dist/dist
RUN CGO_ENABLED=0 go build -o /remoteav -ldflags '-s -w -extldflags "-static"'

FROM alpine AS final
COPY --from=golang-builder /remoteav /bin/remoteav
ENTRYPOINT ["remoteav"]
