<template>
  <div>
    <div id="app">
      <img src="./assets/logo.png" />
      <br />
      {{ messege }}
    </div>
    <div id="http_get">
      <br>
      <textarea 
        style="height:60px; width:300px; font=size:18px" 
        placeholder="编辑要发送的内容..."
        v-model="http_body"
      />
      <p>输入字段:(http://localhost:8090/method)</p>
      <input
        v-model="http_addr"
        placeholder="输入..."
        style="height: 24px; width: 400px; font-size: 21px"
      />
      <button style="height: 25px" @click="http_get_method(http_addr, http_body)">
        GET
      </button>
      <p id="http_return_msg">NULL</p>
    </div>
  </div>
</template>

<script src="https://cdn.staticfile.org/vue/2.4.2/vue.min.js"></script>
<script src="https://cdn.staticfile.org/axios/0.18.0/axios.min.js"></script>

<script>
import axios from "axios";
export default {
  data() {
    return {
      messege: "HTTP_GET_TEST",
      http_info: null,
    };
  },
  methods: {
    http_get_method: function (addr, body) {
      if (addr == null) {
        window.alert("Empty inputArea");
      } else {
        axios
          .get(addr, {
            params:{
              body
            }
          })
          .then(function (response) {
            window.alert("responseCode=" + response.status);
          })
          .catch(function (error) {
            window.alert(error);
          });
      }
    },
  },
};
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
