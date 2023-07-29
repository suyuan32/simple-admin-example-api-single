FROM golang:1.20.5-alpine3.17 as builder

# Define the project name | 定义项目名称
ARG PROJECT=example

WORKDIR /build
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /build/${PROJECT}_api ${PROJECT}.go

FROM alpine:latest

# Define the project name | 定义项目名称
ARG PROJECT=example
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=example.yaml
# Define the author | 定义作者
ARG AUTHOR="example@example.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /build/${PROJECT}_api ./
COPY --from=builder /build/etc/${CONFIG_FILE} ./etc/

EXPOSE 8081

ENTRYPOINT ./${PROJECT}_api -f etc/${CONFIG_FILE}