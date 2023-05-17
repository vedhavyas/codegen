<script>
    import { onMount } from "svelte";
    import Note from "../Note.svelte";
    import { getNotes } from "./noteAPI.js";

    let notes = [];

    onMount(async () => {
        notes = await getNotes();
    });
</script>

<style>
    .notes-list {
        display: flex;
        flex-direction: column;
        gap: 10px;
        padding: 10px;
    }
</style>

<div class="notes-list">
    {#each notes as note}
        <Note {note} />
    {/each}
</div>