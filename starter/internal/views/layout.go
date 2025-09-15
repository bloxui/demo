package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
)

// Layout wraps content with head, tailwind, and collected component assets.
func Layout(title string, content Node) Component {
	assets := NewAssets()
	assets.Collect(content)
	return baseHTML(title, content, assets)
}

// LayoutWithAssets collects assets from content and additional components
// that may not be reachable by the collector (e.g., nested components).
func LayoutWithAssets(title string, content Node, extras ...Component) Component {
	assets := NewAssets()
	assets.Collect(content)
	if len(extras) > 0 {
		assets.Collect(extras...)
	}
	return baseHTML(title, content, assets)
}

// LayoutWithAssetsProvided renders using a pre-collected assets bundle.
// If assets is nil, it falls back to collecting from content.
func LayoutWithAssetsProvided(title string, content Node, assets *Assets) Component {
	if assets == nil {
		assets = NewAssets()
		assets.Collect(content)
	}
	return baseHTML(title, content, assets)
}

func baseHTML(title string, content Node, assets *Assets) Component {
	return Html(
		Lang("en"),
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Meta(Name("description"), Content("Plain - A modern, type-safe HTML component library for Go with beautiful interfaces and compile-time guarantees.")),
			HeadTitle(T(title)),
			Link(LinkRel("preload"), LinkHref("/assets/styles.css"), LinkType("text/css")),
			Link(LinkRel("stylesheet"), LinkHref("/assets/styles.css")),
			assets.CSS(),
		),
		Body(
			Class("bg-background text-foreground antialiased min-h-screen font-sans"),
			siteHeader(title),
			Main(Class("container mx-auto p-6"), content),
			assets.JS(),
		),
	)
}

func siteHeader(title string) Node {
	return Header(
		Class("border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/80 shadow-sm"),
		Div(
			Class("container mx-auto px-6 h-16 flex items-center justify-between"),
			Div(
				Class("font-bold text-lg tracking-tight inline-flex items-center gap-3"),
				Div(
					Class("flex items-center justify-center w-8 h-8 bg-primary rounded-lg"),
					icons.Zap(icons.Size("18"), Class("text-primary-foreground")),
				),
				Span(Class("text-foreground"), T("Plain")),
				Span(Class("text-muted-foreground text-sm font-normal"), T("/ "+title)),
			),
			Nav(
				Class("flex items-center gap-1 text-sm"),
				A(
					Href("/"),
					Class("inline-flex items-center gap-2 px-3 py-2 rounded-md hover:bg-muted transition-colors"),
					icons.House(icons.Size("16"), Class("text-muted-foreground")),
					T("Home"),
				),
				A(
					Href("/users"),
					Class("inline-flex items-center gap-2 px-3 py-2 rounded-md hover:bg-muted transition-colors bg-muted text-foreground"),
					icons.Users(icons.Size("16")),
					T("Users"),
				),
			),
		),
	)
}
