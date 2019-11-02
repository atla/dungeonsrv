import Vue from 'vue'
import Vuetify from 'vuetify'
import App from './App.vue'
import router from './router'
import store from './store'
import moment from "moment";
import 'vuetify/src/stylus/app.styl'
import "vue-material-design-icons/styles.css"
import axios from 'axios'
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader

axios.defaults.withCredentials = false;

Vue.config.productionTip = false
Vue.use(Vuetify, {
  iconfont: 'mdi'
})

Vue.filter('formatDate', function (value) {
  if (value) {
    return moment(String(value)).format('DD.MM.YYYY hh:mm')
  }
});

new Vue({
  router,
  store,
  watch: {
    $route: function () {}
  },


  render: h => h(App)
}).$mount('#app')