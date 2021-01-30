<template>
  <div class="MTTChassis">
    <h1>{{ msg1 }}</h1>
    <h1>{{ msg2 }}</h1>

    <hr />
    <b>
    Monoslide/Chassis à ouverture totale : Ventaux aluminium à rupture de pont thermique, coulissant séparement sur toute la longueur d'un rail unique, et empilables dans une zone de stockage
    </b>
    <pre></pre>      
    Prix : sur devis
    <pre></pre>          
    Délai de livraison : 6 à 8 semaines
    <pre></pre>       
    <img
      alt="FENETRES PLIANTES PISCINE"
      src="../assets/FENETRES PLIANTES PISCINE.jpg"
      img
      width="40%"
    />
    <pre></pre>       
    <video
      muted  = "true"
      autoplay = "true"    
      loop = "true"
      src = "../assets/Fenetres Empilables Mediterranee Techniques Travaux.webm"    
      ref="videoFenetresEmpilables"
      width="40%"
      horizontal-align=left
      controls
    ></video>
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd2"
      style="width:40%"
      :rows="4"
      :readonly="true"
      :autoHeight="autoHeight"
    ></textarea>
    <pre></pre>
    <a href="https://www.archiexpo.fr/prod/sunparadise/product-94948-1252311.html">{{VOIR_AUSSI}}</a>    
    <pre></pre>    
    <input type="checkbox" v-model="produits[0]" />
    {{interet}}


    <hr />
    <b>
    Supertherme 80 : Chassis à ventaux pliants empilables en accordéon
    </b>
    <pre></pre>              
    Prix : 700 Euros/m2
    <pre></pre>              
    Délai de livraison : 6 à 8 semaines
    <pre></pre>
    <img
      alt="FENETRES PLIANTES 5VTX"
      src="../assets/FENETRES PLIANTES 5VTX.jpg"
      img
      width="40%"
    />
    <pre></pre>    
    <!--better to set a percentage like 40%, instead of a hard-coded size       -->    
    <!-- muted  = "true" is VERY important otherwise the autoplay does not work -->
    <video
      muted  = "true"
      autoplay = "true"    
      loop = "true"
      src = "../assets/Fenetres Accordéon Méditerranée Techniques Travaux.webm"    
      ref="videoFenetresAccordeon"
      width="40%"
      horizontal-align=left
      controls
    ></video>
    <!--hr    hr prints a line and makes a carriage return, whereas pre does a simple carriage return-->
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd1"
      style="width:40%"
      :rows="4"
      :readonly="true"
      :autoHeight="autoHeight"
    ></textarea>
    <!--button :disabled="isPlaying" @click="play">Jouer</button>
    <button :disabled="!isPlaying" @click="stop">Arrêter</button-->
    <pre></pre>
    <!--a href="www.sunparadise.com/FR/International/Consumer-Product-Detail/SYSTMES-DE-REPLIER/SUPERTHERM-80/0/2/6"></a-->
    <a href="https://www.sunparadise.com/FR/International/Consumer-Product-Detail/SYSTMES-DE-REPLIER/SUPERTHERM-80/0/2/6">{{VOIR_AUSSI}}</a>    
    <pre></pre>    
    <input type="checkbox" v-model="produits[1]" />
    {{interet}}


    <hr />
    <b>
    Bali 60 : Panneau coulissant suspendu
    </b>
    <pre></pre>      
    Prix : 700 Euros/m2
    <pre></pre>          
    Délai de livraison : 6 à 8 semaines
    <pre></pre>       
    <img
      alt="FENETRES CORSE"
      src="../assets/FENETRES CORSE.jpg"
      img
      width="40%"
    />
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd3"
      style="width:40%"
      :rows="4"
      :readonly="true"
      :autoHeight="autoHeight"
    ></textarea>    
    <pre></pre>
    <a href="https://www.archiexpo.fr/prod/sunparadise/product-62746-569452.html">{{VOIR_AUSSI}}</a>
    <pre></pre>   
    <input type="checkbox" v-model="produits[2]" />
    {{interet}}


    <hr />
    <br />
    <b>
      Nom{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace
      }}{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace
      }}{{ espace }} <input type="text" v-model="nom" />
      *
      <pre></pre>
      Prénom{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace
      }}{{ espace }} <input type="text" v-model="prenom" />
      <pre></pre>
      Téléphone{{ espace }}{{ espace }} <input v-model="telephone" />
      <pre></pre>
      E-mail{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace
      }}{{ espace }}{{ espace }}{{ espace }}{{ espace }}
      <input type="email" v-model="mail" />
      *

      <br />
      <br />      
      <span>Adresse de livraison ou des travaux :</span>
      <br />
      <pre></pre>
      <textarea
        v-model="addrTravaux"
        style="width:600px;"
        height="300"
        v-bind:placeholder=PLACE_HOLDER
      ></textarea>

      <br />
      <br />
      <span>Commentaire optionnel :</span>
      <br />
      <pre></pre>
      <textarea
        v-model="messClient"
        style="width:600px;"
        height="300"
        v-bind:placeholder=PLACE_HOLDER
      ></textarea>

      <p v-if="errors.length">
        <b>SVP corrigez :</b>
        <pre>  <!--jump a line-->
          <template class="laurent" v-for="problem in errors">
            {{problem}}
          </template>    
        </pre>                
        <!--ul>
          <instead of the loop    {{errors}}   >
          <li v-for="(problem, index) in errors"  v-bind:key="index">
            {{problem}}
          </li>
        </ul-->
      </p>
      <h1>
        {{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace }} {{ espace
        }}{{ espace }}{{ espace }}{{ espace }}{{ espace }}
        <button class="button btn-primary" v-on:click="postMTTchassis">
          INFORMER MTT
        </button>
      </h1>
    </b>
    <pre></pre>
    <pre></pre>
  </div>
