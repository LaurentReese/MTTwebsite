<template>
  <div class="MTTChassis">
    <h1>{{ msg1 }}</h1>
    <h1>{{ msg2 }}</h1>
    <hr>
    <img alt="FENETRES PLIANTES PISCINE" src="../assets/FENETRES PLIANTES PISCINE.jpg" img width="40%">
    &nbsp;&nbsp;    
    <!--label-->
      <input type="checkbox" name="chassis1" value="valeur" v-model="chassis1"> Je suis intéressé par ce produit
    <!--/label-->
    <!--pre></pre-->
    <hr>    
    <img alt="FENETRE PISCINE 2" src="../assets/FENETRE PISCINE 2.jpg" img width="40%">
    &nbsp;&nbsp;    
    <input type="checkbox" name="chassis2" value="valeur" v-model="chassis2"> Je suis intéressé par ce produit
    <hr>
    <img alt="FENETRES CORSE" src="../assets/FENETRES CORSE.jpg" img width="40%">    
    &nbsp;&nbsp;    
    <input type="checkbox" name="chassis3" value="valeur" v-model="chassis3"> Je suis intéressé par ce produit
    <hr>
    <img alt="FENETRES PLIANTES 5VTX" src="../assets/FENETRES PLIANTES 5VTX.jpg" img width="40%">
    &nbsp;&nbsp;    
    <input type="checkbox" name="chassis4" value="valeur" v-model="chassis4"> Je suis intéressé par ce produit
    <hr>
    <img alt="FENETRES PLIANTES 6VTX" src="../assets/FENETRES PLIANTES 6VTX.jpg" img width="40%">
    &nbsp;&nbsp;    
    <input type="checkbox" name="chassis5" value="valeur" v-model="chassis5"> Je suis intéressé par ce produit    
    <hr>
    <img alt="PLIANTES 2" src="../assets/PLIANTES 2.jpg" img width="66%">
    &nbsp;&nbsp;
    <input type="checkbox" name="chassis6" value="valeur" v-model="chassis6"> Je suis intéressé par ce produit        
    <hr>
    <video  ref="videoFenetresAccordeon"
            src="../assets/Fenetres Accordéon Méditerranée Techniques Travaux.webm" 
            width="1000">
    </video>
    <!--hr    hr prints a line and makes a carriage return, whereas pre does a simple carriage return-->
    <button class="button btn-primary" v-on:click="playStop"> Jouer / Arrêter </button>    
    <!--button :disabled="isPlaying" @click="play">Jouer</button>
    <button :disabled="!isPlaying" @click="stop">Arrêter</button-->
    &nbsp;&nbsp;
    <input type="checkbox" name="chassis7" value="valeur" v-model="chassis7"> Je suis intéressé par ce produit
    <hr>
    <b>
    Nom{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}
    <input v-model="nom"/><pre></pre>
    Prénom{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}
    <input v-model="prenom"/><pre></pre>
    Téléphone{{espace}}{{espace}}
    <input v-model="telephone"/><pre></pre>
    E-mail{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}{{espace}}
    <input v-model="mail"/>    <pre></pre>
    </b>
    <h1>
    {{espace}}{{espace}}{{espace}}{{espace}}{{espace}}
    {{espace}}{{espace}}{{espace}}{{espace}}{{espace}}    
    <button class="button btn-primary" v-on:click="envoyer"> INFORMER MTT </button>
    </h1>
    <pre></pre>
    <pre></pre>


    <!--p>
      For a guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
    </p>
    <h3>Installed CLI Plugins</h3>
    <ul>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel" target="_blank" rel="noopener">babel</a></li>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank" rel="noopener">eslint</a></li>
    </ul>
    <h3>Essential Links</h3>
    <ul>
      <li><a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a></li>
      <li><a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a></li>
      <li><a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a></li>
      <li><a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a></li>
    </ul>
    <h3>Ecosystem</h3>
    <ul>
      <li><a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a></li>
      <li><a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a></li>
      <li><a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank" rel="noopener">vue-devtools</a></li>
      <li><a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a></li>
      <li><a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a></li>
    </ul-->
  </div>
</template>

<script>

// TO DO : factorize, make arrays for texts, checkboxes, spaces, etc
export default {
  name: 'MTTChassis',
  props: {
    msg1 : String,
    msg2 : String
  },
  data : function () {
    return {isPlaying: false,
            chassis1: false, // TO DO : of course array of bool to regroup all my choices
            chassis2: false, // TO DO : of course array of bool to regroup all my choices
            chassis3: false, // TO DO : of course array of bool to regroup all my choices
            chassis4: false, // TO DO : of course array of bool to regroup all my choices            
            chassis5: false, // TO DO : of course array of bool to regroup all my choices
            chassis6: false, // TO DO : of course array of bool to regroup all my choices
            chassis7: false, // TO DO : of course array of bool to regroup all my choices
            nom : "",
            prenom : "",
            telephone : "",
            mail : "",
            espace : '\xa0'
    }
  },
  methods: {
    play() {
      this.$refs.videoFenetresAccordeon.play()
      this.isPlaying = true
    },
    stop() {
      this.$refs.videoFenetresAccordeon.pause()
      this.isPlaying = false
    },
    playStop() {
      if (this.isPlaying)
        this.$refs.videoFenetresAccordeon.pause()
      else
        this.$refs.videoFenetresAccordeon.play()
      this.isPlaying = !this.isPlaying      
    },
    envoyer() {
      alert("nom=" + this.nom + ",prénom=" + this.prenom + ",téléphone=" + this.telephone + ",mail=" + this.mail)
      // TO DO handle with the array of course
      var produits = []
      if (this.chassis1)
        produits.push(1)
      if (this.chassis2)
        produits.push(2)
      if (this.chassis3)
        produits.push(4)
      if (this.chassis4)
        produits.push(4)        
      if (this.chassis5)
        produits.push(5)
      if (this.chassis6)
        produits.push(6)
      if (this.chassis7)
        produits.push(7)        
        
      alert("intérêt pour les produits : " + produits)
          
    }    
  }  
}
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
