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

    </div>
    <infinite-loading @infinite="infiniteHandler"></infinite-loading>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading';
import axios from 'axios';

const api = window.location.origin + '/api'
console.log(api)

export default {

  data() {
    return {
      page: 0,
      last: 10,
      list: [],
    };
  },

  methods: {
    infiniteHandler($state) {
      axios.get(api, {
        params: {
          page: this.page,
          last: this.last
        }
      }).then((data) => {
        if(data.data.length == 10){
          this.page += 10;
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