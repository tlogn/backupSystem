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
      pack_suc: false,
      compress_suc: false,
      encode_suc: false,
      backup_suc: false,
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
            if (rsp.succeed == true) {
              that.back_status += type + "成功" + "; ";
              if (type == "打包") that.pack_suc = true;
              if (type == "压缩") that.compress_suc = true;
              if (type == "加密") that.encode_suc = true;
              if (type == "备份") that.backup_suc = true;
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
    submit: function () {
      var s_pth = this.source;
      var d_pth = this.destin;
      var opt = this.opt;
      var that = this;
      that.pack_suc = true;
      that.compress_suc = true;
      that.encode_suc = true;
      that.backup_suc = false;
      var filename = s_pth.substring(s_pth.lastIndexOf("/"));
      d_pth += filename;
      if (s_pth == "" || d_pth == "") {
        window.alert("源路径或目标路径为空！");
        return;
      }
      var r = window.confirm(
        "您要将文件(夹)：" +
          s_pth +
          "\n" +
          "备份到：" +
          d_pth +
          "\n" +
          "备份选项：" +
          opt +
          "\n" +
          "注意：若目标地址存在重名文件可能会被覆盖！"
      );
      if (r == true) {
        var that = this;
        that.back_status = "";
        var type = 0; //local
        var pack = 0,
          enco = 0,
          compress = 0,
          custom = 0;
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "自定义备份") {
            custom = 1;
          } else if (that.opt[i] == "压缩") {
            that.compress_suc = false;
            compress = 1;
          } else if (that.opt[i] == "打包") {
            that.pack_suc = false;
            pack = 1;
          } else if (that.opt[i] == "加密") {
            that.encode_suc = false;
            enco = 1;
            if (this.encode_pwd.length == 0) {
              window.alert("密码为空！");
              return;
            }
          }
        }

        that.Body.op = "local_copy";
        that.Body.copy_para.origin_path = s_pth;
        that.Body.copy_para.backup_path = d_pth;
        this.Post("备份");
        if (custom == 1) {
        } //TODO
        if (pack == 1) {
          setTimeout(() => {
            if (this.backup_suc) {
              that.Body.op = "local_pack";
              that.Body.pack_para.is_pack = true;
              that.Body.pack_para.pack_path = d_pth;
              d_pth += ".pack";
              this.Post("打包");
            }
          }, 1500);
        }
        if (compress == 1) {
          setTimeout(() => {
            if (this.backup_suc && this.pack_suc) {
              that.Body.op = "local_compress";
              that.Body.compress_para.is_compress = true;
              that.Body.compress_para.compress_path = d_pth;
              this.Post("压缩");
            }
          }, 3000);
        }
        if (enco == 1) {
          setTimeout(() => {
            if (this.backup_suc && this.pack_suc && this.compress_suc) {
              that.Body.op = "local_encode";
              that.Body.encode_para.is_encode = true;
              that.Body.encode_para.encode_path = d_pth;
              that.Body.encode_para.password = this.encode_pwd;
              this.Post("加密");
            }
          }, 4000);
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