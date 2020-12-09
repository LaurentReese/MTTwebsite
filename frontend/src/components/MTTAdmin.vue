<template>
  <div class="MTTAdmin">
    {{ espace }} {{ espace }}
    <button v-on:click="MTTAdmin">admin</button>
    <span v-if="askPassword">
      {{ espace }}
      Entrez un mot de passe :
      {{ espace }}
      <input type="password" v-model="password" />
      {{ espace }}
      <br /><br />
      <button v-on:click="MTTDatabaseAction">
        Recevoir la database par mail
      </button>
      <br /><br />
      <button v-on:click="MTTJsonAction">Mettre à jour les produits</button>
    </span>
  </div>
</template>

<script>
import axios from "axios";
//import Vue from 'vue'
// import VeeValidate from 'vee-validate'
/* eslint-disable */
//Vue.use(VeeValidate)
export default {
  name: "MTTAdmin",

  data: function () {
    return {
      askPassword: false,
      password: "",
      espace: "\xa0",
    };
  },
  methods: {
    MTTAdmin: function () {
      this.password = "";
      this.askPassword = !this.askPassword; // toggle state
    },

    MTTDatabaseAction: function () {
      if (this.password == "") {
        // no need to get the server's answer in that case
        alert("Le mot de passe est vide");
        return;
      }
      var dataFromMTTPassword = {
        password: this.password,
      };
      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/mttDatabaseAction",
        data: dataFromMTTPassword,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          console.log(result.data);
          alert(result.data["messageServer"]);
        })
        .catch((error) => {
          alert(error);
          console.error(error);
          alert("Serveur MTT indisponible");
        });
    },

    MTTJsonAction: function () {
      if (this.password == "") {
        // no need to get the server's answer in that case
        alert("Le mot de passe est vide");
        return;
      }
      // Lire un .json des produits
      // Vérifier la syntaxe du .json
      // L'envoyer au serveur GOLANG
      // L'intégrer dans la database (le code est déjà fait)

      // appeler
      /*
            axios({
        method: "POST",
        url: "http://127.0.0.1:8090/mttJsonAction",
        data: dataFromMTTPassword,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          console.log(result.data);
          alert(result.data["messageServer"]);
        })
        .catch((error) => {
          alert(error);
          console.error(error);
          alert("Serveur MTT indisponible");
        });
        */
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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