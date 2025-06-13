import Link from "next/link";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowRight, Code, Database, Globe, Zap } from "lucide-react";

export default function HomePage() {
  return (
    <div className="container mx-auto px-4 py-8">
      {/* Hero Section */}
      <section className="text-center py-12 mb-12">
        <h1 className="text-4xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
          欢迎使用 Next.js App Router
        </h1>
        <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto">
          探索现代化的 Next.js App Router 功能，包括服务器组件、路由系统、数据获取等
        </p>
        <div className="flex gap-4 justify-center flex-wrap">
          <Button asChild size="lg">
            <Link href="/features">
              探索功能 <ArrowRight className="ml-2 h-4 w-4" />
            </Link>
          </Button>
          <Button variant="outline" size="lg" asChild>
            <Link href="/blog">查看博客</Link>
          </Button>
        </div>
      </section>

      {/* Features Section */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold text-center mb-10">核心特性</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <Card className="text-center">
            <CardHeader>
              <Zap className="h-12 w-12 mx-auto mb-4 text-yellow-500" />
              <CardTitle>服务器组件</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                默认在服务器端渲染，提供更好的性能和SEO优化
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Globe className="h-12 w-12 mx-auto mb-4 text-blue-500" />
              <CardTitle>基于文件的路由</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                直观的文件系统路由，支持嵌套路由和动态路由
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Database className="h-12 w-12 mx-auto mb-4 text-green-500" />
              <CardTitle>内置数据获取</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                简化的数据获取方式，支持缓存和重新验证
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Code className="h-12 w-12 mx-auto mb-4 text-purple-500" />
              <CardTitle>TypeScript 支持</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                完整的 TypeScript 支持，提供更好的开发体验
              </CardDescription>
            </CardContent>
          </Card>
        </div>
      </section>

      {/* Quick Links */}
      <section>
        <h2 className="text-3xl font-bold text-center mb-10">快速链接</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>功能演示</CardTitle>
              <CardDescription>查看各种 App Router 功能的实际应用</CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild className="w-full">
                <Link href="/features">查看功能</Link>
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>博客文章</CardTitle>
              <CardDescription>阅读关于 Next.js 和现代 Web 开发的文章</CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild className="w-full">
                <Link href="/blog">阅读博客</Link>
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>关于我们</CardTitle>
              <CardDescription>了解更多关于这个项目的信息</CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild className="w-full">
                <Link href="/about">了解更多</Link>
              </Button>
            </CardContent>
          </Card>
        </div>
      </section>
    </div>
  );
}
