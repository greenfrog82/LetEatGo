<template>
  <div id="app">
    <v-app light>
      <v-toolbar class="white">
        <v-toolbar-title v-text=title></v-toolbar-title>
      </v-toolbar>
      <v-content>
        <router-view/>
      </v-content>
      <v-footer height="auto">
      <v-card flat tile class="flex">
        <!-- https://vuetifyjs.com/en/style/colors -->
        <v-card-title class="black grey--text text--lighten-1">
          <strong class="subheading">Let's contribute our code!</strong>
          <v-btn icon dark class="mx-3">
            <v-icon size="24px">fab fa-github</v-icon>
          </v-btn>
        </v-card-title>
      </v-card>
    </v-footer>
    </v-app>
  </div>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    title: 'LetEatGo'
  })
}
</script>

<style>
</style>
