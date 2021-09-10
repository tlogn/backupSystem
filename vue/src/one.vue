<template>
  <div id="one">
    <h1>TEST</h1>
    <hr />
    <br />
    <div id="first">
      <textarea
        v-model="post_params"
        placeholder="请输入POST参数..."
        style="height: 80px; width: 400px; font-size: 18px"
      />
      <br />
      <p style="font-size: 20px">http://localhost:8090/method</p>
      <input
        v-model="post_addr"
        placeholder="请输入URL..."
        style="height: 19px; width: 400px; font-size: 18px"
      />
      <button
        @click="post(post_addr, post_params)"
        style="height: 28px; width: 65px; font-size: 18px"
      >
        POST
      </button>
      <br />
      <p style="font-size: 20px">status = {{ post_status }}</p>
      <p style="font-size: 20px">reponse = {{ post_response_msg }}</p>
    </div>
    <div id="second">
      <textarea
        v-model="get_params"
        placeholder="请输入GET参数..."
        style="height: 80px; width: 400px; font-size: 18px"
      />
      <br />
      <p style="font-size: 20px">http://localhost:8090/method</p>
      <input
        v-model="get_addr"
        placeholder="请输入URL..."
        style="height: 19px; width: 400px; font-size: 18px"
      />
      <button
        @click="get(get_addr, get_params)"
        style="height: 28px; width: 65px; font-size: 18px"
      >
        GET
      </button>
      <br />
      <p style="font-size: 20px">status = {{ get_status }}</p>
      <p style="font-size: 20px">reponse = {{ get_response_msg }}</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "one",
  data() {
    return {
      msg: "Backup",
      post_response_msg: null,
      get_reponse_msg: null,
      post_status: null,
      get_status: null,
    };
  },
  methods: {
    post: function (addr, param) {
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        var that = this;
        axios
          .post(addr, param
          )
          .then(function (response) {
            that.post_response_msg = response;
            that.post_status = response.status;
          })
          .catch(function (error) {
            that.post_response_msg = null;
            window.alert(error);
          });
      }
    },
    get: function (addr, body) {
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        var that = this;
        axios
          .get(addr, {
            params: {
              body,
            },
          })
          .then(function (response) {
            that.get_response_msg = response;
            that.get_status = response.status;
          })
          .catch(function (error) {
            that.get_response_msg = null;
            window.alert(error);
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
  margin-top: 40px;
}
#first {
  width: 48%;
  float: left;
}
#second {
  width: 48%;
  float: right;
  /*height: 100px;
  border: 1px solid #3b6273;*/
}
</style>
