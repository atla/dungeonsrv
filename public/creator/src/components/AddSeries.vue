<template>
  <v-container fill-width>
    <v-form ref="form" @submit="addSeries">
      <v-text-field v-model="name" label="Add Series" required solo-inverted></v-text-field>
    </v-form>
  </v-container>
</template>


<script>
import { mapActions } from "vuex";

export default {
  name: "CreateSeries",
  data() {
    return {
      title: "",
      name: ""
    };
  },
  methods: {
    ...mapActions(["createSeries"]),

    addSeries(e) {
      e.preventDefault();
      console.log(this.data);
      // dispatch with a payload
      this.$store.dispatch("series/createSeries", this.name);
      this.name = "";
    }
  }
};
</script>


