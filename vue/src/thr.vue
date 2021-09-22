<template>
  <div>
    <div v-show="showPop">
      <pop></pop>
    </div>
    <h1 style="text-align: center">本地还原</h1>
    <hr />
    <h3>
      <center>
        <p id="lbl">还原源路径：{{ rec_source }}</p>
        <p id="lbl">文件将被还原到备份时所在的位置</p>
        <p id="lbl">还原状态：{{ back_status }}</p>
        <button
          id="btn2"
          @click="submit()"
          style="height: 35px; width: 120px; font-size: 18px"
        >
          确定还原
        </button>
        <a href="/">
          <button id="btn2" style="height: 35px; width: 120px; font-size: 18px">
            返回首页
          </button>
        </a>
      </center>
    </h3>
    <hr />
    <center>
      <h2>选择还原选项</h2>
      <div id="column3">
        <input type="checkbox" id="ckx" value="解压" v-model="opt" />
        <label>解压</label>
        <input type="checkbox" id="ckx" value="解包" v-model="opt" />
        <label>解包</label>
        <br />
        <input type="checkbox" id="ckx" value="解密" v-model="opt" />
        <label
          >解密
          <input
            v-model="encode_pwd"
            type="password"
            v-show="IsShow"
            placeholder="请输入密码"
          />
        </label>
        <br />
      </div>
      <rec_right @tar="parent_tar"></rec_right>
    </center>
  </div>
</template>

<script>
import Rec_left from "./components/restore_left.vue";
import Rec_right from "./components/restore_right.vue";
import c from "../common.vue";
import pop from "./components/pop.vue";
import axios from "axios";
export default {
  name: "thr",
  components: {
    Rec_left,
    Rec_right,
    pop,
  },
  data() {
    return {
      rec_source: "",
      rec_destin: "",
      header: "http://localhost:8090/method",
      opt: [],
      back_status: "",
      default_pth: "",
      Body: c.Body,
      encode_pwd: "",
      IsShow: true,
      decode: false,
      decompress: false,
      unpack: false,
      download: false,
      s_pth: "",
      d_pth: "",
      showPop: false,
    };
  },

  methods: {
    parent_ori: function (data) {
      var that = this;
      that.rec_destin = data;
      that.back_status = "";
    },
    parent_tar: function (data) {
      var that = this;
      that.rec_source = data;
      that.back_status = "";
    },
    Post: async function (type) {
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
              that.back_status += type + "失败" + "; ";
              window.alert(type + "失败：" + rsp.err);
              throw type + "失败：" + rsp.err;
            }
          })
          .catch(function (error) {
            throw error;
          });
      }
    },
    get_rec_destin: function () {
      return "";
    },
    submit: async function () {
      var that = this;
      that.rec_destin = this.get_rec_destin();
      that.s_pth = this.rec_source;
      that.d_pth = this.rec_destin;
      that.decode = false;
      that.decompress = false;
      that.unpack = false;
      if (that.s_pth == "") {
        window.alert("还原源路径为空！");
        return;
      }
      var opt = this.opt;
      /*var filename = s_pth.substring(s_pth.lastIndexOf("/"));
      d_pth += filename;*/
      var r = window.confirm(
        "您要将文件(夹)：" +
          that.s_pth +
          "\n" +
          "还原到：" +
          that.d_pth +
          "\n" +
          "备份选项：" +
          opt +
          "\n" +
          "注意：若还原地址存在重名文件可能会被覆盖！"
      );
      if (r == true) {
        that.back_status = "";
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "解压") {
            that.decompress = true;
          } else if (that.opt[i] == "解包") {
            that.unpack = true;
          } else if (that.opt[i] == "解密") {
            that.decode = true;
          }
        }
        that.showPop = true;
        await this.Download().catch((err) => {
          that.showPop = false;
        });
        that.showPop = false;
      }
    },
    async remove(pth) {
      var that = this;
      that.Body.op = "local_remove";
      that.Body.dir_para.dir_path = pth;
      await axios.post(that.header, that.Body);
    },
    async Download() {
      var that = this;
      await this.Unpack().catch((err) => {
        throw err;
      });
      that.Body.op = "local_recover";
      that.Body.recover_para.recover_path = that.s_pth;
      //that.newBody.copy_para. = d_pth;
      await this.Post("还原").catch((err) => {
        throw err;
      });
    },
    async Unpack() {
      var that = this;
      if (!that.unpack) {
        await this.Decompress().catch((err) => {
          throw err;
        });
        return;
      }
      await this.Decompress().catch((err) => {
        throw err;
      });
      that.Body.op = "local_pack";
      that.Body.pack_para.is_pack = false;
      that.Body.pack_para.pack_path = that.s_pth;
      await this.Post("解包").catch((err) => {
        throw err;
      });
      await this.remove(that.s_pth);
      that.s_pth = that.s_pth.substring(0, that.s_pth.length - 5);
    },
    async Decompress() {
      var that = this;
      if (!this.decompress) {
        await this.Decode().catch((err) => {
          throw err;
        });
        return;
      }
      await this.Decode().catch((err) => {
        throw err;
      });
      that.Body.op = "local_compress";
      that.Body.compress_para.is_compress = false;
      that.Body.compress_para.compress_path = that.s_pth;
      await this.Post("解压").catch((err) => {
        throw err;
      });
      await this.remove(that.s_pth);
      that.s_pth = that.s_pth.substring(0, that.s_pth.length - 4);
    },
    async Decode() {
      var that = this;
      if (!that.decode) return;
      that.Body.op = "local_encode";
      that.Body.encode_para.is_encode = false;
      that.Body.encode_para.encode_path = that.s_pth;
      that.Body.encode_para.password = that.encode_pwd;
      await this.Post("解密").catch((err) => {
        throw err;
      });
      await this.remove(that.s_pth);
      that.s_pth = that.s_pth.substring(0, that.s_pth.length - 5);
    },
  },
};
</script>

<style>
@import "./backup.CSS";
#first {
  width: 50%;
  float: left;
  text-align: center;
}
#second {
  width: 50%;
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