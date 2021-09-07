<template>
  <div id="one">
    <h1>Backup</h1>
    <br />
    <textarea
      v-model="params"
      placeholder="请输入参数..."
      style="height: 80px; width: 400px; font-size: 18px"
    />
    <br />
    <p style="font-size:20px">http://localhost:8090/method</p>
    <input
      v-model="addr"
      placeholder="请输入URL..."
      style="height: 19px; width: 400px; font-size: 18px"
    />
    <button
      @click="post(addr, params)"
      style="height: 28px; width: 60px; font-size: 18px"
    >
      提交
    </button>
    <br><br>
    <hr>
    <p style="font-size:20px">reponse = {{ response_msg }}</p>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "one",
  data() {
    return {
      msg: "Backup",
      response_msg: null
    };
  },
  methods: {
    post: function (addr, params) {
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        var that = this
        axios
          .post(addr, { params })
          .then(function (response) {
            that.response_msg = response
          })
          .catch(function (error) {
            that.response_msg = null
            window.alert(error)
          });
      }
    },
  },
};
</script>

<style>
#one {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2e659b;
  margin-top: 100px;
}
</style>
