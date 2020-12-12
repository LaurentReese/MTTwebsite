<template>
  <div class="MTTAdmin">
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
      <br /><br />
      <MTTFileSelect v-model="file"></MTTFileSelect>
      <p v-if="file">==> {{ file.path }}{{ file.name }}</p>
    </span>
  </div>
</template>

<script>
// SELECT FILE : see https://www.digitalocean.com/community/tutorials/vuejs-file-select-component
import axios from "axios"
import MTTFileSelect from './MTTFileSelect.vue'
//import Vue from 'vue'
// import VeeValidate from 'vee-validate'
/* eslint-disable */
//Vue.use(VeeValidate)
export default {
  components : {
    MTTFileSelect
  },
  name: "MTTAdmin",

  data: function () {
    return {
      file: null,
      askPassword: false,
      essai: false,
      password: "",
      espace: "\xa0",
    };
  },
  methods: {
    handleFileChange(e) {
      this.$emit("input", e.target.files[0]);
    },
    MTTAdmin: function () {
      this.password = "";
      this.askPassword = !this.askPassword; // toggle state
      this.file = null;
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
      alert(this.file.name);
      let reader = new FileReader();      
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