# NOTE: though this `Dockerfile` assumes that `yarn build` was just invoked.

# NOTE: though this `Dockerfile` assumes the cwd is the root directory of the
# repo.

FROM alpine
LABEL org.opencontainers.image.source = "https://github.com/skeswa/chorro"

# Host the app in `/app`.
WORKDIR /app

# Copy over the compiled binary for the target platform.
COPY ./apps/server/build/server-linux-amd64 ./server

# Specify that we're running in production.
ENV NODE_ENV production

# Specify that we're running on port 8000.
ENV PORT 8000
EXPOSE 8000

ENTRYPOINT ["./server"]