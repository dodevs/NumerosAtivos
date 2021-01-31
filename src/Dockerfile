# Build stage
FROM golang:alpine AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /dockerdev

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .

RUN go build -o /wppdiscovery

# RUN STAGE
FROM golang:alpine AS runner

WORKDIR /
COPY --from=build-env /wppdiscovery /

CMD ["/wppdiscovery"]