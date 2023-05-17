<script>
  import { writable } from 'svelte/store';
  import noteAPI from './noteAPI.js';

  const title = writable('');
  const content = writable('');

  async function addNote() {
    await noteAPI.addNote({
      title: $title,
      content: $content
    });

    title.set('');
    content.set('');
  }
</script>

<style>
  .add-note {
    display: flex;
    flex-direction: column;
    margin-bottom: 16px;
  }

  .add-note__input,
  .add-note__textarea {
    margin-bottom: 8px;
  }
</style>

<div class="add-note">
  <input
    class="add-note__input"
    type="text"
    placeholder="Title"
    bind:value={$title}
  />

  <textarea
    class="add-note__textarea"
    placeholder="Content"
    rows="4"
    bind:value={$content}
  ></textarea>

  <button on:click={addNote}>
    Add Note
  </button>
</div>