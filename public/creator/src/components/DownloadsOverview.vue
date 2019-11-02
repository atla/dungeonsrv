<template>
  <v-container>
    <v-layout justify-start align-center>
      <span class="title">Downloads</span>
    </v-layout>
    <br />

    <v-tabs left>
      <v-tab>Active / Queued</v-tab>
      <v-tab>Completed</v-tab>

      <v-tab-item :key="0">
        <v-card class="header">
          <v-data-table
            :headers="headers"
            :items="downloads"
            :rows-per-page-items="[25,50,100,200]"
          >
            <template v-slot:items="props">
              <td :key="`${props.id}`">
                <span class="text-truncate">{{ props.item.packageId }}</span>
              </td>

              <td>
                <span class="text-truncate">{{ props.item.packageName }}</span>
              </td>
              <td>
                <span class="text-truncate">{{ props.item.searchName }}</span>
              </td>

              <td class="text-xs-right">
                <v-progress-linear v-model="props.item.percentageDownloaded"></v-progress-linear>
              </td>
              <td>
                <span class="text-truncate">{{ Math.ceil(props.item.percentageDownloaded) }}%</span>
              </td>
            </template>

            <template v-slot:no-results>
              <v-alert :value="true" color="error" icon="warning">No active Downloads.</v-alert>
            </template>
          </v-data-table>
        </v-card>
      </v-tab-item>

      <v-tab-item :key="1">
        <v-card class="header">
          <v-data-table
            :headers="headers"
            :items="downloads"
            :rows-per-page-items="[25,50,100,200]"
          >
            <v-alert :value="true" color="error" icon="warning">No active Downloads.</v-alert>
          </v-data-table>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>



<script>
import { mapState } from "vuex";

export default {
  computed: {
    ...mapState({
      downloads: state => state.downloads.all
    })
  },
  created() {
    setInterval(() => {
      this.$store.dispatch("downloads/fetchDownloads");
    }, 1000);
  },
  data() {
    return {
      rowsPerPageItems: [6, 16, 32],
      pagination: {
        rowsPerPage: 8
      },
      headers: [
        {
          text: "Package ID",
          align: "left",
          sortable: false,
          value: "packageId"
        },
        {
          text: "Package",
          align: "left",
          sortable: false,
          value: "packageId"
        },
        {
          text: "Search",
          align: "left",
          sortable: false,
          value: "packageId"
        },
        {
          text: "Download",
          align: "center",
          sortable: false,
          width: 100,
          value: "percentageDownloaded"
        },
        {
          text: "%",
          align: "right",
          sortable: false,
          value: "percentageDownloaded"
        }
      ]
    };
  }
};
</script>

<style scoped>
.hidden {
  background-color: transparent;
}
.brigthen {
  background-color: #ffffff11;
}

.rounded-card {
  border-radius: 15px;
}
.padd {
  padding: 5 px;
}
.bottom-gradient {
  background-image: linear-gradient(
    to top,
    rgba(0, 0, 0, 0.6),
    transparent 60px
  );
}
</style>
