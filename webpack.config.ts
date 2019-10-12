import * as path from 'path';
import * as webpack from 'webpack';
import * as nodeExternals from 'webpack-node-externals';

const config: webpack.Configuration = {
  entry: './src/bootstrap.ts',
  target: 'node',
  mode: 'production',
  node: {
    __dirname: false,
    __filename: false,
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  externals: [
    nodeExternals(),
  ],
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: 'elastic-amqp-stash.js',
    path: path.resolve(__dirname, 'dist'),
  },
};

export default config;
