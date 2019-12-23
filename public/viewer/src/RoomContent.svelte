<script>
  import { fade } from "svelte/transition";

  export let currentRoom;
</script>

<style>
  h1 {
    text-transform: uppercase;
    color: #eceff4;
    text-align: left;

    font-family: "Montserrat", sans-serif;
  }

  p {
    text-align: justify;
    font-family: "Montserrat", sans-serif;
  }
</style>

<article>
  <h1>{$currentRoom.title}</h1>
  <p>{$currentRoom.description}</p>
</article>
