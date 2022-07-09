<script>
    import Counter from './lib/_counter.svelte'
    let _times_to_be_rendered = 2
    let times_to_be_rendered = 2
    const update = () => {
        //@ts-ignore
        times_to_be_rendered = parseInt(_times_to_be_rendered)
    }
</script>

<article>
<h1>Counter: </h1>

<input bind:value={_times_to_be_rendered} placeholder="How many players..."/>
<br>
<button type="submit" on:click={update}>Update</button>
<br>
{#each Array(times_to_be_rendered) as _, i}    
    <bordered>
    <article class="grid">
        <Counter />
    </article>
    <br>
    </bordered>
{/each}
</article>
