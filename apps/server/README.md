# `@chorro/server`

Go application that interfacing between Chorro's frontends and its database
(amond other things).

## Developing

Once you've installed [Go](https://go.dev/learn/), start the development server:

```bash
$ yarn run dev
```

## Building

You can generate the self-contained, compiled Node.js binary within the
`./build` directory by activating SvelteKit's `build` step:

```bash
$ yarn run build
```

## Deploying

You can deploy to Kubernetes using using the `release` and `deploy` scripts:

```bash
$ yarn run release
$ yarn run deploy
```
