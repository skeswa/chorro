# NOTE: though this `Dockerfile` is intended to build a new Docker container
# image for `@chorro/web`, it must be run from the root of the monorepo.

# BUILD ########################################################################

FROM node:17-alpine as build

# Host the monorepo in `~/repo` of the `node` user.
WORKDIR /home/node/repo

# Copy over all relevant source code.
COPY . .

# Fetch the dependencies we need to build `web`.
RUN yarn install --pure-lockfile --non-interactive

# Build everything.
RUN yarn turbo run build

# RUN ##########################################################################

FROM node:17-alpine as run

# Host the app in `~/app` of the `node` user.
WORKDIR /home/node/app

# Copy over `web`'s `package.json` and built server source directory.
COPY --from=build /home/node/repo/apps/web/package.json .
COPY --from=build /home/node/repo/apps/web/build ./build

# Specify that we're running in production.
ENV NODE_ENV production

# Fetch only the dependencies we need to run the server.
RUN yarn install --pure-lockfile --non-interactive --production

# Run `web` on port 3000.
EXPOSE 3000
CMD ["node", "./build"]