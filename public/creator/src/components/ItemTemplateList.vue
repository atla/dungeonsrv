<template>
  <div>
    <v-layout justify-start align-center>
      <span class="title">Item Templates</span>
    </v-layout>

    <v-container fluid grid-list-md>
      <v-data-table
        :headers="headers"
        :items="itemTemplates"
        class="elevation-5"
        expand
        :rows-per-page-items="[25,50,100,200]"
      >
        <template v-slot:items="props" :to="`/itemTemplates/${props.item.id}`">
          <router-link tag="tr" :to="`/itemTemplates/${props.item.id}`">
            <td>
              <span text-truncate>{{ props.item.name }}</span>
            </td>
            <td class="text-xs-right">{{ props.item.description }}</td>
            <td class="text-xs-right">{{ props.item.itemType }}</td>
          </router-link>
        </template>
      </v-data-table>
    </v-container>
  </div>
</template>

<script>
import { mapState, mapGetters } from "vuex";

export default {
  computed: {
    ...mapState({
      itemTemplates: state => state.itemTemplates.all
    }),
    ...mapGetters("items", {})
  },

  data() {
    return {
      headers: [
        {
          text: "Item ID",
          align: "left",
          sortable: true,
          value: "id"
        },
        { text: "Name", value: "name" },
        { text: "Description", value: "description" },
        { text: "Item Type", value: "itemType" }
      ],
      items: []
    };
  }
};
</script>


<style scoped>
a {
  text-decoration: none;
}
</style>
