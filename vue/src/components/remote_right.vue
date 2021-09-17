<template>
  <div>
    <h2>网盘目录</h2>
    <div>
      <button id="btn2" @click="ini_get()">浏览文件</button>
      <br />
      <h4>当前路径：{{ Body.get_dir_para.dir_path }}</h4>
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
axios.defaults.headers.post["content-type"] = "application/json";
export default {
  name: "rright",
  data() {
    return {
      msg: "",
      header: "http://localhost:8090/method",
      lis: "",
      newBody: null,
      default_pth: "",
      Body: {
        op: "remote_dir",
        get_dir_para: {
          dir_path: "",
        },

        user_name: "",

        login_para: {
          username: "",
          password: "",
        },

        copy_para: {
          origin_path: "",
          backup_path: "",
        },

        recover_para: {
          recover_path: "",
        },

        compress_para: {
          is_compress: false,
          compress_path: "",
        },

        encode_para: {
          is_encode: false,
          encode_path: "",
        },

        pack_para: {
          is_pack: false,
          pack_path: "",
        },
      },
    };
  },
  mounted: function () {
    var that = this;
    var usrname = window.location.href.split('?')[1];
    that.default_pth = "/home/lighthouse/backup/" + usrname;
    that.Body.user_name = usrname;
    that.Body.get_dir_para.dir_path = that.default_pth;
  },
  methods: {
    emitToParent: function (para) {
      this.$emit("right", para);
    },
    new_folder: function() {},
    ini_get: function (para = this.default_pth) {
      var that = this;
      that.Body.get_dir_para.dir_path = para + "/";
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
          }
        })
        .catch(function (error) {
          that.get_response_msg = null;
          window.alert(error);
        });
    },
    sele: function (para, type) {
      if (type == true) {
        var pth = this.Body.get_dir_para.dir_path + para;
        this.ini_get(pth);
      }
    },
    Return: function () {
      var that = this;
      var pth = that.Body.get_dir_para.dir_path;
      if (pth == that.default_pth + "/" || pth == that.default_pth) return;
      if (pth.length == 4) {
        that.Body.get_dir_para.dir_path = "";
        return;
      } else if (pth.length == 0) {
        return;
      }
      pth = pth.substring(0, pth.lastIndexOf("/"));
      pth = pth.substring(0, pth.lastIndexOf("/"));
      that.Body.get_dir_para.dir_path = pth;
      this.ini_get(that.Body.get_dir_para.dir_path);
    },
    sel_tar: function (filename = "") {
      var pth;
      if (filename!="")
        pth = this.Body.get_dir_para.dir_path + filename;
      else {
        pth = this.Body.get_dir_para.dir_path;
        if (pth[pth.length-1] == '/')
          pth = pth.substring(0, pth.lastIndexOf('/'));
      }
      window.scrollTo(0, -50);
      this.emitToParent(pth);
    },
  },
};
</script>   