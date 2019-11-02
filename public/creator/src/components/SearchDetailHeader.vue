<template>
  <v-layout>
    <v-card class="detail-container">
      <v-layout>
        <v-flex xs3>
          <v-img :src="getImageForID(series, $route.params.id)" height="200px"></v-img>
        </v-flex>
        <v-flex xs9>
          <v-card-title primary-title>
            <div>
              <div class="headline">{{getSearchByID.keywords}}</div>

              <div>(2014)</div>
            </div>
          </v-card-title>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn flat color="red" @click="deleteSearch">Delete</v-btn>
          </v-card-actions>
        </v-flex>
      </v-layout>
    </v-card>
  </v-layout>
</template>


<script>
import { mapState, mapGetters } from "vuex";

export default {
  name: "SearchDetailHeader",
  computed: {
    ...mapState({
      searches: state => state.searches.all,
      series: state => state.series.all
    }),
    ...mapGetters({
      getImageForID: "searches/pictureForSearch",
      getSearchByID: "searches/getSearchByID",
      getSeriesByID: "series/getSeriesByID"
    })
  },
  created() {
    this.$store.dispatch("series/fetchSeries");
    this.$store.dispatch("searches/getAllSearches");
  },
  data() {
    return {
      title: "",
      name: ""
    };
  },
  methods: {
    deleteSearch() {
      this.$store.dispatch("searches/deleteSearchByID", this.$route.params.id);
      this.$router.go(-1);
    }
  }
};
</script>

<style scoped>
.detail-container {
  width: 100%;
}
</style>


