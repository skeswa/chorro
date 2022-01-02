# NOTE: though this `Dockerfile` assumes that `yarn build` was just invoked.

# NOTE: though this `Dockerfile` assumes the cwd is the root directory of the
# repo.

FROM node:17-alpine

# Host the app in `~/app` of the `node` user.
WORKDIR /home/node/app

# Copy over the built server directory, and its corresponding dependency
# configuration.
COPY ./yarn.lock ./yarn.lock
COPY ./apps/web/package.json ./package.json
COPY ./apps/web/build ./build

# Specify that we're running in production.
ENV NODE_ENV production

# Fetch only the dependencies we need to run the server.
RUN yarn install --pure-lockfile --non-interactive --production

# Run `web` on port 3000.
EXPOSE 3000
CMD ["node", "./build"]