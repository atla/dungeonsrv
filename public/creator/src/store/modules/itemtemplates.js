import api from '../../api/api'
import store from '..'

// initial state
const state = {
    all: []
}

const getters = {
    getItemTemplateByID(state, id) {
        return state.all.find(itemTemplates => itemTemplates.id === id)
    },
    getItemTemplates(state) {
        return state.all
    }
}

// actions
const actions = {
    fetchItemTemplates(context) {
        api.getItemTemplates(itemTemplates => {
            context.commit('SET_ITEM_TEMPLATES', itemTemplates)
        })
    },

    createItemTemplate(context, itemTemplate) {
        api.createItemTemplate(itemTemplate, () => {
            store.dispatch('items/itemTemplates')
        })
    }
}

// mutations
const mutations = {
    SET_ITEM_TEMPLATES(state, itemTemplates) {
        state.all = itemTemplates
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}