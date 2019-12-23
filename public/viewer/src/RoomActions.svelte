<script>
  export let currentRoom;

  import { setCurrentRoom, visible } from "./stores.js";
  import { fade } from "svelte/transition";

  function handleAction(action) {
    console.dir("handle action type" + action.type);
    switch (action.type) {
      case "direction":
        visible.update(v => false);

        setTimeout(function() {
          setCurrentRoom(action.data.room);
          visible.update(v => true);
        }, 500);

        break;
      default:
        console.error("unhandled action type " + action.type);
    }
  }
</script>

<style>
  .actionContainer {
    text-align: left;
  }

  button {
    text-transform: uppercase;
    background: transparent;
    color: #8fbcbb;
    border: 1px #8fbcbb33 solid;
    border-radius: 0.5em;
    padding-left: 1em;
    padding-right: 1em;

    margin-right: 0.5em;
  }
</style>

<div class="actionContainer">
  {#each $currentRoom.actions as action}
    <button
      on:click={() => {
        handleAction(action);
      }}>
      {action.description}
    </button>
  {/each}
</div>
