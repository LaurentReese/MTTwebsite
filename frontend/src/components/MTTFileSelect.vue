<template>
  <label class="MTTFileSelect">
    <div class="select-button">
      <span v-if="value">Fichier Produits:</span>
      <!--span v-if="value">Fichier: {{value.name}}</span-->      
      <span v-else>Fichier Produits:</span>
    </div>
    <input type="file" @change="handleFileChange"/>
  </label>
</template>

<script>
export default {
  props: {
    value: File
  },

  methods: {
    handleFileChange(e) {
      this.$emit('input', e.target.files[0])
    }
  }
}
</script>

<style scoped>

.MTTFileSelect > .select-button {
  max-width: 150px;    
  max-height: 50px;
  padding: 1rem;

  color: white;
  background-color: #2EA169;

  border-radius: .3rem;

  text-align: center;
  font-weight: bold;
}

.MTTFileSelect > input[type="file"] {
  display: none;
}

</style>