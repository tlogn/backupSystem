<template>
  <div id="sel_back">
    <h2>选择备份目标目录</h2>
    默认备份目录(测试用)：
    <input
      placeholder="请输入默认备份目录"
      v-model="default_pth"
      style="height: 23px; width: 400px; font-size: 18px"
    />
    <br /><br />
    <div>
      <button id="btn2" @click="ini_get()">浏览文件</button>
      <br />
      <h4>当前路径：{{ curPth }}</h4>
      <button
        id="btn2"
        @click="Return()"
        style="height: 50px; width: 120px; font-size: 18px"
      >
        返回上一级
      </button>
    </div>
    <br />
    <div id="list">
      <center>
        <ul id="column1" style="width: 350px; padding: 8px">
        <button @click="sel_tar()" id="btn2" >选择当前路径</button>
        </ul>
        <div id="for" v-for="fil in lis">
          <div v-if="fil.is_dir == true">
            <ul id="column1" style="width: 350px; padding: 8px">
              <div id="fil1">
                <button @click="sel_tar(fil.file_name)" id="btn2">选择</button>
                <button @click="sele(fil.file_name, fil.is_dir)" id="btn3">
                  进入
                </button>
                <label style="font-size: 18px">
                  {{ fil.file_name }}
                </label>
              </div>
            </ul>
          </div>
        </div>
        <div id="for" v-for="fil in lis">
          <div v-if="fil.is_dir != true">
            <ul id="column2" style="width: 350px; padding: 8px">
              <div id="fil2">
                <label style="font-size: 18px">
                  {{ fil.file_name }}
                </label>
              </div>
            </ul>
          </div>
        </div>
      </center>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
import c from "../../common.vue"
axios.defaults.headers.post["content-type"] = "application/json";
export default {
  name: "Target",
  data() {
    return {
      msg: "",
      header: "http://localhost:8090/method",
      lis: "",
      newBody: null,
      default_pth: "",
      Body: c.Body,
      curPth: "",
    };
  },
  mounted: function () {
    this.get_os_type();
    var that = this;
    that.curPth = that.default_pth;
  },
  methods: {
    get_os_type: function () {
      var that = this;
      var type = navigator.userAgent.toLowerCase();
      if (type.indexOf("win") > -1) {
        type = "win";
        that.default_pth = "/mnt/d/123/0Bachelor/大四上/软件开发实验/backup";
      } else if (type.indexOf("mac") > -1) {
        type = "mac";
        that.default_pth = "/users/backup";
      } else if (type.indexOf("linux") > -1) {
        type = "linux";
        that.default_pth = "/mnt/d/123";
      } else type = "unknown";
    },
    emitToParent: function (para) {
      this.$emit("tar", para);
    },
    ini_get: function (para = this.default_pth) {
      var that = this;
      that.curPth = para + "/";
      that.Body.op = "local_dir";
      that.Body.dir_para.dir_path = that.curPth;
      //window.alert(that.Body.dir_para.dir_path);
      axios
        .get(that.header, {
          params: {
            body: that.Body,
          },
        })
        .then(function (response) {
          var data = response.data;
          if (data.succeed == true) {
            that.lis = data.dir_files;
          } else {
            window.alert(data.err);
            var pth = that.curPth;
            pth = pth.substring(0, pth.lastIndexOf('/'));
            pth = pth.substring(0, pth.lastIndexOf('/')+1);
            that.curPth = pth;
          }
        })
        .catch(function (error) {
          that.get_response_msg = null;
          window.alert(error);
        });
    },
    sele: function (para, type) {
      if (type == true) {
        var pth = this.curPth + para;
        this.ini_get(pth);
      }
    },
    Return: function () {
      var that = this;
      var pth = that.curPth;
      if (pth == that.default_pth + "/" || pth == that.default_pth) return;
      if (pth.length == 0) {
        return;
      }
      pth = pth.substring(0, pth.lastIndexOf("/"));
      pth = pth.substring(0, pth.lastIndexOf("/"));
      that.curPth = pth;
      this.ini_get(that.curPth);
    },
    sel_tar: function (filename = "") {
      var pth;
      if (filename!="")
        pth = this.curPth + filename;
      else {
        pth = this.curPth;
        if (pth[pth.length-1] == '/')
          pth = pth.substring(0, pth.lastIndexOf('/'));
      }
      window.scrollTo(0, -50);
      this.emitToParent(pth);
    },
  },
};
</script>   