<svelte:head>
    <title>Todo List</title>
    <link rel="stylesheet" href="/global.css">
    <link rel="stylesheet" href="/build/bundle.css">
</svelte:head>

<script>
    import Task from './Task.svelte';
    import NewTask from './NewTask.svelte';
    
    let tasks = [];
    
    function handleAddTask(e) {
        tasks = [...tasks, e.detail];
    }
    
    function handleUpdateTask(updatedTask) {
        tasks = tasks.map(task => task.id === updatedTask.id ? updatedTask : task);
    }
    
    function handleDeleteTask(taskId) {
        tasks = tasks.filter(task => task.id !== taskId);
    }
    
    function handleToggleTaskCompletion(taskId) {
        tasks = tasks.map(task => task.id === taskId ? {...task, isComplete: !task.isComplete} : task);
    }
</script>

<main>
    <section>
        <h1>Todo List</h1>
        <NewTask on:addTask="{handleAddTask}" />
        <ul>
            {#each tasks as task (task.id)}
                <li>
                    <Task bind:task="{task}" on:updateTask="{handleUpdateTask}" on:deleteTask="{handleDeleteTask}" on:toggleTaskCompletion="{handleToggleTaskCompletion}" />
                </li>
            {/each}
        </ul>
    </section>
</main>