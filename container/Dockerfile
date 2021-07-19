FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache musl-dev gcc git build-base

WORKDIR /go/src/github.com/TrevorEdris/template-golang-kubernetes/

# go.mod and go.sum
COPY go.* .

# go source files
COPY *.go .

RUN go build -ldflags "-linkmode external -extldflags \"-static\" -s -w $LDFLAGS" -o the-binary

# Copy the binary from the "builder" docker target into a scratch container
# to vastly reduce the overall size of the image
FROM scratch AS final

EXPOSE 8080
ENTRYPOINT ["/the-binary"]
COPY --from=builder /go/src/github.com/TrevorEdris/template-golang-kubernetes/the-binary /the-binary
