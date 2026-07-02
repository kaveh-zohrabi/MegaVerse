/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: '*.megaverse.dev',
      },
    ],
  },
};

module.exports = nextConfig;
