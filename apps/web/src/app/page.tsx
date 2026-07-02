export default function Home() {
  return (
    <main className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-4">MegaVerse</h1>
        <p className="text-lg text-gray-600 mb-8">
          A modular, polyglot software ecosystem built for scale.
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="p-6 border rounded-lg">
            <h2 className="text-xl font-semibold mb-2">Services</h2>
            <ul className="space-y-2 text-gray-600">
              <li>Auth Service (Go)</li>
              <li>User Service (Go)</li>
              <li>Social Service (Java)</li>
              <li>Messaging Service (Go)</li>
              <li>AI Service (Python)</li>
            </ul>
          </div>

          <div className="p-6 border rounded-lg">
            <h2 className="text-xl font-semibold mb-2">Tech Stack</h2>
            <ul className="space-y-2 text-gray-600">
              <li>Next.js + React</li>
              <li>TypeScript</li>
              <li>Tailwind CSS</li>
              <li>MySQL</li>
            </ul>
          </div>
        </div>

        <div className="mt-8 p-6 border rounded-lg">
          <h2 className="text-xl font-semibold mb-2">Quick Start</h2>
          <pre className="bg-gray-100 p-4 rounded text-sm overflow-x-auto">
{`# Start MySQL
docker-compose -f docker-compose.dev.yml up -d

# Install dependencies
npm install

# Run frontend
npm run dev`}
          </pre>
        </div>
      </div>
    </main>
  );
}
