# Compile binary in full Go container
FROM golang:1.14 as builder
LABEL stage="builder"
WORKDIR /opt/azdevman
COPY go.mod go.sum /
RUN set -ex \
    && go mod download
COPY . .
RUN set -ex \
    && make binary

# Build final lightweight container with the binary
FROM alpine as final
WORKDIR /opt/
COPY --from=builder /opt/azdevman/bin/azdevman .
ENTRYPOINT [ "/opt/azdevman" ]
