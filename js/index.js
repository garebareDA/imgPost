'use strict';

import Vue from 'vue';
import infinite from './components/app.vue';
import user from './components/user.vue';
import post from './components/post.vue';

new Vue({
  el:'#app',

  data:{
    currentPage : 'infie'
  },

  methods:{
    transPage: function(page){
      this.currentPage = page
    }
  },

  components: {
    'infie' : infinite,
    'user' : user,
    'post' : post
  }
});