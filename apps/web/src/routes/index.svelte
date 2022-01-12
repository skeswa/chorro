<script context="module" lang="ts">
  export const prerender = true;
</script>

<svelte:head>
  <title>Chorro | Coming Soon</title>

  <style>
    html, body {
      margin: 0;
      padding: 0;
    }
  </style>
</svelte:head>

<section>
  <img src="/logo.svg" alt="Chorro logo">
  <p>coming soon</p>
</section>


<style>
  section {
    align-items: center;
    display: flex;
    flex-direction: column;
    justify-content: center;
    max-height: 100vh;
    max-width: 100vw;
    min-height: 100vh;
    min-width: 100vw;
  }

  img {
    max-height: 240px;
    max-width: 240px;
    min-height: 240px;
    min-width: 240px;
  }

  p {
    color: #3a425c;
    font-family: sans-serif;
    font-size: 36px;
  }
</style>
