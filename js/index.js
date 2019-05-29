'use strict';

import Vue from 'vue';
import InfiniteLoading from 'vue-infinite-loading'

Vue.component('infinite-loading', InfiniteLoading)

new Vue({
  el:'#app',
  data(){
    return{
      page: 0,
      last:10,
      list: [],
    };
  },

  methods:{
    infiniteHandler($state){
      
    }
  }

});