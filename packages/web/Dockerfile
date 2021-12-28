# NOTE: though this `Dockerfile` is intended to build a new Docker container
# image for `@chorro/web`, it must be run from the root of the monorepo.

# BUILD ########################################################################

FROM node:17-alpine as build

# Host the app in `~/app` of the `node` user.
WORKDIR /home/node/app

# Copy over the stuff that `yarn` needs to do its job.
COPY package.json .
COPY yarn.lock .

# Copy over all relevant source code.
#
# # TODO(skeswa): put any shared packages needed by `web` here.
# COPY packages/shared ./packages/shared
COPY packages/web ./packages/web

# Fetch the dependencies we need to build `web`.
RUN yarn install --pure-lockfile --non-interactive

# # TODO(skeswa): Build any shared packages needed by `web` here.
# WORKDIR /home/node/app/packages/shared
# RUN yarn build

# Build `web`.
WORKDIR /home/node/app/packages/web
RUN yarn build

# RUN ##########################################################################

FROM node:17-alpine as run

# Host the app in `~/app` of the `node` user.
WORKDIR /home/node/app/packages/web

# Copy over `web`'s `package.json` and built server source directory.
COPY --from=build /home/node/app/packages/web/package.json /home/node/app/packages/web/package.json
COPY --from=build /home/node/app/packages/web/build /home/node/app/packages/web/build

# Specify that we're running in production.
ENV NODE_ENV production

# Fetch only the dependencies we need to run the server.
RUN yarn install --pure-lockfile --non-interactive --production

# Run `web` on port 3000.
EXPOSE 3000
CMD ["node", "./build"]