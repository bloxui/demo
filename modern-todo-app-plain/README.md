# Modern Todo App (Plain)

A Plain/Go port of the Modern Todo App originally built with Next.js. The example demonstrates how to recreate the full experience using server-rendered HTML, htmx interactions, and Tailwind CSS without client-side frameworks.

## Features

- Type-safe HTML using `github.com/plainkit/html`
- Progressive enhancement with htmx for filters, mutations, and dialogs
- Reusable Tailwind design tokens aligned with the original Next.js app
- In-memory store with add, edit, toggle, and delete operations
- Stats sidebar that swaps via out-of-band updates

## Getting Started

```bash
# Build Tailwind CSS
cd examples/modern-todo-app-plain
TAILWIND_DISABLE_TOUCH=1 tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css -m

# Run the server
cd cmd/server
go run .
```

Open http://localhost:8080 to browse the app.

## Project Structure

```
modern-todo-app-plain/
├── cmd/server         # Entry point
├── internal/
│   ├── app/           # HTTP wiring
│   ├── css/           # Tailwind source + embedded output
│   ├── handlers/      # HTTP handlers
│   ├── store/         # In-memory data store
│   └── views/         # Plain components & layouts
├── go.mod
└── README.md
```

Regenerate CSS whenever you change Tailwind classes in the view layer.
