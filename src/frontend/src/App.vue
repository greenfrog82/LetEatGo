<template>
  <div id="app">
    <v-app light>
      <v-toolbar flat>
        <v-toolbar-title>
          <a href='/'>
            <b class="black--text">LetEatGo</b>
          </a>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn flat to="/shop">shop</v-btn>
          <v-btn flat to="/discount">discount</v-btn>
          <v-btn flat>
            <v-icon size="24px">fas fa-user</v-icon>
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <v-content>
        <router-view/>
      </v-content>
      <v-footer height="auto">
        <v-card flat title class="flex" style="background-color:#3e3e3e">
          <!-- https://vuetifyjs.com/en/style/colors -->
          <v-card-title class="grey--text text--lighten-1">
            <a href='/' class="subheading">
              <b class="grey--text text--lighten-1">LetEatGo</b>
            </a>
            <v-spacer></v-spacer>
            <span class="subheading">
              Let's contribute our code!
              <v-btn icon dark class="mx-3" href="https://github.com/greenfrog82/LetEatGo">
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
