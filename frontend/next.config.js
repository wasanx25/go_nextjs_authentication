// @ts-check

const { StatsWriterPlugin } = require('webpack-stats-plugin')

/**
 * @type {import('next').NextConfig}
 **/
const nextConfig = {
  webpack: (config, options) => {
    const {dev, isServer} = options

    if (!dev && !isServer) {
      config.plugins.push(
        new StatsWriterPlugin({
            filename: '../artifacts/webpack-stats.json',
            stats: {
              assets: true,
              entrypoints: true,
              chunks: true,
              modules: true
            }
          }
        )
      );
    }

    return config;
  }
}

module.exports = nextConfig
