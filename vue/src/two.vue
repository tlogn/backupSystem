<template>
  <div id="backup">
    <div id="first">
      <h2>选择要备份的文件</h2>
      <div>
        <input
          type="file"
          id="file1"
          style="display:none"
          onchange="document.getElementById('textfield').value=this.value"
          ref="file"
        />
        <button id="btn2" @click="expl" style="cursor:pointer">浏览文件</button>
        <br /><br />
        <input
          type="text"
          name="textfield"
          id="textfield"
          class="txt"
          style="width: 300px; height: 25px; font-size: 18px"
        />
      </div>
      <br />
      <div>
        <button
          id="btn1"
          style="font-size: 25px; height: 70px; width: 130px; cursor: hand"
          title="备份"
          @click="submit()"
        >
          上传
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "backup",
  data() {
    return {
      msg: "",
    };
  },
  methods: {
    submit: function () {
      window.alert(this.$refs.file.files[0].name)
    },
    expl:function() {
      file1.click()
    }
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
  padding: 5px 10px;
  background: #00b0f0;
  color: #fff;
  border: none;
  border-radius: 5px;
}
</style>