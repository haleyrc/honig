FROM golang:1.23-bookworm AS builder

WORKDIR /go/src/app

RUN apt update && \
    apt upgrade -y && \
		dpkg --add-architecture amd64

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build -o /go/bin/server .

###############################################################################

FROM gcr.io/distroless/base-debian12

LABEL "org.opencontainers.image.source"="https://github.com/haleyrc/honig"

COPY --from=builder /go/bin/server /

CMD ["/server"]
