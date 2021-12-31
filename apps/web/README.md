# `@chorro/web`

Node.js application that serves the Chorro web application, which is built with
[SvelteKit](https://kit.svelte.dev/).

## Developing

Once you've created a project and installed dependencies with `yarn`, start a
development server:

```bash
yarn run dev

# or start the server and open the app in a new browser tab
yarn run dev -- --open
```

## Building

You can generate the self-contained, compiled Node.js binary within the
`./build` directory by activating SvelteKit's `build` step:

```bash
yarn run build
```

## Deploying

You can deploy to Kubernetes using using the deploy script:

```bash
yarn run deploy
```

This will tweak the `deployment.yml`, which should be subsequently committed and
pushed.