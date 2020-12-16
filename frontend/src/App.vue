<template>
  <div id="app">
    <img alt="MTT logo" src="./assets/LOGO MTT HABITAT.png" img width="5%">
    <MTTAdmin />
    <MTTChassis msg1="Bienvenue sur le site de MÉDITERRANÉE TECHNIQUES ET TRAVAUX" msg2="Présentation des chassis de fenêtre" interet="Je suis intéressé par ce produit" />
  </div>
</template>

<script>
import MTTChassis from './components/MTTChassis.vue'
import MTTAdmin from './components/MTTAdmin.vue'

export default {
  name: 'App',
  // TODO : stuff like const msg1 = "Bienvenue sur le site de MÉDITERRANÉE TECHNIQUES ET TRAVAUX"
  // TODO in mounted() ?
  components: {
    MTTChassis,
    MTTAdmin
  },
  data : function() {
    return {
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
