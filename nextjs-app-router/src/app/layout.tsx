import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Navigation from "@/components/Navigation";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
});

export const metadata: Metadata = {
  title: {
    default: "Next.js App Router 示例",
    template: "%s | Next.js App Router"
  },
  description: "使用 Next.js App Router 构建的现代化 Web 应用",
  keywords: ["Next.js", "React", "TypeScript", "App Router"],
  authors: [{ name: "Developer" }],
  creator: "Next.js Developer",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <body className={`${inter.variable} font-sans antialiased bg-background text-foreground`}>
        <div className="min-h-screen flex flex-col">
          <Navigation />
          <main className="flex-1">
            {children}
          </main>
          <footer className="border-t bg-muted/50 py-6 text-center text-sm text-muted-foreground">
            <p>&copy; 2024 Next.js App Router 示例. 保留所有权利.</p>
          </footer>
        </div>
      </body>
    </html>
  );
}
