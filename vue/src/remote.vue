<template>
  <div>
    <h1 style="text-align: center">网盘模式(用户：{{ Body.user_name }})</h1>
    <hr />
    <h3>
      <p id="lbl">本地路径：{{ local }}</p>
      <p id="lbl">网盘路径：{{ server }}</p>
      <p id="lbl">备份状态：{{ back_status }}</p>
      <center>
        <button
          id="btn2"
          @click="submit('backup')"
          style="height: 35px; width: 120px; font-size: 18px"
        >
          备份到网盘
        </button>
        <button
          id="btn2"
          @click="submit('rec')"
          style="height: 35px; width: 120px; font-size: 18px"
        >
          还原到本地
        </button>
        <a href="/">
          <button id="btn2" style="height: 35px; width: 120px; font-size: 18px">
            返回首页
          </button>
        </a>
      </center>
    </h3>
    <hr />
    <div id="first">
      <h2>选择备份选项</h2>
      <center>
        <div id="column3">
          <input type="checkbox" id="ckx" value="自定义备份" v-model="b_opt" />
          <label>自定义备份</label>
          <input type="checkbox" id="ckx" value="压缩" v-model="b_opt" />
          <label>压缩</label>
          <input type="checkbox" id="ckx" value="打包" v-model="b_opt" />
          <label>打包</label>
          <input type="checkbox" id="ckx" value="加密" v-model="b_opt" />
          <label>加密</label>
          <br />
        </div>
      </center>
      <rleft @left="left"></rleft>
    </div>
    <div id="second">
      <h2>选择还原选项</h2>
      <center>
        <div id="column3">
          <input type="checkbox" id="ckx" value="压缩" v-model="b_opt" />
          <label>解压</label>
          <input type="checkbox" id="ckx" value="打包" v-model="b_opt" />
          <label>解包</label>
          <input type="checkbox" id="ckx" value="加密" v-model="b_opt" />
          <label>解密</label>
          <br />
        </div>
      </center>
      <rright @right="right"></rright>
    </div>
  </div>
</template>

<script>
import rleft from "./components/remote_left.vue";
import rright from "./components/remote_right.vue";
import axios from "axios";
export default {
  name: "remote",
  components: {
    rleft,
    rright,
  },
  data() {
    return {
      local: "",
      server: "",
      header: "http://localhost:8090/method",
      newBody: "",
      default_pth: "",
      b_opt: [],
      r_opt: [],
      back_status: "",
      Body: {
        op: "",
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
    var usrname = window.location.href.split("?")[1];
    that.Body.user_name = usrname;
  },
  methods: {
    left: function (data) {
      var that = this;
      that.local = data;
      that.back_status = "";
    },
    right: function (data) {
      var that = this;
      that.server = data;
      that.back_status = "";
    },
    Post: function (type) {
      var addr = this.header,
        data = this.newBody;
      var that = this;
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == "true") {
              that.back_status += type + "成功" + "; ";
            } else {
              window.alert(type + "失败:" + rsp.err);
              that.back_status += type + "失败" + "; ";
            }
          })
          .catch(function (error) {
            that.post_response_msg = null;
            window.alert(error);
          });
      }
    },
    submit: function (type = "") {
      var l_pth = this.local;
      var r_pth = this.server;
      var b_opt = this.b_opt,
        r_opt = this.r_opt;
      var filename;
      if (type == "backup") {
        filename = l_pth.substring(l_pth.lastIndexOf("/"));
        r_pth += filename;
      } else {
        filename = r_pth.substring(r_pth.lastIndexOf("/"));
        l_pth += filename;
      }
      if (l_pth == "" || r_pth == "") {
        window.alert("本地路径或网盘路径为空！");
        return;
      }
      var r;
      if (type == "backup") {
        r = window.confirm(
          "您要将文件(夹)：" +
            l_pth +
            "\n" +
            "备份到：" +
            r_pth +
            "\n" +
            "备份选项：" +
            b_opt +
            "\n" +
            "注意：若目标路径存在重名文件可能会被覆盖！"
        );
      } else {
        r = window.confirm(
          "您要将文件(夹)：" +
            r_pth +
            "\n" +
            "还原到：" +
            l_pth +
            "\n" +
            "还原选项：" +
            r_opt +
            "\n" +
            "注意：若目标路径存在重名文件可能会被覆盖！"
        );
      }
      var pack = 0,
        enco = 0,
        compress = 0,
        custom = 0;
      var that = this;
      if (r == true) {
        if (type == "backup") {
          that.back_status = "";
          that.newBody = that.Body;
          for (var i = 0; i < b_opt.length; i++) {
            if (b_opt[i] == "自定义备份") {
              custom = 1;
            } else if (b_opt[i] == "压缩") {
              compress = 1;
            } else if (b_opt[i] == "打包") {
              pack = 1;
            } else if (b_opt[i] == "加密") {
              enco = 1;
            }
          }
          that.newBody.op = "remote_copy";
          that.newBody.copy_para.origin_path = l_pth;
          that.newBody.copy_para.backup_path = r_pth;
          this.Post("备份");
          if (custom == 1) {
          } //TODO
          if (pack == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_pack";
            that.newBody.pack_para.is_pack = true;
            that.newBody.pack_para.pack_path = r_pth;
            this.Post("打包");
          }
          if (compress == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_compress";
            that.newBody.compress_para.is_compress = true;
            that.newBody.compress_para.compress_path = r_pth;
            this.Post("压缩");
          }
          if (enco == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_encode";
            that.newBody.encode_para.is_encode = true;
            that.newBody.encode_para.encode_para = r_pth;
            this.Post("加密");
          }
        } else {
          /*==============================还原================================= */
          that.back_status = "";
          that.newBody = that.Body;
          for (var i = 0; i < r_opt.length; i++) {
            if (r_opt[i] == "解压") {
              compress = 1;
            } else if (r_opt[i] == "解包") {
              pack = 1;
            } else if (r_opt[i] == "解密") {
              enco = 1;
            }
          }
          if (enco == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_encode";
            that.newBody.encode_para.is_encode = false;
            that.newBody.encode_para.encode_para = r_pth;
            this.Post("解密");
          }
          if (compress == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_compress";
            that.newBody.compress_para.is_compress = false;
            that.newBody.compress_para.compress_path = r_pth;
            this.Post("解压");
          }
          if (pack == 1) {
            that.newBody = that.Body;
            that.newBody.op = "remote_pack";
            that.newBody.pack_para.is_pack = false;
            that.newBody.pack_para.pack_path = r_pth;
            this.Post("解包");
          }
          that.newBody = that.Body;
          that.newBody.op = "remote_recover";
          that.newBody.recover_para.recover_path = r_pth;
          //that.newBody.recover_para. = l_pth;
          this.Post("还原");
        }
      }
    },
  },
};
</script>

<style>
@import "./backup.CSS";
#first {
  width: 49%;
  float: left;
  text-align: center;
}
#second {
  width: 49%;
  float: right;
  text-align: center;
}
#showpth {
  margin: 3px;
  margin-left: 20px;
  text-align: left;
  font-size: 18px;
}
#lbl {
  text-align: left;
  padding: 5px;
  margin: 0%;
}
</style>