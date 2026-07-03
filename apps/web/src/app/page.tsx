import Link from 'next/link';

const services = [
  { name: 'Auth Service', lang: 'Go', icon: '🔐', desc: 'JWT, OAuth, secure authentication' },
  { name: 'User Service', lang: 'Go', icon: '👤', desc: 'Profiles, follows, social graph' },
  { name: 'Social Service', lang: 'Java', icon: '📱', desc: 'Posts, comments, reactions' },
  { name: 'Messaging', lang: 'Go', icon: '💬', desc: 'Real-time WebSocket chat' },
  { name: 'AI Service', lang: 'Python', icon: '🤖', desc: 'Embeddings, vector search' },
  { name: 'API Gateway', lang: 'Go', icon: '🌐', desc: 'Routing, rate limiting, proxy' },
];

const features = [
  { title: 'Polyglot', desc: 'Go, Java, Python, TypeScript - each service in its best language', icon: '⚡' },
  { title: 'Microservices', desc: 'Independent, scalable, deployable services', icon: '🏗️' },
  { title: 'Real-time', desc: 'WebSocket support for live messaging', icon: '🔄' },
  { title: 'AI-Powered', desc: 'Built-in ML inference and embeddings', icon: '🧠' },
  { title: 'Type-Safe', desc: 'Shared types across all TypeScript packages', icon: '🛡️' },
  { title: 'Open Source', desc: 'MIT licensed, community driven', icon: '💚' },
];

const stats = [
  { value: '6', label: 'Microservices' },
  { value: '4', label: 'Languages' },
  { value: '3', label: 'SDKs' },
  { value: '∞', label: 'Scalability' },
];

export default function Home() {
  return (
    <div className="min-h-screen bg-grid">
      {/* Hero Section */}
      <section className="relative min-h-screen flex items-center justify-center px-4">
        {/* Background Effects */}
        <div className="absolute inset-0 overflow-hidden">
          <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-purple-500/20 rounded-full blur-3xl animate-pulse-slow" />
          <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-pink-500/20 rounded-full blur-3xl animate-pulse-slow" style={{ animationDelay: '2s' }} />
          <div className="absolute top-1/2 left-1/2 w-96 h-96 bg-cyan-500/20 rounded-full blur-3xl animate-pulse-slow" style={{ animationDelay: '4s' }} />
        </div>

        <div className="relative z-10 text-center max-w-5xl mx-auto">
          <div className="animate-float mb-8">
            <span className="text-8xl">🚀</span>
          </div>
          
          <h1 className="text-6xl md:text-8xl font-bold mb-6">
            <span className="gradient-text">MegaVerse</span>
          </h1>
          
          <p className="text-xl md:text-2xl text-gray-400 mb-8 max-w-2xl mx-auto">
            A modular, polyglot software ecosystem built for scale. 
            One platform, infinite possibilities.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link href="#services" className="px-8 py-4 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl font-semibold text-lg hover:opacity-90 transition-opacity glow">
              Explore Services
            </Link>
            <a href="https://github.com/kaveh-zohrabi/MegaVerse" target="_blank" rel="noopener noreferrer" className="px-8 py-4 glass rounded-xl font-semibold text-lg hover:bg-white/10 transition-colors">
              View on GitHub
            </a>
          </div>

          <div className="mt-16 grid grid-cols-2 md:grid-cols-4 gap-8">
            {stats.map((stat) => (
              <div key={stat.label} className="text-center">
                <div className="text-4xl md:text-5xl font-bold gradient-text">{stat.value}</div>
                <div className="text-gray-500 mt-2">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-24 px-4">
        <div className="max-w-6xl mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold text-center mb-16">
            Why <span className="gradient-text">MegaVerse</span>?
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {features.map((feature) => (
              <div key={feature.title} className="glass rounded-2xl p-8 hover:bg-white/10 transition-all duration-300 group">
                <span className="text-4xl mb-4 block group-hover:scale-110 transition-transform">{feature.icon}</span>
                <h3 className="text-xl font-bold mb-2">{feature.title}</h3>
                <p className="text-gray-400">{feature.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Services Section */}
      <section id="services" className="py-24 px-4">
        <div className="max-w-6xl mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold text-center mb-16">
            <span className="gradient-text">Services</span>
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {services.map((service) => (
              <div key={service.name} className="glass rounded-2xl p-6 hover:glow transition-all duration-300">
                <div className="flex items-center gap-3 mb-4">
                  <span className="text-3xl">{service.icon}</span>
                  <div>
                    <h3 className="text-lg font-bold">{service.name}</h3>
                    <span className="text-sm text-purple-400">{service.lang}</span>
                  </div>
                </div>
                <p className="text-gray-400">{service.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Tech Stack Section */}
      <section className="py-24 px-4">
        <div className="max-w-4xl mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold text-center mb-16">
            Built with <span className="gradient-text">Modern Tech</span>
          </h2>

          <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
            {[
              { name: 'Go', icon: '🔵' },
              { name: 'Java', icon: '☕' },
              { name: 'Python', icon: '🐍' },
              { name: 'TypeScript', icon: '📘' },
              { name: 'Next.js', icon: '⚛️' },
              { name: 'MySQL', icon: '🗄️' },
              { name: 'Docker', icon: '🐳' },
              { name: 'Tailwind', icon: '🎨' },
            ].map((tech) => (
              <div key={tech.name} className="glass rounded-xl p-6 text-center hover:bg-white/10 transition-colors">
                <span className="text-3xl mb-2 block">{tech.icon}</span>
                <span className="text-sm text-gray-400">{tech.name}</span>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-24 px-4">
        <div className="max-w-4xl mx-auto text-center">
          <div className="glass rounded-3xl p-12 glow">
            <h2 className="text-4xl font-bold mb-6">Ready to Build?</h2>
            <p className="text-gray-400 mb-8 text-lg">
              Start building with MegaVerse today. Open source, production ready.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <a href="https://github.com/kaveh-zohrabi/MegaVerse" target="_blank" rel="noopener noreferrer" className="px-8 py-4 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl font-semibold text-lg hover:opacity-90 transition-opacity">
                Get Started
              </a>
              <Link href="#services" className="px-8 py-4 glass rounded-xl font-semibold text-lg hover:bg-white/10 transition-colors">
                Learn More
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 px-4 border-t border-white/10">
        <div className="max-w-6xl mx-auto flex flex-col md:flex-row justify-between items-center gap-4">
          <div className="flex items-center gap-2">
            <span className="text-2xl">🚀</span>
            <span className="font-bold text-xl">MegaVerse</span>
          </div>
          <p className="text-gray-500">© 2026 MegaVerse. MIT License.</p>
          <a href="https://github.com/kaveh-zohrabi/MegaVerse" target="_blank" rel="noopener noreferrer" className="text-gray-400 hover:text-white transition-colors">
            GitHub →
          </a>
        </div>
      </footer>
    </div>
  );
}
