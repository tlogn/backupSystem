<template>
  <div>
    <div v-show="showPop">
      <pop></pop>
    </div>
    <h1 style="text-align: center">网盘模式(用户：{{ curUser }})</h1>
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
          <input type="checkbox" id="ckx" value="压缩" v-model="b_opt" />
          <label>压缩</label>
          <input type="checkbox" id="ckx" value="打包" v-model="b_opt" />
          <label>打包</label>
          <br />
          <input type="checkbox" id="ckx" value="加密" v-model="b_opt" />
          <label
            >加密
            <input
              v-model="encode_pwd"
              type="password"
              v-show="IsShow"
              placeholder="请输入密码"
            />
          </label>
          <br />
        </div>
      </center>
      <rleft @left="left"></rleft>
    </div>
    <div id="second">
      <h2>选择还原选项</h2>
      <center>
        <div id="column3">
          <input type="checkbox" id="ckx" value="解压" v-model="r_opt" />
          <label>解压</label>
          <input type="checkbox" id="ckx" value="解包" v-model="r_opt" />
          <label>解包</label>
          <br />
          <input type="checkbox" id="ckx" value="解密" v-model="r_opt" />
          <label
            >解密
            <input
              v-model="decode_pwd"
              type="password"
              v-show="IsShow"
              placeholder="请输入密码"
            />
          </label>
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
import c from "../common.vue";
import pop from "./components/pop.vue";
import axios from "axios";
export default {
  name: "remote",
  components: {
    rleft,
    rright,
    pop,
  },
  data() {
    return {
      local: "",
      server: "",
      header: "http://localhost:8090/method",
      default_pth: "",
      b_opt: [],
      r_opt: [],
      back_status: "",
      Body: c.Body,
      curUser: "",
      upload: false,
      encode: false,
      compress: false,
      pack_suc: false,
      download: false,
      decode: false,
      decompress: false,
      unpack: false,
      encode_pwd: "",
      decode_pwd: "",
      IsShow: true,
      showPop: false,
      l_pth: "",
      r_pth: "",
    };
  },
  mounted: function () {
    var that = this;
    var cur_usrname = window.location.href.split("?")[1];
    var usrname = sessionStorage.getItem("user_name");
    //console.log(usrname);
    if (usrname != cur_usrname) {
      window.location.href = "/";
    } else {
      that.Body.user_name = cur_usrname;
      that.curUser = cur_usrname;
    }
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
    Post: async function (type) {
      console.log("pot" + type);
      var addr = this.header,
        data = this.Body;
      var that = this;
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        await axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == true) {
              that.back_status += type + "成功" + "; ";
            } else {
              that.back_status += type + "失败：" + rsp.err + ";" + "\n ";
              window.alert(type + "失败:" + rsp.err);
              throw type + "失败:" + rsp.err;
            }
          })
          .catch(function (error) {
            throw error;
          });
      }
    },
    submit: async function (type = "") {
      var that = this;
      that.l_pth = this.local;
      that.r_pth = this.server;
      var b_opt = this.b_opt,
        r_opt = this.r_opt;
      var filename;
      if (that.l_pth == "" || that.r_pth == "") {
          window.alert("本地路径或网盘路径为空！");
          return;
      }
      if (type == "backup") {
        filename = that.l_pth.substring(that.l_pth.lastIndexOf("/"));
        that.r_pth += filename;
      } else {
        filename = that.r_pth.substring(that.r_pth.lastIndexOf("/"));
        that.l_pth += filename;
      }
      var r;
      if (type == "backup") {
        r = window.confirm(
          "您要将文件(夹)：" +
            that.l_pth +
            "\n" +
            "上传到：" +
            that.r_pth +
            "\n" +
            "附加选项：" +
            b_opt +
            "\n" +
            "注意：若目标路径存在重名文件可能会被覆盖！"
        );
      } else {
        r = window.confirm(
          "您要将文件(夹)：" +
            that.r_pth +
            "\n" +
            "下载到：" +
            that.l_pth +
            "\n" +
            "附加选项：" +
            r_opt +
            "\n" +
            "注意：若目标路径存在重名文件可能会被覆盖！"
        );
      }
      if (r == true) {
        that.upload = false;
        that.compress = false;
        that.pack = false;
        that.encode = false;
        that.download = false;
        that.decompress = false;
        that.unpack = false;
        that.decode = false;
        if (type == "backup") {
          that.back_status = "";
          for (var i = 0; i < b_opt.length; i++) {
            if (b_opt[i] == "自定义备份") {
            } else if (b_opt[i] == "压缩") {
              that.compress = true;
            } else if (b_opt[i] == "打包") {
              that.pack = true;
            } else if (b_opt[i] == "加密") {
              that.encode = true;
              if (this.encode_pwd.length == 0) {
                window.alert("密码为空！");
                return;
              }
            }
          }
          that.showPop = true;
          await this.Encode().catch((err) => {
            that.showPop = false;
          });
          that.showPop = false;
        } else {
          that.back_status = "";
          for (var i = 0; i < r_opt.length; i++) {
            if (r_opt[i] == "解压") {
              that.decompress = true;
            } else if (r_opt[i] == "解包") {
              that.unpack = true;
            } else if (r_opt[i] == "解密") {
              that.decode = true;
            }
          }
          that.showPop = true;
          await this.Download().catch((err) => {
            that.showPop = false;
            return;
          });
          that.showPop = false;
          that.Body.op = "remote_dir";
          that.Body.dir_para.dir_path = that.r_pth;
          axios.get(that.header, {
            params: {
              body: that.Body,
            },
          });
        }
      }
    },
    async Remove(pth) {
      console.log(pth)
      var that = this;
      that.Body.op = "remote_remove";
      that.Body.dir_para.dir_path = pth;
      await axios.post(that.header, that.Body);
    },
    /************************************上传 ************************************/
    async Upload() {
      var that = this;
      that.Body.op = "remote_upload";
      that.Body.trans_para.local_path = that.l_pth;
      that.Body.trans_para.remote_path = that.r_pth;
      await this.Post("上传").catch((err) => {
        console.log("upload caught");
        throw err;
      });
    },
    async Pack() {
      var that = this;
      if (!that.pack) {
        await that.Upload().catch((err) => {
          console.log("caught upload");
          throw err;
        });
        console.log("not pack");
        return;
      }
      await that.Upload().catch((err) => {
        console.log("caught upload");
        throw err;
      });
      that.Body.op = "remote_pack";
      that.Body.pack_para.is_pack = true;
      that.Body.pack_para.pack_path = that.r_pth;
      await this.Post("打包").catch((err) => {
        throw err;
      });
      await this.Remove(that.r_pth);
      that.r_pth += ".pack";
    },
    async Compress() {
      var that = this;
      if (!that.compress) {
        await that.Pack().catch((err) => {
          console.log("caught pack");
          throw err;
        });
        console.log("not compress");
        return;
      }
      await that.Pack().catch((err) => {
        console.log("caught pack");
        throw err;
      });
      that.Body.op = "remote_compress";
      that.Body.compress_para.is_compress = true;
      that.Body.compress_para.compress_path = that.r_pth;
      await this.Post("压缩").catch((err) => {
        throw err;
      });
      await this.Remove(that.r_pth);
      that.r_pth += ".ylx";
    },
    async Encode() {
      var that = this;
      if (!that.encode) {
        await that.Compress().catch((err) => {
          console.log("caught comopress");
          throw err;
        });
        console.log("not encode");
        return;
      }
      await that.Compress().catch((err) => {
        console.log("caught comopress");
        throw err;
      });
      that.Body.op = "remote_encode";
      that.Body.encode_para.is_encode = true;
      that.Body.encode_para.encode_path = that.r_pth;
      that.Body.encode_para.password = that.encode_pwd;
      await this.Post("加密").catch((err) => {
        throw err;
      });
      await this.Remove(that.r_pth);
      that.r_pth += ".lock"
    },
    /***********************************下载 ************************************/
    async Download() {
      var that = this;
      await that.Unpack().catch((err) => {
        throw err;
      });
      that.Body.op = "remote_download";
      that.Body.trans_para.remote_path = that.r_pth;
      that.Body.trans_para.local_path = that.l_pth;
      await this.Post("下载").catch((err) => {
        console.log("download caught");
        throw err;
      });
    },
    async Unpack() {
      var that = this;
      if (!that.unpack) {
        await that.Decompress().catch((err) => {
          throw err;
        });
        return;
      }
      await that.Decompress().catch((err) => {
        console.log("unpac");
        throw err;
      });
      that.Body.op = "remote_pack";
      that.Body.pack_para.is_pack = false;
      that.Body.pack_para.pack_path = that.r_pth;
      await this.Post("解包").catch((err) => {
        console.log("unpack caught");
        throw err;
      });
      await that.Remove(that.r_pth);
      that.r_pth = that.r_pth.substring(0, that.r_pth.length - 5);
    },
    async Decompress() {
      var that = this;
      if (!that.decompress) {
        await that.Decode().catch((err) => {
          console.log("decom");
          throw err;
        });
        return;
      }
      await that.Decode().catch((err) => {
        throw err;
      });
      that.Body.op = "remote_compress";
      that.Body.compress_para.is_compress = false;
      that.Body.compress_para.compress_path = that.r_pth;
      await this.Post("解压").catch((err) => {
        console.log("decom caught");
        throw err;
      });
      await that.Remove(that.r_pth);
      that.r_pth = that.r_pth.substring(0, that.r_pth.length - 4);
    },
    async Decode() {
      var that = this;
      if (!that.decode) {
        return;
      }
      that.Body.op = "remote_encode";
      that.Body.encode_para.is_encode = false;
      that.Body.encode_para.encode_path = that.r_pth;
      that.Body.encode_para.password = that.decode_pwd;
      await this.Post("解密").catch((err) => {
        console.log("decode caught");
        throw err;
      });
      await that.Remove(that.r_pth);
      that.r_pth = that.r_pth.substring(0, that.r_pth.length - 5);
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