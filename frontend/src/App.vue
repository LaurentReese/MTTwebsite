<template>
  <div id="app">
    <img alt="MTT logo" src="./assets/MTT.jpg" img width="5%">
    <MTTChassis msg1="Bienvenue sur le site de MÉDITERRANÉE TECHNIQUES ET TRAVAUX" msg2="Présentation des chassis de fenêtre"/>
  </div>
</template>

<script>
import MTTChassis from './components/MTTChassis.vue'

export default {
  name: 'App',
  // TODO : stuff like const msg1 = "Bienvenue sur le site de MÉDITERRANÉE TECHNIQUES ET TRAVAUX"
  // TODO in mounted() ?
  components: {
    MTTChassis
  },
  data : function() {
    return {}
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  vertical-align: middle;  
  color: #2c3e50;
  margin-top: 60px;
}
</style>
