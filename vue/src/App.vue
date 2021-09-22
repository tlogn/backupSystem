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
      <div>
        <center>
          <a href="./one.html"> <button>测试页面</button> </a><br /><br />
          <a href="./two.html">
            <button
              id="btn1"
              style="font-size: 25px; height: 70px; width: 130px; cursor: hand"
              title="备份"
            >
              备份
            </button>
          </a>
          <a href="./thr.html">
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
      </div>
      <div id="reg">
        <button id="but" style="cursor:pointer" @click="login(u,p)">登录</button>  
        <button id="but" style="cursor:pointer" @click="reg(u,p)">注册</button> 
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import c from "../common.vue"
export default {
  data() {
    return {
      messege: "",
      header: "http://localhost:8090/method",
      Body:c.Body,
      u: "",
      p: "",
      vis: false,
      src: "../../../1.课程概述与需求分析.pdf",
    };
  },
  methods: {
    login: function (u, p) {
      if (u.length == 0 || p.length == 0) {
        window.alert("用户名或密码为空！");
        return;
      }
      var that = this;
      that.Body.op = "login";
      that.Body.login_para.username = u;
      that.Body.login_para.password = p;
      this.Post("登陆");
    },
    reg: function (u, p) {
      if (u.length == 0 || p.length == 0) {
        window.alert("用户名或密码为空！");
        return;
      }
      var r = window.confirm("是否确定注册用户名为" + u + "的新用户?");
      if (r == true) {
        var that = this;
        that.Body.op = "register";
        that.Body.login_para.username = u;
        that.Body.login_para.password = p;
        this.Post("注册");
      }
    },
    Post: function (type) {
      var addr = this.header,
        data = this.Body;
      var that = this;
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == false) { //login/reg fail
              var err = rsp.err;
              if (type == "登陆") { //login
                if (err == "wrong password") window.alert("密码错误！");
                else if (err == "username not exists")
                  window.alert("用户不存在！请先注册");
                else window.alert("未知错误:"+err);
              } else {  //reg
                if (err == "username already exists")
                  window.alert("注册失败：用户名已被注册");
                else  
                  window.alert("未知错误:"+err);
              }
            } else {  //login/reg succeed
              if (type == "登陆") { //login
                sessionStorage.setItem("user_name",that.Body.login_para.username);
                //console.log(that.Body.user_name)
                window.location.href =
                  "./remote.html?" + that.Body.login_para.username;
              } else {  //reg
                window.alert( 
                  "注册成功！您的用户名为" + that.Body.login_para.username
                );
              }
            }
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
  height: 150px;
}
#login {
  color: rgb(197, 8, 235);
  text-shadow: 0 0 2px;
  letter-spacing: 1px;
  text-align: center;
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
#but {
  width: 150px;
  min-height: 20px;
  background-color: #d851bb;
  border: 2px solid #3762bc;
  color: #fff;
  padding: 9px 14px;
  font-size: 15px;
  line-height: normal;
  border-radius: 5px;
  margin: 0;
}
#reg {
  position: relative;
  padding: 10px;
  top: 40%;
}
</style>
