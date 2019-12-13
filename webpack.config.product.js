const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin"); //可以根据不同模版，定义不同的title
//const ExtractTextPlugin = require('extract-text-webpack-plugin');
const UglifyJSPlugin = require('uglifyjs-webpack-plugin');
//const ManifestPlugin = require('webpack-manifest-plugin'); //追踪所有模块并映射到输出bundle中的
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer')
const CompressionPlugin  = require('compression-webpack-plugin')

/*
 ** 关于CleanWebpackPlugin 的错误引入方法
 ** 1.const CleanWebpackPlugin = require("clean-webpack-plugin");
 **   new CleanWebpackPlugin(['dist'])
 ** 2.const CleanWebpackPlugin = require("clean-webpack-plugin");
 **   new CleanWebpackPlugin(['dist'], {
 **        root: path.resolve(__dirname, '../'),   //根目录
 **   })
 */
//const { CleanWebpackPlugin } = require("clean-webpack-plugin"); //在每次构建前清理 /dist 文件夹，只会生成用到的文件

function getPluginList() {
  var plugins = [
    //new CleanWebpackPlugin(),
    new HtmlWebpackPlugin({
      //可以针对不同的模版，定义不同的名称
      title: "嘿哈",
      template: "./src/index.html",
      chunks: ["littleCai"],
      filename: "index.html"
    }),
    new webpack.DefinePlugin({
      //这可能会对开发模式和发布模式的构建允许不同的行为非常有用
      "process.env": {
        NODE_ENV: JSON.stringify(process.env.NODE_ENV)
      }
    }),
    new webpack.DefinePlugin({
            'process.env':{
                'NODE_ENV': JSON.stringify(process.env.NODE_ENV)
            }
    }),
    //new ExtractTextPlugin({ filename: 'css/[name]-[md5:contenthash:hex:8].css', allChunks: false }),
    //new ExtractTextPlugin({ filename: 'css/[name]-[md5:contenthash:hex:8].css'}),
//    new BundleAnalyzerPlugin({ analyzerPort: 8919 }),
    new UglifyJSPlugin({
      parallel: true,
      cache: true,
      uglifyOptions: {
        output: {
          comments: false,
          beautify: false,
        },
        compress: {
         drop_console: true,
        },
        warnings: false
      },
    }),
    new CompressionPlugin({
      filename: '[path].gz[query]',
      algorithm: 'gzip',
      test: new RegExp('\\.(js|css)$'),
      threshold: 10240,
      minRatio: 0.8
    }),
  ];
    return plugins;
}

module.exports = {
  mode: "development",
  entry: {
    //入口文件
    littleCai: "./src/firstWeb/react/littleCai/index.jsx",
  },
/* optimization: {
    splitChunks: {
      chunks: 'all',
      minSize: 30000,
      maxSize: 0,
      minChunks: 1,
      maxAsyncRequests: 5,
      maxInitialRequests: 3,
      automaticNameDelimiter: '~',
      name: true,
      cacheGroups: {
        vendors: {
          test: /[\\/]node_modules[\\/]/,
          priority: -10,
          filename: 'js/[name]-bundle.js'
        },
        default: {
          minChunks: 2,
          priority: -20,
          reuseExistingChunk: true
        }
      }
    }
  },   */
  //devtool: "inline-source-map",
  devtool: 'false',
  //devServer: {
  //  contentBase: "./dist" //告诉开发服务器(dev server)，在哪里查找文件：
  //},
  plugins: getPluginList(),
  output: {
    filename: "js/[name]-[hash].js",
    chunkFilename: "[name].bundle.js",
    path: path.join(__dirname, "dist"),
    publicPath: '/static/',
    library: "_LittleCaiPageRender"
  },
  module: {
    rules: [
       {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env"],
            outputPath: 'js',
            plugins: [
                     ["import", { "libraryName": "antd", "style": "true"}],
              ],
          }
        }
      },
      {
        test: /\.jsx$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env", "@babel/preset-react"],
             plugins: [
                      ["import", { "libraryName": "antd", "style": "true"}],
               ],     
          }
        }
      },
      {
                test: /\.css?$/,
                use: [
                    'style-loader',
                    'css-loader'
                ]
      }, 
      {
        test: /\.(png|svg|jpg|gif|jpeg)$/,
        use: [{
          loader: "file-loader",
          options: {
            outputPath : 'images'
          }
        },
    //     {
    //                loader: 'image-webpack-loader',
    //                options: {
    //                    bypassOnDebug: true,
   //                 }
    //       }
        ]
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: ["file-loader"]
      },
  ]
  }
};
