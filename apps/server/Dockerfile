# NOTE: though this `Dockerfile` assumes the cwd is the root directory of the
# repo.

# builder ######################################################################

FROM golang:1.17.5-alpine AS builder
LABEL org.opencontainers.image.source = "https://github.com/skeswa/chorro"

# Host the app in `/app`.
WORKDIR /app

# Copy over just the information we need to fetch dependencies, and then
# download the dependecies using `go mod download`. We do it this way because it
# benefits from improved cacheability from Docker's perspective.
COPY ./apps/server/go.mod ./go.mod
COPY ./apps/server/go.sum ./go.sum
RUN go mod download

# Copy over all the source code.
COPY ./apps/server .

# Build the `server` binary.
RUN go build -o server .

# runner #######################################################################

FROM alpine

# Host the app in `/app`.
WORKDIR /app

# Copy the `server` binary we built in the previous step.
COPY --from=builder /app/server ./server

# Specify that we're running in production.
ENV NODE_ENV production

# Specify other miscellaneous application configuration.
ENV HOME_URL https://chorro.app

# Specify that we're running on port 8000.
ENV PORT 8000
EXPOSE 8000

ENTRYPOINT ["./server"]