module.exports = {
    lintOnSave: false,
    devServer: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8081', // 源地址
                changeOrigin: true, // 改变源
            }
        }
    },
    configureWebpack: {

    },
    css: {
        loaderOptions: {
            less: {
                javascriptEnabled: true
            }
        }
    }
}