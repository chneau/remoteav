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
RUN apk add --no-cache portaudio-dev gcc musl-dev
RUN go build -o /remoteav

FROM alpine AS final
RUN apk add --no-cache portaudio
COPY --from=golang-builder /remoteav /bin/remoteav
ENTRYPOINT ["remoteav"]
