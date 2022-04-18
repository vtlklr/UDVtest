FROM golang:1.16 AS build_dep_cache
ENV GO111MODULE=on
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

FROM build_dep_cache AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ./
RUN go build -o udvtest ./main.go

FROM debian:10-slim
WORKDIR /app
COPY --from=build /build/udvtest ./udvtest
ENTRYPOINT [ "/app/udvtest" ]
EXPOSE 8000