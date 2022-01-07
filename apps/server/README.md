# `@chorro/server`

Go application that interfacing between Chorro's frontends and its database
(amond other things).

## Developing

### Pre-requisites

- Tools & Launguages
  - [Git](https://git-scm.com/) 2.34+
  - [Go](https://go.dev/) v1.17.5+
  - [Node.js](https://nodejs.dev/) v17+
  - [Yarn](https://nodejs.dev/) v1.22+
- Databases
  - [Postgres](https://www.postgresql.org/) v14.1
  - [Redis](https://www.postgresql.org/) v6+

### Getting Started

1.  Create a Postgres database, user, and password for Chorro
    [guide](https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e)
2.  Create a `.env` file with environment variables that suit your particular
    setup - use the included `./env.example` as a starting point:

    ```bash
    $ cp ./env.example .env
    ```

3.  You need to make sure you download all the necessary Go dependencies

    ```bash
    $ go mod download
    ```

### Running the dev server

The dev server should restart if its source code is changed while it is running:

```bash
$ yarn dev
```

## Building the self-contained binary

You can generate the self-contained, compiled Go binary within the `./build`
using the `build` step:

```bash
$ yarn build
```

## Building and deploying the Docker image

You can deploy to Kubernetes using using the `release` and `deploy` scripts:

```bash
$ yarn release
$ yarn deploy
```
