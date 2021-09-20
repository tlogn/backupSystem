<template>
  <div id="one">
    <h1>TEST</h1>
    <hr />
    <br />
    <div id="first">
      <textarea
        v-model="post_params"
        placeholder="请输入POST参数..."
        style="height: 300px; width: 400px; font-size: 18px"
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
        style="height: 300px; width: 400px; font-size: 18px"
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
    <hr>
    <br />
    <p>
      {<br />
      "op":"local_dir",<br />

      "dir_para":{<br />
      "dir_path":"D:\"<br />
      },<br />

      "copy_para":{<br />
      "origin_path":"",<br />
      "backup_path":""<br />
      },<br />

      "recover_para":{<br />
      "recover_path":""<br />
      },<br />

      "compress_para":{<br />
      "is_compress":false,<br />
      "compress_path":""<br />
      },<br />

      "encode_para":{<br />
      "is_encode":false,<br />
      "encode_path":""<br />
      },<br />

      "pack_para":{<br />
      "is_pack":false,<br />
      "pack_path":""<br />
      }<br />
      }<br />
    </p>
  </div>
</template>

<script>
import axios from "axios";
axios.defaults.headers.post["content-type"] = "application/json";
export default {
  name: "one",
  data() {
    return {
      msg: "Backup",
      post_response_msg: "",
      get_response_msg: "",
      post_status: "",
      get_status: "",
      json_msg: ""
    };
  },
  methods: {
    post: function (addr, param) {
      
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        var that = this;
        //var data = '{' + '"' + 'body' + '"' + ':' + param + '}'

        var data = param;
        axios
          .post(addr, data)
          .then(function (response) {
            that.post_response_msg = response.data;
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
            that.get_response_msg = response.data
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
  width: 50%;
  float: left;
}
#second {
  width: 50%;
  float: right;
}
</style>
