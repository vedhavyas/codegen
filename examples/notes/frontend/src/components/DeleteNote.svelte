<script>
  import { deleteNote } from './noteAPI.js';

  export let noteId;

  async function handleDelete() {
    await deleteNote(noteId);
    window.location.reload();
  }
</script>

<button on:click={handleDelete}>Delete</button>