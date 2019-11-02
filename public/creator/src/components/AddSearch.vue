<template>
  <v-combobox
    v-model="name"
    :items="getSeriesSuggestions()"
    label="Add Search"
    @keyup.native.enter="addSearch"
    solo-inverted
    :menu-props="{openOnClick: false}"
    :open-on-clear="false"
    :validate-on-blur="false"
    :auto-select-first="true"
    ref="combobox"
  ></v-combobox>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  name: "AddTodo",
  data() {
    return {
      title: "",
      name: ""
    };
  },
  methods: {
    ...mapActions(["createSearch"]),
    ...mapGetters({
      getSeriesSuggestions: "series/getSeriesSuggestions"
    }),
    onChange(e) {
      console.dir("ON change " + e);
    },
    addSearch(e) {
      e.preventDefault();

      if (this.name === "") return;

      // dispatch with a payload
      this.$store.dispatch("searches/createSearch", this.name);
      this.name = "";
      this.$refs["combobox"].blur();
    }
  }
};
</script>


