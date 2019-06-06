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
    transPage(){
      if (this.currentPage === 'infie'){
        this.currentPage = 'post';
        const postButton = document.getElementById('post');
        postButton.innerHTML = "戻る";
        }else{
          this.currentPage = 'infie';
          const postButton = document.getElementById('post');
          postButton.innerHTML = "投稿";
        }
      }
    },

  components: {
    'infie' : infinite,
    'user' : user,
    'post' : post
  }
});