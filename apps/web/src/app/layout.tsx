import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'MegaVerse - Modular Software Ecosystem',
  description: 'A modular, polyglot software ecosystem built for scale. Go, Java, Python, TypeScript.',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
