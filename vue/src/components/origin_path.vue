<template>
  <div >
    <h2>选择要备份的文件</h2>
    <input
      placeholder="请输入目录，如'C:'或'Users'。注意不带斜杠"
      v-model="root"
      style="height: 23px; width: 400px; font-size: 18px"
    />
    <br /><br />
    <div>
      <button id="btn2" @click="ini_get(root)">浏览文件</button>
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
        <div id="for" v-for="fil in lis">
          <div v-if="fil.is_dir == true">
            <ul id="column1" style="width: 350px; padding: 8px">
              <div id="fil1">
                <button
                  @click="sel_ori(fil.file_name)"
                  id="btn2"
                >
                  选择
                </button>
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
                <button
                  @click="sel_ori(fil.file_name)"
                  id="btn2"
                >
                  选择
                </button>
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
  name: "Backup",
  data() {
    return {
      msg: "",
      header: "http://localhost:8090/method",
      newBody: null,
      lis: "",
      Body: c.Body,
      curPth: "",
    };
  },
  methods: {
    emitToParent: function(para) {
      this.$emit('ori', para);
    },
    ini_get: function (para = "/mnt/d/123/0Bachelor/大四上/软件开发实验") {
      console.log(this.Body.user_name);
      this.curPth = para + "/";
      var that = this;
      that.Body.op = "local_dir";
      that.Body.dir_para.dir_path = this.curPth;
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
      var pth = this.curPth;
      if (pth.length == 0) {
        return;
      }
      pth = pth.substring(0, pth.lastIndexOf("/"));
      pth = pth.substring(0, pth.lastIndexOf("/"));
      this.curPth = pth;
      this.ini_get(this.curPth);
    },
    sel_ori: function(filename) {
      var oripth = this.curPth + filename;
      window.scrollTo(0, -50);
      this.emitToParent(oripth);
    }
  },
};
</script>