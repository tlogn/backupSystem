<template>
  <div>
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
      </center>
    </h3>
    <hr />
    <div id="first">
      <h2>选择要还原到的路径</h2>
      <center>
        <div id="column3">
          <input type="checkbox" id="ckx" value="压缩" v-model="opt" />
          <label>解压</label>
          <input type="checkbox" id="ckx" value="打包" v-model="opt" />
          <label>解包</label>
          <input type="checkbox" id="ckx" value="加密" v-model="opt" />
          <label>解密</label>
          <br />
        </div>
      </center>
      <rec_left @ori="parent_ori"></rec_left>
    </div>
    <div id="second">
      <rec_right @tar="parent_tar"></rec_right>
    </div>
  </div>
</template>

<script>
import Rec_left from "./components/restore_left.vue";
import Rec_right from "./components/restore_right.vue";
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
      newBody: "",
      opt: [],
      back_status: "",
      default_pth: "",
      Body: {
        op: "local_dir",
        get_dir_para: {
          dir_path: "",
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
        data = this.newBody;
      var that = this;
      //console.log(addr);
      //console.log(data);
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == "false") {
              window.alert("还原失败:" + rsp.err);
              that.back_status += type + "失败" + "; ";
            } else {
              that.back_status += type + "成功" + "; ";
            }
          })
          .catch(function (error) {
            that.post_response_msg = null;
            window.alert(error);
          });
      }
    },
    submit: function () {
      var s_pth = this.rec_source;
      var d_pth = this.rec_destin;
      var opt = this.opt;
      var filename = s_pth.substring(s_pth.lastIndexOf("/"));
      d_pth += filename;
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
        that.newBody = that.Body;
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "解压") {
            compress = 1;
          } else if (that.opt[i] == "解包") {
            pack = 1;
          } else if (that.opt[i] == "解密") {
            enco = 1;
          }
        }

        if (enco == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_encode";
          else that.newBody.op = "remote_encode";
          that.newBody.encode_para.is_encode = false;
          that.newBody.encode_para.encode_para = s_pth;
          this.Post("解密");
        }
        if (compress == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_compress";
          else that.newBody.op = "remote_compress";
          that.newBody.compress_para.is_compress = false;
          that.newBody.compress_para.compress_path = s_pth;
          this.Post("解压");
        }
        if (pack == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_pack";
          else that.newBody.op = "remote_pack";
          that.newBody.pack_para.is_pack = false;
          that.newBody.pack_para.pack_path = s_pth;
          this.Post("解包");
        }
        that.newBody = that.Body;
        that.newBody.op = "local_recover";
        that.newBody.recover_para.recover_path = s_pth;
        //that.newBody.copy_para. = d_pth;
        this.Post("还原");
        console.log(that.back_status);
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