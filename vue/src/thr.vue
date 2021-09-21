<template>
  <div>
    <h1 style="text-align: center">本地还原</h1>
    <hr />
    <h3>
      <center>
        <p id="lbl">还原源路径：{{ rec_source }}</p>
        <p id="lbl">还原目标路径：{{ rec_destin }}</p>
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
import axios from "axios";
export default {
  name: "thr",
  components: {
    Rec_left,
    Rec_right,
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
      decode_suc: false,
      decompress_suc: false,
      unpack_suc: false,
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
              if (type == "解密") that.decode_suc = true;
              if (type == "解压") that.decompress_suc = true;
              if (type == "解包") that.unpack_suc = true;
            } else {
              if (rsp.err == "redis: nil") {
                if (type == "还原")
                  window.alert(type + "失败：该文件(夹)不是备份文件(夹)");
                else window.alert(type + "失败: 该文件未加密");
              } else window.alert(type + "失败：" + rsp.err);
              that.back_status += type + "失败" + "; ";
            }
          })
          .catch(function (error) {
            that.post_response_msg = null;
            window.alert(error);
          });
      }
    },
    get_rec_destin: function () {
      return "";
    },
    submit: function () {
      var that = this;
      that.rec_destin = this.get_rec_destin();
      var s_pth = this.rec_source;
      var d_pth = this.rec_destin;
      that.decode_suc = true;
      that.decompress_suc = true;
      that.unpack_suc = true;
      if (s_pth == "") {
        window.alert("还原源路径为空！");
        return;
      }
      var opt = this.opt;
      /*var filename = s_pth.substring(s_pth.lastIndexOf("/"));
      d_pth += filename;*/
      var r = window.confirm(
        "您要将文件(夹)：" +
          s_pth +
          "\n" +
          "还原到：" +
          d_pth +
          "\n" +
          "备份选项：" +
          opt +
          "\n" +
          "注意：若还原地址存在重名文件可能会被覆盖！"
      );
      if (r == true) {
        var that = this;
        that.back_status = "";
        var type = 0; //local
        var pack = 0,
          enco = 0,
          compress = 0;
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "解压") {
            compress = 1;
            that.decompress_suc = false;
          } else if (that.opt[i] == "解包") {
            pack = 1;
            that.unpack_suc = false;
          } else if (that.opt[i] == "解密") {
            enco = 1;
            that.decode_suc = false;
          }
        }
        if (enco == 1) {
          that.Body.op = "local_encode";
          that.Body.encode_para.is_encode = false;
          that.Body.encode_para.encode_path = s_pth;
          that.Body.encode_para.password = this.encode_pwd;
          this.Post("解密");
        }
        if (compress == 1) {
          setTimeout(() => {
            if (this.decode_suc) {
              that.Body.op = "local_compress";
              that.Body.compress_para.is_compress = false;
              that.Body.compress_para.compress_path = s_pth;
              this.Post("解压");
            }
          }, 1500);
        }
        if (pack == 1) {
          setTimeout(() => {
            if (this.decompress_suc && this.decode_suc) {
              that.Body.op = "local_pack";
              that.Body.pack_para.is_pack = false;
              that.Body.pack_para.pack_path = s_pth;
              s_pth = s_pth.substring(0, s_pth.length - 5);
              this.Post("解包");
              console.log("abc");
            }
          }, 3000);
        }

        setTimeout(() => {
          console.log([this.unpack_suc,this.decode_suc, this.decode_suc]);
          if (this.unpack_suc && this.decompress_suc && this.decode_suc) {
            that.Body.op = "local_recover";
            that.Body.recover_para.recover_path = s_pth;
            //that.newBody.copy_para. = d_pth;
            this.Post("还原");
          }
        }, 4500);
      }
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