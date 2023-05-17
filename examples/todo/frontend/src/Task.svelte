<svelte:options immutable={true} />
<script>
  export let task;
  export let index;

  function handleToggle() {
    this.$dispatch("tasktoggle", { id: task.id });
  }

  function handleEdit() {
    this.$dispatch("taskedit", { id: task.id, title: task.title });
  }

  function handleDelete() {
    this.$dispatch("taskdelete", { id: task.id });
  }
</script>

<div class="task">
  <input type="checkbox" checked={task.completed} on:change={handleToggle}>
  </input>
  <label class="title" contenteditable="true" on:blur={handleEdit}>
    {task.title}
  </label>
  <button on:click={handleDelete}>Delete</button>
</div>

<style>
  .task {
    display: flex;
    align-items: center;
  }

  .task input[type=checkbox] {
    margin-right: 12px;
  }

  .task label {
    flex: 1;
  }

  .task button {
    background-color: red;
    color: white;
    padding: 3px 8px;
    border: none;
    border-radius: 3px;
    margin-left: 12px;
    cursor: pointer;
  }

  .task button:hover {
    background-color: #d44;
  }
</style>