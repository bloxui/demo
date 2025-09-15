# BloxUI Starter

A production-ready BloxUI application showcasing modern web development with Go. This starter demonstrates type-safe HTML generation, beautiful UI components, and achieves a perfect 100/100/100/100 Lighthouse score.

## ✨ Features

- **Type-safe HTML**: Compile-time validation with zero runtime overhead
- **Modern UI**: Beautiful components with shadcn/ui styling
- **Perfect Lighthouse Score**: 100/100/100/100 across all metrics
- **Production Ready**: Built with performance and SEO best practices
- **Accessible**: WCAG compliant markup and semantic HTML
- **1000+ Icons**: Full Lucide icon set included

## 🏗️ Architecture

- `cmd/server/main.go`: Application entry point
- `internal/app`: Application wiring (dependency injection, routing, middleware)
- `internal/httpx`: HTTP middleware (compression, logging, recovery)
- `internal/handlers`: HTTP handlers (controllers)
- `internal/views`: UI composition with type-safe HTML
- `internal/ui`: Reusable UI components
- `internal/service`: Business logic layer
- `internal/repo`: Data access layer
- `internal/css`: Tailwind CSS compilation and embedding

## 🚀 Quick Start

### Prerequisites

- Go 1.20+
- [Tailwind CSS CLI](https://tailwindcss.com/docs/installation) (for CSS compilation)

### Development

```bash
# Build CSS and start server
tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css -m && go run ./cmd/server/main.go

# Or run separately
tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css -m
go run ./cmd/server/main.go
```

Visit http://localhost:8080

### Production Build

```bash
# Build optimized CSS
tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --minify

# Build binary
go build -o bloxui-starter ./cmd/server
```

## 📁 Project Structure

```
starter/
├── cmd/server/           # Application entry point
├── internal/
│   ├── app/             # App configuration and routing
│   ├── css/             # Tailwind CSS compilation
│   ├── handlers/        # HTTP request handlers
│   ├── httpx/           # HTTP middleware
│   ├── ui/              # Reusable UI components
│   └── views/           # Page templates and layouts
├── go.mod
└── README.md
```

## 🎯 What You'll Learn

- **Type-safe HTML**: Build UIs with compile-time guarantees
- **Component composition**: Reusable, testable UI components
- **Modern Go patterns**: Clean architecture and dependency injection
- **Production deployment**: Best practices for real-world applications

## 🔧 Customization

### Adding New Pages

1. Create handler in `internal/handlers/`
2. Add route in `internal/app/app.go`
3. Create view in `internal/views/`

### Adding UI Components

1. Create component in `internal/ui/`
2. Follow the established patterns for type safety
3. Use semantic CSS classes

### Styling

- Modify `internal/css/index.css` for custom styles
- Rebuild CSS with `tailwindcss` command
- Components use design system tokens

## 🚀 Lighthouse Performance

This starter achieves a perfect **100/100/100/100** Lighthouse score across:
- **Performance**: Fast loading and rendering
- **Accessibility**: WCAG compliant and screen reader friendly
- **Best Practices**: Modern web standards and security
- **SEO**: Search engine optimized markup and metadata

## 📝 License

MIT License - see the BloxUI main repository for details.
