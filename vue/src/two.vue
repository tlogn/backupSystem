<template>
  <div>
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
          <input type="checkbox" id="ckx" value="加密" v-model="opt" />
          <label>加密</label>
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
      newBody: "",
      opt: [],
      back_status: "",
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
        data = this.newBody;
      var that = this;
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == "false") {
              window.alert("上传失败:" + rsp.err);
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
      var s_pth = this.source;
      var d_pth = this.destin;
      var opt = this.opt;
      var filename = s_pth.substring(s_pth.lastIndexOf("/"));
      d_pth += filename;
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
        that.newBody = that.Body;
        for (var i = 0; i < that.opt.length; i++) {
          if (that.opt[i] == "自定义备份") {
            custom = 1; 
          } else if (that.opt[i] == "压缩") {
            compress = 1;
          } else if (that.opt[i] == "打包") {
            pack = 1;
          } else if (that.opt[i] == "加密") {
            enco = 1;
          }
        }  
        that.newBody.op = "local_copy";
        that.newBody.copy_para.origin_path = s_pth;
        that.newBody.copy_para.backup_path = d_pth;
        this.Post("备份");
        if (custom == 1) {} //TODO
        if (pack == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_pack";
          else that.newBody.op = "remote_pack";
          that.newBody.pack_para.is_pack = true;
          that.newBody.pack_para.pack_path = d_pth;
          this.Post("打包");
        }
        if (compress == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_compress";
          else that.newBody.op = "remote_compress";
          that.newBody.compress_para.is_compress = true;
          that.newBody.compress_para.compress_path = d_pth;
          this.Post("压缩");
        }
        if (enco == 1) {
          that.newBody = that.Body;
          if (type == 0) that.newBody.op = "local_encode";
          else that.newBody.op = "remote_encode";
          that.newBody.encode_para.is_encode = true;
          that.newBody.encode_para.encode_para = d_pth;
          this.Post("加密");
        }
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