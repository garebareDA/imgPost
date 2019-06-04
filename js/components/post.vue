<template>
  <div>
    <button @click="transPage('post')">投稿</button>

    <form method="POST" action="/" enctype="multipart/form-data">
      <input type="file" name="file">
      <input type="text" name="text">
      <input type="submit">
      <input type="hidden" name="_csrf" v-bind:value="csrf">
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return{
      csrf : document.querySelector("meta[name='csrf_token']").content
    };
  },
}
</script>

