<script>
  import { fade } from "svelte/transition";
  import { visible } from "./stores.js";
</script>

<style>
  .box {
    background: #434c5eff;

    border: #4c566a 1px solid;
    border-radius: 2px;
    box-shadow: 2px 2px 16px rgba(0, 0, 0, 0.1);
    margin: 1em 0 1em 0;

    width: 40em;
  }

  input {
    text-transform: lowercase;
    background: transparent;
    color: #8fbcbb;
    border: 0px;
    margin-left: 0.2em;
    margin-right: 0.2em;
  }
</style>

{#if $visible}
  <input class="box" type="text" transition:fade />
{/if}
