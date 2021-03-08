<template>
  <div class="MTTChassis">
    <h1>{{ msg1 }}</h1>
    <h1>{{ msg2 }}</h1>

    <hr />
    <b>
    Monoslide/Chassis à ouverture totale : Vantaux aluminium à rupture de pont thermique, coulissant séparement sur toute la longueur d'un rail unique, et empilables dans une zone de stockage
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
      :width="percent"
    />
    <pre></pre>       
    <video
      muted  = "true"
      autoplay = "true"    
      loop = "true"
      src = "../assets/Fenetres Empilables Mediterranee Techniques Travaux.webm"    
      ref="videoFenetresEmpilables"
      :width="percent"
      horizontal-align=left
      controls
    ></video>
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd2"
      style="width:99%"
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
    Supertherme 80 : Chassis à vantaux pliants empilables en accordéon
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
      :width="percent"
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
      :width="percent"
      horizontal-align=left
      controls
    ></video>
    <!--hr    hr prints a line and makes a carriage return, whereas pre does a simple carriage return-->
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd1"
      style="width:99%"
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
    Bali 60/ThermoSlide 77 : Panneau coulissant suspendu
    </b>
    <pre></pre>      
    Prix : 700 Euros/m2
    <pre></pre>          
    Délai de livraison : 6 à 8 semaines
    <pre></pre>       
    <img
      alt="11_Image_bali-60"
      src="../assets/11_Image_bali-60.jpg"
      img
      :width="percent"
    />
    <pre></pre>    
    <video
      muted  = "true"
      autoplay = "true"    
      loop = "true"
      src = "../assets/sunparadise-monoslide.mp4"    
      ref="sunparadise-monoslide"
      :width="percent"
      horizontal-align=left
      controls
    ></video>
    <pre></pre>    
    <textarea
      class="productText"
      v-model="textProd3"
      style="width:99%"
      :rows="4"
      :readonly="true"
      :autoHeight="autoHeight"
    ></textarea>    
    <pre></pre>
    <a href="https://www.sunparadise.com/EN/Int/Consumer-Product-Detail/SLIDING-SYSTEMS/BALI-60/0/3/11">{{VOIR_AUSSI}}</a>
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
        style="width:99%"
        height="300"
        v-bind:placeholder=PLACE_HOLDER
      ></textarea>

      <br />
      <br />      
      <span>Description du projet en quelques lignes :</span>
      <br />
      <pre></pre>
      <textarea
        v-model="descProjet"
        style="width:99%"
        height="300"
        v-bind:placeholder=PLACE_HOLDER
      ></textarea>

      <pre></pre>
      Date prévue du projet{{ espace }}{{ espace }}{{ espace }}{{ espace }}{{ espace  }}
      <input type="date" v-model="dateProjet" />

      <br />
      <br />
      <span>Commentaire optionnel :</span>
      <br />
      <pre></pre>
      <textarea
        v-model="messClient"
        style="width:99%"
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
    <input class="comment" type="text" v-model="comment"/>
  </div>
</template>

<script>
import axios from "axios";
//import Vue from 'vue'
// import VeeValidate from 'vee-validate'
/* eslint-disable */
//Vue.use(VeeValidate)

export default {
  mounted: function () {
    this.myStyle = "width:" + this.percent; // cannot do that at declaration time
  },

  name: "MTTChassis",

  data: function () {
    return {
      percent: "100%", // change it to cascade change all
      // use myStyle with something like :style="myStyle" if necessary
      myStyle: "", // will be initialized at "mounted time" by using the percent variable
      PLACE_HOLDER: "Ajoutez une ou plusieurs lignes",
      VOIR_AUSSI: "Voir Aussi...",
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
      descProjet: "",
      dateProjet: "",
      commment: "",
      errors: [],
      myUrl: "",
      textProd1:
        "Les cloisons pliantes suspendues Supertherme 80 offrent une multitude de possibilités d'exécution. Pliables vers la gauche, la droite, centrale ou bilatérales, vers l'intérieur ou l'extérieur. Les cloisons pliantes peuvent être réunies dans un angle avec un poteau fixe ou mobile. Largeur de 60 à 1200 mm, Hauteur de 1000 à 2700 mm. Vitrage de 24 à 62mm d'épaisseur. Les cadres ont un coefficient Uf de 1.8, combinés avec les vitrages adéquats ils permettent d'obtenir un coefficient Uw entre 0.8 et 1.4 Watt/m/K. Très facile à utiliser en rénovation.",
      textProd2:
        "Combine fonctionnalité et flexibilité d'une manière unique: ouverture sans limites, même autour des coins, simplement en faisant coulisser la paroi sur le rail fixe au sol. Ce système avec rupture thermique permet de faire coulisser un nombre illimité de parois et de les empiler dans un coin. Ceci peut être combiné avec une porte ouvrante pour les entrées de tous les jours. La possibilité de faire coulisser les parois autour des coins allant jusqu'à 90° permet une grande liberté de design. Principales caractéristiques du système: flexibilité de par l'application d'option de rotation, Système en aluminium avec rupture de pont thermique, Angles de l'élément entre 90° et 180°, Largeur de l'élément: jusqu'à 1300mm, Hauteur de l'élément: jusqu'à 2700mm, Epaisseur du verre: jusqu'à 40mm.",
      textProd3:
        "Porte coulissante suspendue sur un à trois rails, à haut rendement thermique et phonique. Les éléments se soulèvent de 6mm par rotation manuelle de 180° de la poignée à levier. Les joints en caoutchouc sont ainsi libérés et l'élément peut coulisser facilement. Vitrage de 16 à 42mm. Dimensions des vantaux : largeur max = 2500 mm, hauteur max = 2800 mm.",
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
        descProjet: this.descProjet,
        dateProjet: this.dateProjet,
        comment: this.comment
      };

      console.log(dataFromMTT);
      // SEE  https://www.jenniferbland.com/how-to-use-environment-variables-in-vue-js/
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

.comment {
  color: transparent;
  background-color: transparent;
  text-emphasis-color: transparent;
  text-decoration-color: transparent;
  stop-color: transparent;
  flood-color: transparent;
  color-adjust: transparent;
  lighting-color: transparent;
  caret-color: transparent;
  scrollbar-highlight-color: transparent;
  text-emphasis-color: transparent;
  scrollbar-base-color: transparent;
  scrollbar-face-color: transparent;
  scrollbar-darkshadow-color: transparent;
  scrollbar-3dlight-color: transparent;
  scrollbar-track-color: transparent;
  scrollbar-arrow-color: transparent;
  outline-color: transparent;
  border-inline-end-color: transparent;
  border-inline-start-color: transparent;
  border-block-start-color: transparent;
  border-block-end-color: transparent;
  border-color: transparent;
  border-top-color: transparent;
  border-left-color: transparent;
  border-right-color: transparent;
  border-bottom-color: transparent;
  border-block-color: transparent;
  border-inline-color: transparent;
  margin: none;
  padding: none;
  border: none;
  border-width: 0;
  outline: none;
}
</style>
