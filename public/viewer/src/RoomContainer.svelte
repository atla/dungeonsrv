<script>
  import { fade } from "svelte/transition";
  import { visible } from "./stores.js";
</script>

<style>
  .box {
    background: url("https://picsum.photos/500/300?blur");
    background-size: cover;

    border: #4c566a 1px solid;
    border-radius: 2px;
    box-shadow: 2px 2px 16px rgba(0, 0, 0, 0.1);
    margin: 0 0 1em 0;

    margin: auto;
    width: 40em;
  }
  .background {
    padding: 10em 1em 1em 1em;

    background-image: linear-gradient(
      #434c5e00,
      #434c5e00 5em,
      #434c5eaa 10em,
      #434c5eff
    );
  }
</style>

{#if $visible}
  <div class="box" transition:fade>
    <div class="background">
      <slot />
    </div>
  </div>
{/if}
