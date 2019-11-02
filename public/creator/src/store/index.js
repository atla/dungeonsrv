import Vue from 'vue'
import Vuex from 'vuex'
import itemTemplates from './modules/itemtemplates'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        itemTemplates
    },
    strict: debug,
    plugins: []
})