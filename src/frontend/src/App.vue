<template>
  <div id="app">
    <v-app light>
      <v-toolabr class="white">
        <v-toolbar-title v-text=title></v-toolbar-title>
      </v-toolabr>
      <v-content>
        <router-view/>
      </v-content>
      <v-footer height="auto">
      <v-card flat tile class="flex">
        <v-card-title class="teal white--text">
          <strong class="subheading">Get connected with us on social networks!</strong>
          <v-spacer></v-spacer>
          <v-btn
            icon
            dark
            class="mx-3"
          >
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
