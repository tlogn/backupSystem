<template>
  <div>
    <h2>网盘目录</h2>
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
          <button @click="sel_tar()" id="btn2">选择当前路径</button>
          <button @click="new_folder()" id="btn3">创建文件夹</button>
        </ul>
        <div id="for" v-for="fil in lis">
          <div v-if="fil.is_dir == true">
            <ul id="column1" style="width: 350px; padding: 8px">
              <div id="fil1">
                <button @click="sel_tar(fil.file_name)" id="btn2">选择</button>
                <button @click="sele(fil.file_name, fil.is_dir)" id="btn3">
                  进入
                </button>
                <button @click="del(fil.file_name)" id="btn4">删除</button>
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
                <button @click="sel_tar(fil.file_name)" id="btn2">选择</button>
                <button @click="del(fil.file_name)" id="btn4">删除</button>
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
import c from "../../common.vue";
axios.defaults.headers.post["content-type"] = "application/json";
export default {
  name: "rright",
  data() {
    return {
      msg: "",
      header: "http://localhost:8090/method",
      lis: "",
      curPth: "",
      default_pth: "",
      Body: c.Body,
    };
  },
  mounted: function () {
    var that = this,
      usrname = sessionStorage.getItem("user_name");
    that.default_pth = "/home/lighthouse/backup/" + usrname;
    that.curPth = that.default_pth;
    that.ini_get();
  },
  methods: {
    emitToParent: function (para) {
      this.$emit("right", para);
    },
    new_folder: function () {
      var folder_name = prompt("请命名文件夹：", "新建文件夹");
      var that = this;
      that.Body.op = "remote_mkdir";
      that.Body.dir_para.dir_path = that.curPth + folder_name;
      axios
        .post(that.header, that.Body)
        .then(function (response) {
          var rsp = response.data;
          if (rsp.succeed == false) {
            window.alert("创建失败:" + rsp.err);
          } else that.Reload();
        })
        .catch(function (error) {
          window.alert(error);
        });
    },
    Reload() {
      var that = this;
      that.Body.op = "remote_dir";
      that.Body.dir_para.dir_path = that.curPth;
      axios
        .get(that.header, {
          params: {
            body: that.Body,
          },
        })
        .then((res) => {
          that.lis = res.data.dir_files;
        });
    },
    del: function (filename) {
      var r = confirm("确定删除 " + this.curPth + filename + " ? 不可恢复！");
      if (r == false) return;
      var that = this;
      that.Body.op = "remote_remove";
      that.Body.dir_para.dir_path = that.curPth + filename;
      axios
        .post(that.header, that.Body)
        .then(function (response) {
          var rsp = response.data;
          if (rsp.succeed == false) window.alert("删除失败：" + rsp.err);
          else {
            that.Reload();
          }
        })
        .catch(function (error) {
          window.alert(error);
        });
    },
    ini_get: function (para = this.default_pth) {
      var that = this;
      that.curPth = para + "/";
      that.Body.op = "remote_dir";
      that.Body.dir_para.dir_path = this.curPth;
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
            pth = pth.substring(0, pth.lastIndexOf("/"));
            pth = pth.substring(0, pth.lastIndexOf("/") + 1);
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
      if (filename != "") pth = this.curPth + filename;
      else {
        pth = this.curPth;
        if (pth[pth.length - 1] == "/")
          pth = pth.substring(0, pth.lastIndexOf("/"));
      }
      window.scrollTo(0, -50);
      this.emitToParent(pth);
    },
  },
};
</script>   