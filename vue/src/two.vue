<template>
  <div>
    <h1 style="text-align: center">本地备份</h1>
    <hr />
    <h3>
      <p id="lbl">源路径：{{ source }}</p>
      <p id="lbl">目标路径：{{ destin }}</p>
      <p id="lbl">备份状态：{{ back_status }}</p>
      <center>
        <button
          id="btn2"
          @click="submit()"
          style="height: 35px; width: 120px; font-size: 18px"
        >
          确定备份
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
          <input type="checkbox" id="ckx" value="自定义备份" v-model="opt" />
          <label>自定义备份</label>
          <input type="checkbox" id="ckx" value="压缩" v-model="opt" />
          <label>压缩</label>
          <input type="checkbox" id="ckx" value="打包" v-model="opt" />
          <label>打包</label>
          <br />
          <input type="checkbox" id="ckx" value="加密" v-model="opt" />
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
      <backup @ori="parent_ori"></backup>
    </div>
    <div id="second">
      <target @tar="parent_tar"></target>
    </div>
  </div>
</template>

<script>
import Backup from "./components/origin_path.vue";
import Target from "./components/target_path.vue";
import c from "../common.vue";
import axios from "axios";
export default {
  name: "two",
  components: {
    Backup,
    Target,
  },
  data() {
    return {
      source: "",
      destin: "",
      header: "http://localhost:8090/method",
      opt: [],
      back_status: "",
      Body: c.Body,
      IsShow: true,
      encode_pwd: "",
      s_pth: "",
      d_pth: "",
      compress: false,
      encode: false,
      pack: false,
      custom: false,
    };
  },
  methods: {
    parent_ori: function (data) {
      var that = this;
      that.source = data;
      that.back_status = "";
    },
    parent_tar: function (data) {
      var that = this;
      that.destin = data;
      that.back_status = "";
    },
    async Post(type) {
      console.log("123" + type);
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
              var err = rsp.err;
              that.back_status += type + "失败" + "; ";
              if (type == "打包") throw type + "失败：" + err;
              if (type == "压缩") throw type + "失败：" + err;
              if (type == "加密") throw type + "失败：" + err;
              if (type == "备份") throw type + "失败：" + err;
            }
          })
          .catch(function (error) {
            throw error;
          });
      }
      return "OK";
    },
    submit: function () {
      var that = this;
      that.s_pth = this.source;
      that.d_pth = this.destin;
      var filename = this.s_pth.substring(this.s_pth.lastIndexOf("/"));
      that.d_pth += filename;
      if (this.s_pth == "" || this.d_pth == "") {
        window.alert("源路径或目标路径为空！");
        return;
      }
      var r = window.confirm(
        "您要将文件(夹)：" +
          this.s_pth +
          "\n" +
          "备份到：" +
          this.d_pth +
          "\n" +
          "备份选项：" +
          this.opt +
          "\n" +
          "注意：若目标地址存在重名文件可能会被覆盖！"
      );
      if (r == true) {
        that.back_status = "";
        that.pack = 0;
        that.encode = 0;
        that.compress = 0;
        that.custom = 0;
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "自定义备份") {
            that.custom = 1;
          } else if (that.opt[i] == "压缩") {
            that.compress = 1;
          } else if (that.opt[i] == "打包") {
            that.pack = 1;
          } else if (that.opt[i] == "加密") {
            that.encode = 1;
            if (this.encode_pwd.length == 0) {
              window.alert("密码为空！");
              return;
            }
          }
        }
        this.Encode();
      }
    },
    async Copy() {
      var that = this;
      that.Body.op = "local_copy";
      that.Body.copy_para.origin_path = this.s_pth;
      that.Body.copy_para.backup_path = this.d_pth;
      console.log("4");
      await this.Post("备份").catch((err) => {
        window.alert(err);
      });
      console.log("5");
    },
    async Pack() {
      if (!this.pack) {
        await this.Copy();
        return;
      }
      console.log("3");
      await this.Copy();
      console.log("6");
      var that = this;
      that.Body.op = "local_pack";
      that.Body.pack_para.is_pack = true;
      that.Body.pack_para.pack_path = this.d_pth;
      await this.Post("打包").catch((err) => {
        console.log("packErr")
        window.alert(err);
        return;
      });
      that.Body.op = "local_remove";
      that.Body.dir_para.dir_path = this.d_pth;
      await axios.post(that.header, that.Body);
      that.d_pth += ".pack";
      console.log(that.d_pth);
    },
    async Compress() {
      if (!this.compress) {
        await this.Pack();
        return;
      }
      console.log("2");
      await this.Pack();
      console.log("7");
      var that = this;
      that.Body.op = "local_compress";
      that.Body.compress_para.is_compress = true;
      that.Body.compress_para.compress_path = this.d_pth;
      await this.Post("压缩").catch((err) => {
        window.alert(err);
        return;
      });
    },
    async Encode() {
      if (!this.encode) {
        await this.Compress();
        return;
      }
      console.log("1");
      await this.Compress();
      console.log("8");
      var that = this;
      that.Body.op = "local_encode";
      that.Body.encode_para.is_encode = true;
      that.Body.encode_para.encode_path = this.d_pth;
      that.Body.encode_para.password = this.encode_pwd;
      await this.Post("加密").catch((err) => {
        window.alert(err);
      });
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