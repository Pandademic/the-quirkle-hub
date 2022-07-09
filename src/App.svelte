<script>
  import Counter from './components/Counter.svelte'
  import Winners from './components/Winners.svelte'
  import Title from './components/Title.svelte'
  import About from './components/About.svelte'
  import Rules from './components/Rules.svelte'

  import "@picocss/pico";
</script>

<main>
  <Title />
  <br>
  <Counter />
  <br>
  <Winners />
  <br>
  <Rules />
  <br>
  <About />
</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
  }
</style>
