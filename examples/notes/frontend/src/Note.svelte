<svelte:options tag="Note" />

<script>
  import EditNote from './components/EditNote.svelte';
  import DeleteNote from './components/DeleteNote.svelte';

  export let id = 0;
  export let text = "";
</script>

<div class="note">
  <h3>{text}</h3>
  <EditNote id={id} text={text} />
  <DeleteNote id={id} />
</div>

<style>
  .note {
    display: flex;
    align-items: center;
    justify-content: space-between;
    border: 1px solid #ccc;
    padding: 10px;
    margin-bottom: 5px;
  }
  h3 {
    margin: 0;
  }
</style>