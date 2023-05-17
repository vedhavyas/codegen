<svelte:head>
	<style>
		.new-task {
			display: flex;
			flex-direction: row;
			align-items: center;
		}
		.new-task input {
			flex-grow: 1;
			margin-right: 8px;
		}
	</style>
</svelte:head>

<script>
	import { createEventDispatcher } from "svelte";

	const dispatch = createEventDispatcher();

	let inputValue = "";

	function onSubmit() {
		if (inputValue.trim() !== "") {
			dispatch("addTask", { title: inputValue.trim() });
			inputValue = "";
		}
	}
</script>

<div class="new-task">
	<input
		placeholder="Enter new task title"
		bind:value="{inputValue}"
		on:keydown="{(e) => {if (e.key === 'Enter') onSubmit()}}"
	/>
	<button on:click="{onSubmit}">Add Task</button>
</div>