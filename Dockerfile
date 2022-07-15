FROM node:alpine AS node-builder
WORKDIR frontend
COPY frontend/package.json frontend/package-lock.json .
RUN npm clean-install
COPY frontend .
RUN npm run build
