<template>
  <div id="backup">
    <div id="first">
      <h2>选择要备份的文件</h2>
      <form method="post" enctype="multipart/form-data">
        <div>
          <input type="file" id="image_uploads" name="image_uploads" />
        </div>
        <div class="preview">
          <p>No files currently selected for upload</p>
        </div>
        <div>
          <button
            id="btn1"
            @click="submit_path"
            style="font-size: 25px; height: 70px; width: 130px; cursor: hand"
            title="备份"
          >
            上传
          </button>
        </div>
      </form>
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
</style>