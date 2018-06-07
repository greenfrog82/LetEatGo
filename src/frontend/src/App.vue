<template>
  <div id="app">
    <v-app light>
      <v-toolbar app class="white">
        <v-toolbar-title v-text=title></v-toolbar-title>
      </v-toolbar>
      <v-content>
        <router-view/>
      </v-content>
      <v-footer height="auto">
        <v-card flat title class="flex">
          <!-- https://vuetifyjs.com/en/style/colors -->
          <v-card-title class="black grey--text text--lighten-1">
            <a href='/' class="subheading">
              <b class="grey--text text--lighten-1">LetEatGo</b>
            </a>
            <v-spacer></v-spacer>
            <span class="subheading">
              Let's contribute our code!
              <v-btn icon dark class="mx-3">
                <v-icon size="24px">fab fa-github</v-icon>
              </v-btn>
            </span>
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
a {
  text-decoration:none;
}
</style>
