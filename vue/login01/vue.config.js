const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false
})
module.exports = {
  devServer:{
    open: true,
    port:3000
  },
  lintOnSave:false
}
