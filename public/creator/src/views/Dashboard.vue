<template>
  <v-container>
    <ItemTemplateList />
  </v-container>
</template>
<script>
import ItemTemplateList from "../components/ItemTemplateList";

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