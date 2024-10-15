#####################################
#   STEP 1 build executable binary  #
#####################################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk --no-cache add \
        alpine-sdk \
        pkgconf && \
        rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .
COPY go.sum .

RUN go mod download

# Build the binary.
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo \
     -o gitrocket /app/cmd/gitrocket/main.go

#####################################
#   STEP 2 build a small image      #
#####################################
FROM scratch

# Copy our static executable.
COPY --from=builder /app/gitrocket /app/gitrocket
COPY --from=builder /app/cfg/gitrocket.yaml /app/cfg/gitrocket.yaml

# Run the hello binary.
ENTRYPOINT ["/app/gitrocket"]