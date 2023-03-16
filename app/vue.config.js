const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  assetsDir: "statics",
  devServer: {
    proxy: {
      "^/api": {
        target: "http://127.0.0.1:3119",
        changeOrigin: true,
      },
    },
  }
})
