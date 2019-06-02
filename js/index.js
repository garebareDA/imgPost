'use strict';

import Vue from 'vue';
import infinite from './components/app.vue';
import user from './components/user.vue'

new Vue({
  el:'#app',

  components: {
    'infie' : infinite,
    'user' : user
  }
});