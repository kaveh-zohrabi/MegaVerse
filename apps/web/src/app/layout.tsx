import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'MegaVerse | پیامرسان مدرن',
  description: 'پیامرسان مدرن و سریع با طراحی زیبا',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="fa" dir="rtl">
      <body>{children}</body>
    </html>
  );
}
