<template>
  <div>
    <div class="imgPost" v-if="preview">
        <img class="viewImage" :src="preview">
    </div>

    <form class="imgPost" method="POST" action="/" enctype="multipart/form-data">

      <div>
        <input type="file" name="file" @change="change">
      </div>

      <input cols="50" rows="10" type="text" name="text">
      <input type="submit">
      <input type="hidden" name="_csrf" v-bind:value="csrf">
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return{
      csrf : document.querySelector("meta[name='csrf_token']").content,
      preview: '',
    };
  },

  methods:{
    change: function(e){
      const file =  e.target.files[0];
      if(file && file.type.match(/^image\/(png|jpeg)$/)){
          this.preview = URL.createObjectURL(file);
      }
    },
  },
}
</script>

