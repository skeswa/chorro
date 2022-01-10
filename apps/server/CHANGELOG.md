# @chorro/server

## 0.3.2

### Patch Changes

- 24fe76b: Moved all server endpoints under /api.

## 0.3.1

### Patch Changes

- 81dab59: The GraphQL API now does real stuff which is pretty cool.

## 0.3.0

### Minor Changes

- 0592283: Added a GraphQL API (and a playground too) while ditching gofiber.io
  in favor of plain vanilla Go `net/http`.

## 0.2.1

### Patch Changes

- 1761a4b: Got connection to Postgres up and running - next stop, GraphQL
  station!

## 0.2.0

### Minor Changes

- 480a7d1: Enabled session management in the server.

### Patch Changes

- 480a7d1: Made `yarn changeset` and `yarn reset` work in every individual
  package.

## 0.1.0

### Minor Changes

- 9117eb0: Server is now capable of authenticating against Google via OAuth 2.0.

### Patch Changes

- 9117eb0: Server now reloads upon file change while in development.
- 9117eb0: Started using .env files to manage local environment variables.
- 9117eb0: We now have a placeholder logo that I hacked together in a code
  editor with a blindfold on.

## 0.0.2

### Patch Changes

- de13d6c: Added a backend for Chorro - though it doesn't really do anything
  yet...
