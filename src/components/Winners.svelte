<script>
    import Icon from '@iconify/svelte';

    let winners = [
        {name: "This is a example",points: "2",wonBy: "1"}
    ]

    let name = ""
    let points = ""
    let wonBy = ""

    const addWinner = () => {
        winners = [
            ...winners,
            {
                name,
                points,
                wonBy
            }
        ]
        name = ""
        points = ""
        wonBy = ""
        console.log(winners)
        
    }
</script>
<article>
    <h1>Winners:</h1>
    <h6>First at top , latest at bottom.</h6>
    <br>
    <form>
        <label for="name">Name</label>
        <input id="name" type="text" bind:value={name} />

        <label for="name">Points</label>
        <input id="name" type="text" bind:value={points} />
    
        <label for="name">Won By (to nearest competitor)</label>
        <input id="name" type="text" bind:value={wonBy} />
    </form>
    <br>
    <button type="submit" on:click={addWinner}>Submit!</button>
    <br>
    <br>
    {#each winners as winner}
      <article>
        <span><Icon icon="emojione:1st-place-medal" /> {winner.name}</span>
        <br>
        <span><Icon icon="openmoji:hundred-points" /> Total Points: {winner.points}</span>
        <br>
        <span><Icon icon="fxemoji:trophy" /> Beat the nearest competitor by: {winner.wonBy}</span>
      </article>
    {/each}
</article>