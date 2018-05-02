var path = require('path');

module.exports = function(env, argv) {
    return {
        entry: './main',

        output: {
            path: path.join(__dirname, 'build'),
            filename: 'main.js'
        },

        resolve: {
            extensions: ['.ts', '.js']
        },

        module: {
            rules: [
                {
                    test: /\.ts$/,
                    use: ['ts-loader']
                }
            ]
        },

        devtool: 'inline-source-map',

        devServer: {
            contentBase: 'public',
            host: '0.0.0.0',
            port: 8081
        }
    }
};

