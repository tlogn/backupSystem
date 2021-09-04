<template>
  <div>
    <div id="app">
      <img src="./assets/logo.png" />
      <br/>
      {{messege}}
    </div>
    <div id="http_get">
      <p>输入字段:</p>
      <input v-model="http_get_messege" placeholder="输入..." />
      <button @click="http_get_method(http_info)">GET</button>
      <p>输入的信息为:{{ http_get_messege }}</p>
    </div>
  </div>
</template>

<script src="https://cdn.staticfile.org/vue/2.4.2/vue.min.js"></script>
<script src="https://cdn.staticfile.org/axios/0.18.0/axios.min.js"></script>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      messege:"abc",
      http_info:null
    }
  },
  methods: {
    http_get_method: function(info) {
      window.alert("GOT:"+info)
    }
  },
  mounted () {
    axios
      .get('https://www.runoob.com/try/ajax/json_demo.json')
      .then(response => (this.http_info = response))
      .catch(function (error) { 
        console.log(error);
      });
  }
}
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2e659b;
  margin-top: 60px;
}
#http_get {
  text-align: center;
}
</style>
