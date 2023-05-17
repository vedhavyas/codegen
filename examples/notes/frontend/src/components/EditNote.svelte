<script>
  import { createEventDispatcher } from 'svelte';
  import notesAPI from './noteAPI.js';
  
  let noteId;
  let noteContent;
  
  const dispatch = createEventDispatcher();
  
  async function saveChanges() {
    try {
      await notesAPI.updateNote(noteId, noteContent);
      dispatch('noteUpdated', { noteId, noteContent });
    } catch (error) {
      console.error('Error updating note:', error);
    }
  }
</script>

<style>
  /* Define your styles here */
</style>

<div>
  <textarea bind:value="{noteContent}"></textarea>
  <button on:click="{saveChanges}">Save Changes</button>
</div>