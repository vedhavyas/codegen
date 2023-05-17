<script>
  import NotesList from './components/NotesList.svelte';
  import AddNote from './components/AddNote.svelte';
  import { onMount } from 'svelte';
  import noteAPI from './components/noteAPI.js';

  let notes = [];

  onMount(async () => {
    notes = await noteAPI.getNotes();
  });

</script>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>

<main>
  <h1>Note Taking App</h1>
  <AddNote />
  <NotesList {notes} />
</main>