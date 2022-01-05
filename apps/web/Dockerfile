# NOTE: though this `Dockerfile` assumes the cwd is the root directory of the
# repo.

# builder ######################################################################

FROM node:17-alpine AS builder
LABEL org.opencontainers.image.source = "https://github.com/skeswa/chorro"

# Host the repo in `~/repo` of the `node` user.
WORKDIR /home/node/repo

# Copy over all the source code.
COPY . .

# Fetch only the dependencies we need to run the server.
RUN yarn install --pure-lockfile --non-interactive

# Move into the `web` app directory.
WORKDIR /home/node/repo/apps/web

# Specify that we're building for production.
ENV NODE_ENV production

# Specify the production server base URL.
ENV VITE_SERVER_BASE_URL https://chorro.app

# Build the production executable.
RUN yarn build

# runner ######################################################################

FROM node:17-alpine

# Host the app in `~/app` of the `node` user.
WORKDIR /home/node/app

# Copy over just the information we need to fetch dependencies, and then
# download the dependecies using `yarn install`. We do it this way because it
# benefits from improved cacheability from Docker's perspective.
COPY --from=builder /home/node/repo/apps/web/package.json ./package.json
COPY --from=builder /home/node/repo/yarn.lock ./yarn.lock
RUN yarn install --pure-lockfile --non-interactive --production

# Copy over the production executable files built in the previous step.
COPY --from=builder /home/node/repo/apps/web/build .

# Specify that we're running in production.
ENV NODE_ENV production

# Run `web` on port 3000.
ENV PORT 3000
EXPOSE 3000

CMD ["node", "./index.js"]