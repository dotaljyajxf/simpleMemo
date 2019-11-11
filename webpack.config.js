const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin"); //可以根据不同模版，定义不同的title
/*
 ** 关于CleanWebpackPlugin 的错误引入方法
 ** 1.const CleanWebpackPlugin = require("clean-webpack-plugin");
 **   new CleanWebpackPlugin(['dist'])
 ** 2.const CleanWebpackPlugin = require("clean-webpack-plugin");
 **   new CleanWebpackPlugin(['dist'], {
 **        root: path.resolve(__dirname, '../'),   //根目录
 **   })
 */
const { CleanWebpackPlugin } = require("clean-webpack-plugin"); //在每次构建前清理 /dist 文件夹，只会生成用到的文件

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
    })
  ];

  // if (process.env.NODE_ENV !== "dev") {
  //   plugins.push(
  //     new UglifyJSPlugin({
  //       cache: true,
  //       parallel: true,

  //       uglifyOptions: {
  //         // minimize: true,
  //         compress: {
  //           warnings: false
  //         },
  //         output: {
  //           comments: false
  //         }
  //       }
  //     })
  //   );
  // }

  return plugins;
}

module.exports = {
  mode: "development",
  entry: {
    //入口文件
    littleCai: "./src/firstWeb/react/littleCai/index.jsx"
  },
  devtool: "inline-source-map",
  devServer: {
    contentBase: "./dist" //告诉开发服务器(dev server)，在哪里查找文件：
  },
  plugins: getPluginList(),
  output: {
    filename: "[name]-[hash].js",
    chunkFilename: "[name].bundle.js",
    path: path.join(__dirname, "dist"),
    publicPath: '/static/',
    library: "_LittleCaiPageRender"
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      },
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader",
            options: {
              outputPath: 'css'
            }
          },
          {
            loader: "css-loader" // translates CSS into CommonJS
          },
          {
            loader: "less-loader", // compiles Less to CSS
            options: {
              modifyVars: {
                "primary-color": "#1890ff",
                "link-color": "#1890ff",
                "border-radius-base": "2px",
                // or
                hack: `true; @import "your-less-file-path.less";` // Override with less file
              },
              javascriptEnabled: true
            }
          }
        ]
      },
      {
        test: /\.(png|svg|jpg|gif|jpeg)$/,
        use: {
          loader: "file-loader",
          options: {
            outputPath : 'images'
          }
        }
      },
      {
        test: /\.(woff|woff2|eot|ttf|otf)$/,
        use: ["file-loader"]
      },
      {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env"],
            outputPath: 'js'
          }
        }
      },
      {
        test: /\.jsx$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env", "@babel/preset-react"]
          }
        }
      }
    ]
  }
};
