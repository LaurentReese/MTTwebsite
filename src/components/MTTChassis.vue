<template>
  <div class="MTTChassis">
    <h1>{{ msg1 }}</h1>
    <h1>{{ msg2 }}</h1>
    <img alt="FENETRES PLIANTES PISCINE" src="../assets/FENETRES PLIANTES PISCINE.jpg" img width="20%">
    <img alt="FENETRE PISCINE 2" src="../assets/FENETRE PISCINE 2.jpg" img width="20%">
    <img alt="FENETRES CORSE" src="../assets/FENETRES CORSE.jpg" img width="20%">    
    <img alt="FENETRES PLIANTES 5VTX" src="../assets/FENETRES PLIANTES 5VTX.jpg" img width="20%">
    <pre></pre>
    <img alt="FENETRES PLIANTES 6VTX" src="../assets/FENETRES PLIANTES 6VTX.jpg" img width="20%">        
    <img alt="PLIANTES 2" src="../assets/PLIANTES 2.jpg" img width="33%">
    <hr>    
    <video  ref="videoFenetresAccordeon"
            src="../assets/Fenetres Accordéon Méditerranée Techniques Travaux.webm" 
            width="800">
    </video>
    <!--hr-->
    <button :disabled="isPlaying" @click="play">Jouer</button>
    <button :disabled="!isPlaying" @click="stop">Arrêter</button>      


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
export default {
  name: 'MTTChassis',
  props: {
    msg1 : String,
    msg2 : String
  },
  data : function () {
    return {isPlaying: false}
  },
  methods: {
    play() {
      this.$refs.videoFenetresAccordeon.play()
      this.isPlaying = true
    },
    stop() {
      this.$refs.videoFenetresAccordeon.pause()
      this.isPlaying = false
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
