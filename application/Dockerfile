FROM golang:1.17 AS app
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /root/app
COPY server/. .
RUN CGO_ENABLED=0 go build -v -o "app" .

FROM node:16.17.1-alpine AS web
WORKDIR /root/web
COPY web/package*.json ./
RUN yarn config set registry https://registry.npmmirror.com/ && yarn install
COPY web/. .
RUN yarn run build:prod

FROM scratch
WORKDIR /root/app-cli/
COPY --from=app /root/app/app ./
COPY --from=app /root/app/config.yaml ./
COPY --from=web /root/web/dist/ ./dist/
ENTRYPOINT ["./app"]