</template>

<script>
import axios from "axios";
//import Vue from 'vue'
// import VeeValidate from 'vee-validate'
/* eslint-disable */
//Vue.use(VeeValidate)
export default {

  name: "MTTChassis",

  data: function () {
    return {
      PLACE_HOLDER: "Ajoutez une ou plusieurs lignes",
      VOIR_AUSSI : "Voir Aussi...",      
      // STRANGE : I want to make it work with stuff like v-model="produits[0]"
      // but declaring produits: Boolean[3] // does not work, I *must* use the following line instead (don't know why)
      produits: [false, false, false],
      nom: "",
      prenom: "",
      telephone: "",
      mail: "",
      espace: "\xa0",
      messClient: "",
      addrTravaux: "",
      errors: [],
      textProd1 : "Les cloisons pliantes suspendues Supertherme 80 offrent une multitude de possibilités d'exécution. Pliables vers la gauche, la droite, centrale ou bilatérales, vers l'intérieur ou l'extérieur. Les cloisons pliantes peuvent être réunies dans un angle avec un poteau fixe ou mobile. Largeur de 60 à 1200 mm, Hauteur de 1000 à 2700 mm. Vitrage de 24 à 62mm d'épaisseur. Les cadres ont un coefficient Uf de 1.8, combinés avec les vitrages adéquats ils permettent d'obtenir un coefficient Uw entre 0.8 et 1.4 Watt/m/K. Très facile à utiliser en rénovation.",
      textProd2 : "Combine fonctionnalité et flexibilité d'une manière unique: ouverture sans limites, même autour des coins, simplement en faisant coulisser la paroi sur le rail fixé au sol. Ce système avec rupture thermique permet de faire coulisser un nombre illimité de parois et de les empiler dans un coin. Ceci peut être combiné avec une porte tournante pour les entrées de tous les jours. La possibilité de faire coulisser les parois autours des coins allant jusqu'à 90° permet une grande liberté de design. Principales caractéristiques du système: flexibilité de par l'application d'option de rotation, Système en aluminium avec rupture de pont thermique, Angles de l'élément entre 90° et 180°, Largeur de l'élément: jusqu'à 1300mm, Hauteur de l'élément: jusqu'à 2700mm, Epaisseur du verre: jusqu'à 40mm.",
      textProd3 : "Porte coulissante suspendue sur un à trois rails. Les éléments se soulèvent de 6mm par rotation manuelle de 180° de la poignée à levier. Les joints en caoutchouc sont ainsi libérés et l'élément peut coulisser facilement. Vitrage de 16 à 42mm. Dimensions des ventaux : largeur max = 2500mm, hauteur max = 2800 mm."
    };
  },

  props: {
    msg1: String,
    msg2: String,
    interet: String,
  },

  methods: {

    postMTTchassis: function () {
      if (!this.checkForm(this.nom, this.mail)) return;

      var dataFromMTT = {
        nom: this.nom,
        prenom: this.prenom,
        telephone: this.telephone,
        mail: this.mail,
        produits: this.produits,
        addrTravaux: this.addrTravaux,
        messClient: this.messClient,
      };

      console.log(dataFromMTT);

      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/mttChassis",
        data: dataFromMTT,
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
    checkForm: function (name, email) {
      this.errors = [];
      if (!name) this.errors.push("Nom requis.");
      if (!email) {
        this.errors.push("Email requis.");
      } else if (!this.validEmail(email)) {
        this.errors.push("L'email est invalide.");
      }
      if (!this.validProducts(this.produits))
        this.errors.push("Vous devez sélectionner au moins un produit.");
      if (!this.errors.length) return true;
      return false;
    },
    validEmail: function (email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    },
    validProducts: function (produits) {
      for (var produit of produits) {
        // rather use the "of" keyword instead of the "in" keyword (which is deprecated.. gasp ==> the interpreter says nothing)
        if (produit) return true;
      }
      return false;
      /*      for (var i=0;i<produits.length;i++) {
          if (produits[i]) return true;
      }
      return false;*/
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.productText {
  text-align: justify;
  text-justify: auto;
}

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


hr {
	border-top: 5px double #8c8b8b;
  border-style: solid; 
}

/*
hr { 
  height: 30px; 
  border-style: solid; 
  border-color: #8c8b8b; 
  border-width: 1px 0 0 0; 
  border-radius: 20px; 
} 
hr:before { 
  display: block; 
  content: ""; 
  height: 30px; 
  margin-top: -31px; 
  border-style: solid; 
  border-color: #8c8b8b; 
  border-width: 0 0 1px 0; 
  border-radius: 20px; 
}
*/


</style>