<template>
  <v-card class="header" v-if="getSearchByID.searchResults !== null">
    <v-card-title>
      Search Results
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-download"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="getSearchByID.searchResults"
      :search="search"
      :rows-per-page-items="[25,50,100,200]"
    >
      <template v-slot:items="props">
        <td>
          <v-btn v-on:click="download(props.item)" flat icon color="green">
            <v-icon>mdi-layers</v-icon>
          </v-btn>
        </td>

        <td :key="`${props.id}`">
          <span class="text-truncate">{{ props.item.name }}</span>
        </td>
        <td class="text-xs-right">
          <span class="text-truncate">{{ props.item.bot }}</span>
        </td>
        <td class="text-xs-right">{{ props.item.packageNumber }}</td>
        <td class="text-xs-right">{{ props.item.server }}</td>
        <td class="text-xs-right">{{ props.item.channel }}</td>
        <td class="text-xs-right">{{ props.item.downloads }}</td>
      </template>
      <template v-slot:no-results>
        <v-alert
          :value="true"
          color="error"
          icon="warning"
        >Your search for "{{ search }}" found no results.</v-alert>
      </template>
    </v-data-table>
  </v-card>
</template>




<script>
import { mapState, mapGetters } from "vuex";

export default {
  computed: {
    ...mapState({}),
    ...mapGetters({
      getSearchByID: "searches/getSearchByID"
    })
  },

  methods: {
    download: function(pack) {
      this.$store.dispatch("searches/downloadPackage", pack);
    }
  },

  data() {
    return {
      search: "",
      headers: [
        {
          text: "DL Active",
          align: "left",
          sortable: false,
          value: "download.isDownloading"
        },
        {
          text: "Package name",
          align: "left",
          sortable: true,
          value: "name",
          width: "80"
        },
        { text: "Bot", value: "bot", sortable: true, width: "80%" },
        { text: "Package Number", value: "packageNumber", sortable: true },
        { text: "Server", value: "server", sortable: true },
        { text: "Channel", value: "channel", sortable: true },
        { text: "Downloads", value: "download", sortable: true }
      ],
      items: []
    };
  }
};
</script>

<style scoped>
.detail-container {
  width: 800px;
}
.header {
  margin-top: 20px;
  margin-bottom: 10px;
}
</style>


