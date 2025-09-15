# BloxUI Starter (Sample App)

A minimal, idiomatic Go starter that showcases BloxUI (blox + ui) with a simple MVC-ish structure suitable for backend/full‑stack developers.

## Layout (Blueprint)

- `cmd/server/main.go`: entry point
- `internal/app`: app wiring (repos, services, handlers, router, middleware)
- `internal/httpx`: small middleware and handler chain
- `internal/handlers`: HTTP handlers (controllers)
- `internal/views`: UI composition with `blox` + `ui` (pages, layout)
- `internal/service`: application services (use cases)
- `internal/repo`: infrastructure (in-memory repo example)
- `internal/domain`: domain types and repository interfaces
- `starter.go`: exports `Routes()` to present a clean public surface

Guidelines:
- Keep packages shallow (one level under `internal/`).
- Split by responsibility: views don’t import services; handlers are the boundary.
- Add more features by adding files in existing packages (e.g., `handlers/orders.go`).

## Run

The module uses local `replace` directives to point to `../..` copies of blox, ui, and icons. From this folder:

```
go run ./cmd/server
```

Then open http://localhost:8080

## What it shows

- Composition of semantic HTML via `blox` (no templates)
- UI components: buttons, inputs, textarea, checkbox, radio, card; plus a Users feature
- Asset collection: component CSS/JS injected once per page
- App layering: domain → repo/service → handlers → views

Note: Tailwind utilities are provided via CDN in the page `<head>` for convenience. For production, precompile your CSS.
