<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>Essential functions</h2>
    <ul>
      <div id="backUp">
        <li>
          <button @click="sendBackUp">backUp</button>
        </li>
      </div>
      <br />
      <li>
        <button href="https://forum.vuejs.org" target="_blank">reStore</button>
      </li>
    </ul>
    <hr />
    <h2>Advanced functions</h2>
    
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  data() {
    return {
      msg: "Welcome to backupSystem",
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  /*display: inline-block;*/
  margin: 0 20px;
}
a {
  color: #42b983;
}
</style>
