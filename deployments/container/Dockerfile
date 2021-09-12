FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache musl-dev gcc git build-base

WORKDIR /go/src/github.com/TrevorEdris/api-template/

# Naively copy everything. The final binary will be copied into a scratch container.
COPY . .

RUN go build -ldflags "-linkmode external -extldflags \"-static\" -s -w $LDFLAGS" -o the-binary cmd/main.go

# Copy the binary from the "builder" docker target into a scratch container
# to vastly reduce the overall size of the image
FROM scratch AS final

EXPOSE 8080
ENTRYPOINT ["/the-binary"]
COPY --from=builder /go/src/github.com/TrevorEdris/api-template/the-binary /the-binary
