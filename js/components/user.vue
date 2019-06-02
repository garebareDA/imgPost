<template>
  <div>
    <p>テスト</p>
    <tr v-for="(item, index) in list" :key="index">
      <td>{{item.Text}}</td>
      <td>{{item.userName}}</td>

    </tr>
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