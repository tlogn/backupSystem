<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
<script src='https://unpkg.com/vue-router/dist/vue-router.js'></script>

<template>
  <div id="background">
    <div
      v-if="img"
      :style="{
        backgroundImage: './assets/background.jpg',
        backgroundSize: cover,
      }"
    ></div>
    <div id="first">
      <br>
      <h1>本地模式</h1>
      <div id="button1" style="flaot: left" align="left">
        <center>
          <a href="./../one.html"> <button>测试页面</button> </a><br /><br />
          <a href="./../two.html">
            <button
              id="btn1"
              style="font-size: 25px; height: 70px; width: 130px; cursor: hand"
              title="备份"
            >
              备份
            </button>
          </a>
          <a href="./../thr.html">
            <button
              id="btn1"
              style="font-size: 25px; height: 70px; width: 130px"
              title="还原"
            >
              还原
            </button>
          </a>
        </center>
      </div>
    </div>
    <div id="second">
      <h1>网盘模式</h1>
      <div id="login">  
        <h1>Login</h1>  
        <input type="text" required="required" placeholder="用户名" v-model="u" autocomplete="off"></input>  
        <input type="password" required="required" placeholder="密码" v-model="p"></input>  
        <button class="but" type="submit" style="cursor:pointer" @click="login(u,p)">登录</button>  
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messege: "",
    };
  },
  methods: {
    login: function (u, p) {
      console.log(u, p);
      window.location.href="./../one.html";
    },
  },
};
</script>

<style>
#background {
  background: url("./assets/bkg.jpeg") center center no-repeat;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%; /*宽高100%是为了图片铺满屏幕 */
  z-index: -1;
  position: absolute;
  background-size: cover;
}
#first {
  height: 48%;
  margin: 0;
  text-align: center;
}
#second {
  height: 48%;
  margin: 0;
  text-align: center;
}
#btn1 {
  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
  padding: 16px 32px;
  text-align: center;
  background-color: rgb(234, 211, 240);
  color: black;
  border: 4px solid rgb(202, 108, 240);
  border-radius: 5px;
}
#btn1:hover {
  background-color: #e67ecc;
  color: rgba(247, 240, 244, 0.925);
  cursor: pointer;
}
button {
  display: inline-block;
  margin-right: 20px;
  margin-left: 20px;
}
#login {
  position: relative;
  top: 40%;
  left: 50%;
  margin: -150px 0 0 -150px;
  width: 300px;
  height: 300px;
}
#login h1 {
  color: rgb(197, 8, 235);
  text-shadow: 0 0 2px;
  letter-spacing: 1px;
  text-align: center;
}
h1 {
  font-size: 2em;
  margin: 0.67em 0;
}
input {
  width: 278px;
  height: 18px;
  margin-bottom: 10px;
  outline: none;
  padding: 10px;
  font-size: 13px;
  color: #fff;
  border-top: 1px solid #312e3d;
  border-left: 1px solid #312e3d;
  border-right: 1px solid #312e3d;
  border-bottom: 1px solid #56536a;
  border-radius: 4px;
  background-color: #2d2d3f;
}
.but {
  width: 300px;
  min-height: 20px;
  display: block;
  background-color: #d851bb;
  border: 2px solid #3762bc;
  color: #fff;
  padding: 9px 14px;
  font-size: 15px;
  line-height: normal;
  border-radius: 5px;
  margin: 0;
}
</style>
