'use strict';

import Vue from 'vue';
import infinite from './components/app.vue';

const api = window.location + "/api";

new Vue({
  el:'#app',

  data: {
      page: 0,
      last:10,
      getData: [],
  },

  components: {
    'infie' : infinite
  }
});