const path = require('path')
const { VueLoaderPlugin } = require('vue-loader');

module.exports = {
  entry: "./index.js",
    output: {
      path: path.resolve(__dirname, '../static'),
      filename: "bundle.js"
    },

    resolve: {
      alias: {
        'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' webpack 1 用
      }
    },

    module: {
      rules: [
        {test: /\.vue$/, loader: 'vue-loader' },
      ]
    },

    plugins:[
      new VueLoaderPlugin()
    ]
  }