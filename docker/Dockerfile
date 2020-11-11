FROM golang:1.15.4 AS build
WORKDIR /app
COPY main.go .
RUN GO111MODULE=on go build -o proba main.go

FROM scratch
COPY --from=build /app/proba .
ENTRYPOINT ["./proba"]

