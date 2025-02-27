FROM golang:1.23 AS go-builder
COPY main.go /
RUN CGO_ENABLED=0 go build -ldflags "-w" -o /main /main.go

FROM scratch
COPY --from=go-builder --chmod=755 /main /simple-web
CMD ["./simple-web"]
