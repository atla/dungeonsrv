<script>
  import RoomContainer from "./RoomContainer.svelte";
  import RoomContent from "./RoomContent.svelte";
  import RoomInput from "./RoomInput.svelte";
  import RoomActions from "./RoomActions.svelte";

  import { afterUpdate } from "svelte";

  import { currentRoom, rooms, setCurrentRoom, visible } from "./stores.js";

  setCurrentRoom(1);
</script>

<style>
  :global(body) {
    background-color: #3b4252;
    color: #d8dee9;
    transition: background-color 0.3s;
  }

  main {
    text-align: center;

    margin: 0 auto;
    max-width: none;
  }
  h1 {
    text-transform: uppercase;
    font-size: 1em;
    text-align: left;
    font-weight: 100;
    color: #5b6272;
  }
</style>

<main>
  <h1>Dungeonsrv Viewer</h1>
  <RoomContainer>
    <RoomContent {currentRoom} />
    <RoomActions {currentRoom} />

  </RoomContainer>
  <RoomInput />
</main>
