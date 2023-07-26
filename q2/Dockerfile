FROM golang:1.19 AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /run
COPY --from=build /build/app ./
ENTRYPOINT ["./app"]
