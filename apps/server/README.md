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

Once you've installed [Go](https://go.dev/learn/), start the development server:

```bash
$ yarn dev
```

### Getting Started

```bash
$ go mod download
```

Before you run the dev server, you'll first need to create a `.env` file with
environment variables that suit your particular setup. As a starting point,
simply copy over the included `./env.example` and then tweak it to your heart's
content:

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
