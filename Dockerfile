FROM cgr.dev/chainguard/go:latest AS builder
COPY . /app
RUN cd /app && go build

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/go-rest-api /usr/bin/
EXPOSE 8000
ENTRYPOINT ["/usr/bin/go-rest-api"]