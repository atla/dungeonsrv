
<template>
  <v-container>
    <ItemTemplateList />
  </v-container>
</template>

<script>
export default {
  components: {
    ItemTemplateList
  },
  created() {
    this.$store.dispatch("itemTemplates/fetchItemTemplates");
  },
  data: () => ({
    drawer: false
  }),

  props: {
    source: String
  }
};
</script>

