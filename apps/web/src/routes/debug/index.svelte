<script context="module" lang="ts">
  import type { Load } from '@sveltejs/kit';

  // see https://kit.svelte.dev/docs#loading
  export const load: Load = async ({ fetch }) => {
    const response = await fetch('/debug/hello.text');

    if (response.ok) {
      const message = await response.text();

      return {
        props: { message },
      };
    }

    return {
      props: { message: null },
    };
  };
</script>

<script lang="ts">
  export let message: string         | null;
</script>

<h1>Message</h1>
<p>{message}</p>
