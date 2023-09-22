# =============================================================================
#  Test Container for Vaious Go Versions (Last Updated: 2022-09-21)
# =============================================================================
# Default version: golang:alpine
ARG VARIANT="alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache gcc musl-dev

COPY ./go.mod /workspaces/go.mod
COPY ./go.sum /workspaces/go.sum

WORKDIR /workspaces

RUN go mod download

ENTRYPOINT go mod download && go test -race ./...
