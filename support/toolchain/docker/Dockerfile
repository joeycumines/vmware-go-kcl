FROM golang:1.24
ENV PATH=/go/bin:/src/bin:/root/go/bin:/usr/local/go/bin:$PATH \
    GOPATH=/go:/src
RUN set -x && \
    go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.2.2 && \
    go install golang.org/x/tools/cmd/...@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    curl -sfL https://raw.githubusercontent.com/securego/gosec/d2d3ae66bd8d340b78b5142b6fe610691783c2fe/install.sh | sh -s v2.22.5 && \
    chmod -R a+rw /go
