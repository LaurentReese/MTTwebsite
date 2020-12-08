<template>
  <div class="MTTAdmin">
    <button v-on:click="postMTTAdmin">admin</button>
    <span v-if="askPassword">
      {{ espace }}
      Entrez un mot de passe :
      {{ espace }}
      <input type="text" v-model="password" />
      {{ espace }}
      <button v-on:click="postMTTValidatePassword">OK</button>
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
    };
  },

  props: {},

  methods: {
    postMTTAdmin: function () {
      this.password = "";
      this.askPassword = !this.askPassword; // toggle state
    },
    postMTTValidatePassword: function () {
      if (this.password == "") { // no need to get the server's answer in that case
        alert("Le mot de passe est vide");
        return;
      }
      // TO DO : encrypt the password here, it will be decrypted on the server side
      // This is better in case one day HTTP is used instead of HTTPS
      this.askForDatabaseByMail(this.password);
    },
    askForDatabaseByMail(password) {
      var dataFromMTTPassword = {
        password: this.password,
      };
      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/mttAdmin",
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