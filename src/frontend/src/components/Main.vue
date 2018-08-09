<template>
  <div class="main">
    <section>
        <v-parallax :src="require('@/assets/hero.jpeg')" height='500'>
          <v-layout
            column
            align-center
            justify-center
            class="white--text"
          >
          </v-layout>
        </v-parallax>
      </section>
  </div>
</template>

<script>
export default {
  name: 'Main',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
