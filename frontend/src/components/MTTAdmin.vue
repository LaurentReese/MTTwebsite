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
      <hr />
      <pre></pre>
      <input type="file" ref="myFile" @change="selectedFile" />
      <br/>            
      <br/>                  
      <button v-on:click="MTTJsonAction">Mettre à jour les produits</button>
      <br /><br />
      <hr />
    </span>
  </div>
</template>

<script>
// SELECT FILE : see https://www.digitalocean.com/community/tutorials/vuejs-file-select-component
import axios from "axios";

//import MTTFileSelect from './MTTFileSelect.vue'
//import Vue from 'vue'
// import VeeValidate from 'vee-validate'
/* eslint-disable */
//Vue.use(VeeValidate)
export default {
  components: {
    // MTTFileSelect
  },
  name: "MTTAdmin",

  data: function () {
    return {
      askPassword: false,
      password: "",
      espace: "\xa0",
      text: "",
    };
  },
  methods: {
    selectedFile() {
      // see: https://www.raymondcamden.com/2019/05/21/reading-client-side-files-for-validation-with-vuejs
      let file = this.$refs.myFile.files[0];
      if (!file) /* || file.type !== "text/plain") */ {
        this.text = ""
        return;
      }
      let reader = new FileReader();
      reader.readAsText(file, "UTF-8");
      reader.onload = (evt) => {
      this.text = evt.target.result;
      };
      reader.onerror = (evt) => {
        console.error(evt);
      };
    },

    MTTAdmin: function () {
      this.password = "";
      this.askPassword = !this.askPassword; // toggle state
      this.text = ""
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
      //url: "http://127.0.0.1:8090/mttChassis",   // local
      //url: "http://mtt-backend.sloppy.zone:80/mttChassis", // production
      // TO DO : factorize VUE_APP_EXECUTION from App.vue
      var myUrl = "http://127.0.0.1:8090/mttChassis";
      if (process.env.VUE_APP_EXECUTION == "PRODUCTION") {
        myUrl = "http://mtt-backend.sloppy.zone:80/mttChassis";
      }

      axios({
        method: "POST",
        url: myUrl,
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
      if (this.text=="") {
        alert("Pas de fichier json sélectionné");
        return;
      }
      var dataFromMTTJson = {
        password: this.password,
        text: this.text,
      };
      //url: "http://127.0.0.1:8090/mttChassis",   // local
      //url: "http://mtt-backend.sloppy.zone:80/mttChassis", // production
      // TO DO : factorize VUE_APP_EXECUTION from App.vue
      var myUrl = "http://127.0.0.1:8090/mttChassis";
      if (process.env.VUE_APP_EXECUTION == "PRODUCTION") {
        myUrl = "http://mtt-backend.sloppy.zone:80/mttChassis";
      }

      axios({
        method: "POST",
        url: myUrl,
        data: dataFromMTTJson,
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
    }
  }
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