<template>
  <h3>welcome,来自{{profession}}的{{name}}</h3>
</template>

<script>

export default {
  data(){
    return{
      name:'',
      profession:''
    }
  },

  created(){
    console.log(1)
    this.name=window.sessionStorage.getItem('name')
    this.profession=window.sessionStorage.getItem('profession')
    console.log(this.name)
  },

}
</script>

<style>

</style>