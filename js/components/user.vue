<template>
  <div>
    <div class="box" v-for="(item, index) in list" :key="index">

      <div>

        <img class="image" v-bind:src ="'/img/' + item.id + '.jpg'">
        <div class="text">{{item.Text}}</div>

      </div>

      <div>
        <img class="icon" v-bind:src="'/icon/' + item.icon + '.jpg'">
        <div class="userName">{{item.userName}}</div>
      </div>
      <div v-if="item.delete == true">
      <form action="/delete" method="post">
        <input type="hidden" name="_method" value="DELETE">
        <input type="hidden" name="ID" v-bind:value="item.id">
        <input type="hidden" name="userID" v-bind:value="item.userID">
        <input class="deleteButton" type="submit" value="削除">
        <input type="hidden" name="_csrf" v-bind:value="csrf">
      </form>
      </div>

    </div>
    <infinite-loading @infinite="infiniteHandler"></infinite-loading>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading';
import axios from 'axios';

const api = window.location.origin + '/api/user'
const userID = window.location.pathname.split('/')[2]

export default {

  data() {
    return {
      id: userID,
      last: 0,
      list: [],
      csrf : document.querySelector("meta[name='csrf_token']").content,
    };
  },

  methods: {
    infiniteHandler($state) {
      axios.get(api, {
        params: {
          id: this.id,
          last: this.last
        }
      }).then((data) => {
        if(data.data.length == 10){
          this.last += 10;
          this.list = data.data;
          console.log(data.data);
          $state.loaded();
        }else{
          this.list = data.data;
          console.log(data.data)
          $state.complete();
        }
      });
    }
  },

  components: {
    InfiniteLoading,
  },
};
</script>