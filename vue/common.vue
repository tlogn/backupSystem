
<script>
const Body = {
  op: "remote_upload",

  user_name: "",

  trans_para: {
    local_path: "",
    remote_path: "",
  },

  login_para: {
    username: "",
    password: "",
  },

  dir_para: {
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
    password: "",
  },

  pack_para: {
    is_pack: false,
    pack_path: "",
  },
};

export default {
  Body,
};
</script>