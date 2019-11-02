<template>
  <div id="app">
    <v-app dark id="inspire">
      <v-navigation-drawer app clipped floating>
        <v-list class="pt-0">
          <v-list-tile
            v-for="item in items"
            :key="item.text"
            @click.stop="drawer = !drawer"
            :to="item.route"
          >
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>

            <v-list-tile-content>
              <v-list-tile-title>{{ item.text }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>

      <v-toolbar color="darkblue" dense fixed app>
        <!-- <v-toolbar-side-icon v-show="!show" @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <v-btn icon v-show="show" @click.stop="$router.go(-1)">
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>-->

        <v-toolbar-title class="mx-1 align-left">
          <span class="title">Creator</span>
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <!-- <v-toolbar-items>
          <v-btn
            flat
            icon
            v-on:click="toggleBots()"
            round
            :color="`${isActive() ? 'green' : 'grey'}`"
            dark
          >
            <v-icon>mdi-cached</v-icon>
          </v-btn>
        </v-toolbar-items>-->
      </v-toolbar>
      <v-content>
        <router-view></router-view>
      </v-content>
    </v-app>
  </div>
</template>
<script>
import { mapState, mapGetters } from "vuex";
export default {
  components: {},

  computed: {
    ...mapState({}),
    ...mapGetters({})
  },

  methods: {
    isActive: function() {
      return true;
    }
    // ,

    // toggleBots: function() {
    //   this.$store.dispatch("systemStates/fetchSystemStates");
    // }
  },
  created() {
    this.$store.dispatch("systemStates/fetchSystemStates");
  },
  data: () => ({
    drawer: false,
    mini: true,
    backArrow: true,
    items: [
      { icon: "mdi-view-dashboard", text: "Overview", route: "/" },
      {
        icon: "mdi-view-module",
        text: "Item Templates",
        route: "/ItemTemplates"
      }
    ]
  }),
  props: {
    source: String
  }
};
</script>

<style scoped>
</style>


