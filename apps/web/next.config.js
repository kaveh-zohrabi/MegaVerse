module.exports = {
  reactStrictMode: true,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: '*.megaverse.dev',
      },
    ],
  },
  experimental: {
    serverActions: true,
  },
};
