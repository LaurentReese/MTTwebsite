<template>
  <div id="app">
    <img alt="MTT logo" src="./assets/LOGO MTT HABITAT.png" img width="20%" />
    <MTTAdmin />
    <MTTChassis
      v-bind:msg1="MSG1"
      v-bind:msg2="MSG2"
      v-bind:interet="INTERET"
    />
  </div>
</template>

<script>
import MTTChassis from "./components/MTTChassis.vue";
import MTTAdmin from "./components/MTTAdmin.vue";

export default {
  name: "App",
  components: {
    MTTChassis,
    MTTAdmin,
  },
  data: function () {
    return {
      MSG1: "Bienvenue sur le site de MTT HABITAT",
      MSG2:
        "Nous présentons des solutions spéciales de système à ouverture totale. De conception suédoise, ces produits sont de haute qualité et offrent de grandes performances d'isolations thermique et phonique.",
      INTERET: "Je suis intéressé par ce produit",
    };
  },
};
</script>

<style>
#app {
  background-color: #999999;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  vertical-align: middle;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
