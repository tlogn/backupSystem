<template>
  <div class="popContainer">
    <img
      id="img"
      style="height: 40%; position: relative"
      src="../assets/LOADGIF.gif"
    />
  </div>
</template>

<script>
export default {
  name: "pop",
  data() {
    return {
      showPop: false,
    };
  },
  methods: {
    Show: function () {
      var that = this;
      that.showPop = true;
      console.log("show");
    },
    Hide: function () {
      var that = this;
      console.log("hide");
      that.showPop = false;
    },
  },
};
</script>

<style>
div.popContainer {
  text-align: center;
  font-size: 100px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
}
#img {
  padding-top: 15%;
}
</style>