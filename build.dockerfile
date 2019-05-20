FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code
ADD go.mod go.sum /code/
RUN go mod download
ADD . .
RUN go build -o /exporter main.go

FROM gcr.io/distroless/base
EXPOSE 9421
WORKDIR /
COPY --from=builder /exporter /usr/bin/exporter
ENTRYPOINT ["/usr/bin/exporter"]
