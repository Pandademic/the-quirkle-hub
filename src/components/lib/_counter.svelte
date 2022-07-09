<script>
  let count = 0
  let incrementer = 0

  const increment = () => {
    // @ts-ignore
    count += parseInt(incrementer)
    incrementer = 0
  }

  const reset = () => {
    count = 0
    incrementer = 0
  }

</script>

<button on:click={increment} data-tooltip="Tap me!" class="btn-green">
  Count: {count} (Click to add points)
</button>
<input bind:value={incrementer} placeholder="Add how many points..."/>
<button on:click={reset} data-tooltip="Be careful...." class="btn-red">
  Reset
</button>
<style>
  button {
    font-family: inherit;
    font-size: inherit;
    padding: 1em 2em;
    border-radius: 2em;
    border: 2px solid rgba(255, 62, 0, 0);
    font-variant-numeric: tabular-nums;
    cursor: pointer;
  }
  .btn-red {
    background-color:  #E53E3E;
  }
  .btn-green {
    background-color: #68D391;
  }
</style>
