<template>
  <div id="background">
    <div
      v-if="img"
      :style="{
        backgroundImage: './assets/background.jpg',
        backgroundSize: cover,
      }"
    ></div>
    <div id="button1">
      <center>
        <a href="./../one.html">
          <button
            id="btn1"
            style="font-size: 25px; height: 70px; width: 130px; cursor: hand"
            title="备份"
          >
            备份
          </button>
        </a>
      </center>
    </div>
    <div id="button2">
      <center>
        <a href="./../two.html">
          <button
            id="btn2"
            style="font-size: 25px; height: 70px; width: 130px"
            title="还原"
          >
            还原
          </button>
        </a>
      </center>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messege: "HTTP_GET_TEST",
    };
  },
  methods: {},
};
</script>

<style>
#background {
  background: url("./assets/background.jpg") center center no-repeat;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%; /**宽高100%是为了图片铺满屏幕 */
  z-index: -1;
  position: absolute;
  background-size: cover;
}
#button1 {
  margin-top: 100px;
  width: 100%;
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

#button2 {
  margin-top: 100px;
  width: 100%;
}
#btn2 {
  -webkit-transition-duration: 0.4s;
  transition-duration: 0.4s;
  padding: 16px 32px;
  text-align: center;
  background-color: white;
  color: black;
  border: 4px solid #e4e009;
  border-radius: 5px;
}
#btn2:hover {
  background-color: #e4e009;
  color: white;
  cursor: pointer;
}
</style>
