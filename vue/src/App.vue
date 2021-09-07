<template>
  <div id="background">
    <div
      v-if="img"
      :style="{
        backgroundImage: './assets/background.jpg',
        backgroundSize: cover,
      }"
    ></div>
    <map name="backup">
      <area
        shape="rect"
        coords="0, 0, 300, 120"
        href="./one.html"
        alt="click_backup"
        title="备份"
      />
    </map>
    <map name="restore">
      <area
        shape="rect"
        coords="0, 0, 300, 120"
        href="./two.html"
        alt="click_restore"
        title="恢复"
      />
    </map>
    <div id="button">
      <router-link to="./backup">
      <img
        src="./assets/button1.png"
        style="height: 120px; width: 300px"
        alt="backup"
        usemap="#backup"
      />
      </router-link>
      <img
        src="./assets/button2.png"
        style="height: 120px; width: 300px"
        alt="restore"
        usemap="#restore"
      />
      <br />
      {{ messege }}
    </div>
    <div id="http_get">
      <br />
      <textarea
        style="height: 60px; width: 300px; font=size: 18px"
        placeholder="编辑要发送的内容..."
        v-model="http_body"
      />
      <p>输入字段:(http://localhost:8090/method)</p>
      <input
        v-model="http_addr"
        placeholder="输入..."
        style="height: 24px; width: 400px; font-size: 21px"
      />
      <button
        style="height: 25px"
        @click="http_get_method(http_addr, http_body)"
      >
        GET
      </button>
      <p>responseCode = {{ http_status_code }}</p>
      <hr />
      <p>responseMsg = {{ http_msg }}</p>
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
      http_status_code: null,
      http_msg: null,
    };
  },
  methods: {
    http_get_method: function (addr, body) {
      if (addr == null) {
        window.alert("Empty addr");
      } else {
        var that = this;
        axios
          .get(addr, {
            params: {
              body,
            },
          })
          .then(function (response) {
            that.http_status_code = response.status;
            that.http_msg = response;
          })
          .catch(function (error) {
            window.alert(error);
            that.http_info = null;
          });
      }
    },
  },
};
</script>

<style>
#button {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2e659b;
  margin-top: 100px;
}
#http_get {
  text-align: center;
}
#background {
  background: url('./assets/background.jpg') center center no-repeat;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%; /**宽高100%是为了图片铺满屏幕 */
  z-index: -1;
  position: absolute;
}
</style>
