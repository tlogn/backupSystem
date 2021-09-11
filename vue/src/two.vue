<template>
  <div id="backup">
    <div id="first">
      <h2>选择备份选项</h2>
      <center>
        <div id="column3">
          <input
            type="checkbox"
            id="ckx"
            value="备份到服务器"
            v-model="checkedNames"
          />
          <label>备份到服务器</label>
          <input type="checkbox" id="ckx" value="压缩" v-model="checkedNames" />
          <label>压缩</label>
          <input type="checkbox" id="ckx" value="打包" v-model="checkedNames" />
          <label>打包</label>
          <input type="checkbox" id="ckx" value="加密" v-model="checkedNames" />
          <label>加密</label>
          <br />
        </div>
      </center>
      <h2>选择要备份的文件</h2>
      <input
        placeholder="请输入目录，如'C:'或'Users'。注意不带斜杠"
        v-model="root"
        style="height: 23px; width: 400px; font-size: 18px"
      />
      <br><br>
      <div>
        <button id="btn2" @click="ini_get(root)">浏览文件</button>
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
          <div id="for" v-for="fil in lis">
            <div v-if="fil.is_dir == true">
              <ul id="column1" style="width: 350px; padding: 8px">
                <div id="fil1">
                  <button
                    @click="send(fil.file_name, fil.is_dir, checkedNames)"
                    id="btn2"
                  >
                    备份
                  </button>
                  <button @click="sele(fil.file_name, fil.is_dir)" id="btn3">
                    选择
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
                    @click="send(fil.file_name, fil.is_dir, checkedNames)"
                    id="btn2"
                  >
                    备份
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
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
axios.defaults.headers.post["content-type"] = "application/json";
export default {
  name: "backup",
  data() {
    return {
      msg: "",
      header: "http://localhost:8090/method",
      lis: "",
      last_len: 0,
      checkedNames: [],
      newBody: null,
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
    ini_get: function (para) {
      var that = this;
      that.Body.get_dir_para.dir_path = para + "/";
      //window.alert(that.Body.get_dir_para.dir_path);
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
      } else {
      }
    },
    Return: function () {
      var that = this;
      var pth = that.Body.get_dir_para.dir_path;
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
    post: function() {
      var addr = this.header, data = this.newBody;
      if (addr == null) {
        window.alert("Empty URL");
      } else {
        axios
          .post(addr, data)
          .then(function (response) {
            var rsp = response.data;
            if (rsp.succeed == "true") {
              window.alert(上传成功);
            }
          })
          .catch(function (error) {
            that.post_response_msg = null;
            window.alert(error);
          });
      }
    },
    send: function (para, type, opt) {
      var pth = this.Body.get_dir_para.dir_path + para;
      var r = window.confirm(
        "您要备份的文件(夹)是：" + pth + "\n" + "备份选项：" + opt
      );
      if (r == true) {
        var that = this;
        var type = 0; //local
        var pack = 0, enco = 0, compress = 0;
        that.newBody = that.Body;
        for (var i = 0; i < that.checkedNames.length; i++) {
          if (that.checkedNames[i] == "备份到服务器") {
            type = 1; //remote
          } else if (that.checkedNames[i] == "压缩") {
            compress = 1;
          } else if (that.checkedNames[i] == "打包") {
            pack = 1;
          } else if (that.checkedNames[i] == "加密") {
            enco = 1;
          }
        }
        if (type == 0) {
          that.newBody.op = "local_copy"; 
        } else {
          that.newBody.op = "remote_copy";
        }
        that.newBody.copy_para.origin_path = that.newBody.get_dir_para.dir_path;
        that.newBody.copy_para.backup_path = "xxx";
        this.post();
        if (that.checkedNames.length == 0 || (that.checkedNames.length == 1 && type == 1)) {
          return; 
        }
        if (pack == 1) {
          that.newBody = that.Body;
          if (type == 0) 
            that.newBody.op = "local_pack";
          else 
            that.newBody.op = "remote_pack";
          that.newBody.pack_para.is_pack = true;
          that.newBody.pack_para.pack_path = "xxx";
          this.post();
        }
        if (compress == 1) {
          that.newBody = that.Body;
          if (type == 0) 
            that.newBody.op = "local_compress";
          else 
            that.newBody.op = "remote_compress";
          that.newBody.compress_para.is_compress = true;
          that.newBody.compress_para.compress_path = "xxx";
          this.post();
        }
        if (enco == 1) {
          that.newBody = that.Body;
          if (type == 0) 
            that.newBody.op = "local_encode";
          else 
            that.newBody.op = "remote_encode";
          that.newBody.encode_para.is_encode = true;
          that.newBody.encode_para.encode_para = "xxx";
          this.post();
        }
      }
    },
  },
};
</script>

<style>
#first {
  width: 50%;
  float: left;
  text-align: center;
}
#second {
  width: 50%;
  float: right;
}
#btn1 {
  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
  padding: 16px 32px;
  text-align: center;
  background-color: white;
  color: black;
  border: 4px solid #e4e009;
  border-radius: 5px;
}
#btn1:hover {
  background-color: #e4e009;
  color: white;
  cursor: pointer;
}
#btn2 {
  cursor: pointer;
  padding: 5px 10px;
  background: #00b0f0;
  color: #fff;
  border: none;
  border-radius: 5px;
}
#btn3 {
  cursor: pointer;
  padding: 5px 10px;
  background: #f560099d;
  color: #fff;
  border: none;
  border-radius: 5px;
}
#fil2 {
  color: #00b0f0;
  text-align: left;
}
#column2 {
  border: #00b0f0;
  border-style: solid;
  border-radius: 5px;
}
#fil1 {
  color: #f560099d;
  text-align: left;
}
#column1 {
  border: #f560099d;
  border-style: solid;
  border-radius: 5px;
}
#column3 {
  border: #1a0b029d;
  border-style: solid;
  border-radius: 5px;
  width: 80%;
  height: 60px;
  font-size: 20px;
}
#ckx {
  height: 15px;
  width: 15px;
  margin-left: 25px;
  margin-top: 10px;
  padding: 70px 0;
}
</style>