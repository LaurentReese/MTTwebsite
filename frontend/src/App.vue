<template>
  <div id="app">
    <img alt="MTT logo" src="./assets/LOGO MTT HABITAT.png" img width="5%">
    <MTTAdmin />
    <MTTChassis v-bind:msg1=MSG1 v-bind:msg2=MSG2 v-bind:interet=INTERET />    
  </div>
</template>

<script>
import MTTChassis from './components/MTTChassis.vue'
import MTTAdmin from './components/MTTAdmin.vue'

export default {
  name: 'App',
  components: {
    MTTChassis,
    MTTAdmin
  },
  data : function() {
    return {
      MSG1 : "Bienvenue sur le site de MÉDITERRANÉE TECHNIQUES ET TRAVAUX",
      MSG2 : "Présentation des chassis de fenêtre",
      INTERET : "Je suis intéressé par ce produit"
    }
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
