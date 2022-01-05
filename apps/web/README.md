# `@chorro/web`

Node.js application that serves the Chorro web application, which is built with
[SvelteKit](https://kit.svelte.dev/).

## Developing

### Pre-requisites

- [Git](https://git-scm.com/) 2.34+
- [Node.js](https://nodejs.dev/) v17+
- [Yarn](https://nodejs.dev/) v1.22+

### Getting Started

```bash
$ yarn install
```

Before you run the dev server, you'll first need to create a `.env` file with
environment variables that suit your particular setup. As a starting point,
simply copy over the included `./env.example` and then tweak it to your heart's
content:

```bash
$ cp .env.example .env
```

### Running the development server

Once you've created a project and installed dependencies with `yarn`, start a
development server:

```bash
$ yarn dev
```

or start the server and open the app in a new browser tab:

```bash
$ yarn dev -- --open
```

## Building the compiled binary

You can generate the self-contained, compiled Node.js binary within the
`./build` directory by activating SvelteKit's `build` step:

```bash
$ yarn build
```

## Building and deploying the Docker image

You can deploy to Kubernetes using using the `release` and `deploy` scripts:

```bash
$ yarn release
$ yarn deploy
```

## Future

In future, it would make a whole lot of sense to deploy to
https://developers.cloudflare.com/workers.