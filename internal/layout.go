package internal

import (
	x "github.com/bloxui/blox"
	"github.com/bloxui/lucide"
)

// PageConfig holds all configuration for building a page
type PageConfig struct {
	Title       string
	Description string
	ActiveRoute string
	Body        []x.Component
	Assets      *x.Assets
}

// DefaultPageConfig returns a PageConfig with sensible defaults
func DefaultPageConfig() PageConfig {
	return PageConfig{
		Title:       "Blox",
		Description: "Blox – a modern function-based HTML component library for Go with compile-time safety",
	}
}

// BuildPage creates a complete HTML page from the config
func BuildPage(config PageConfig) string {
	head := Head(config)
	body := Body(config)

	return "<!DOCTYPE html>\n" + x.Render(x.Html(
		x.Lang("en"),
		x.C(head),
		x.C(body),
	))
}

// buildHead creates the HTML head section
func Head(config PageConfig) x.Component {
	components := []x.Component{
		x.HeadTitle(x.T(config.Title)),
		x.Meta(x.Charset("utf-8")),
		x.Meta(x.Name("viewport"), x.Content("width=device-width, initial-scale=1")),
		x.Meta(x.Name("description"), x.Content(config.Description)),
		x.Meta(x.Name("theme-color"), x.Content("#0b1220")),
		x.Link(x.LinkRel("stylesheet"), x.LinkHref("/assets/styles.css"), x.LinkType("text/css")),
	}

	// Add assets if provided
	if config.Assets != nil && config.Assets.HasAssets() {
		components = append(components, config.Assets.CSS())
		components = append(components, config.Assets.JS())
	}

	var headArgs []x.HeadArg
	for _, component := range components {
		headArgs = append(headArgs, x.C(component))
	}

	return x.Head(headArgs...)
}

// buildBody creates the HTML body section
func Body(config PageConfig) x.Component {
	nav := Navigation(config.ActiveRoute)
	footer := Footer()

	allComponents := []x.Component{nav}
	allComponents = append(allComponents, config.Body...)
	allComponents = append(allComponents, footer)

	var wrapperArgs []x.DivArg
	wrapperArgs = append(wrapperArgs, x.Class("min-h-screen bg-gradient-to-b from-white to-zinc-50 dark:from-zinc-900 dark:to-zinc-950"))
	for _, component := range allComponents {
		wrapperArgs = append(wrapperArgs, x.C(component))
	}

	wrapper := x.Div(wrapperArgs...)
	return x.Body(x.Class("bg-black"), x.C(wrapper))
}

// buildNavigation creates the site navigation header
func Navigation(activeRoute string) x.Component {
	logo := x.A(
		x.Href("/"),
		x.Class("flex items-center gap-2"),
		x.C(x.Span(x.Class("inline-block h-8 w-8 rounded-md bg-gradient-to-br from-violet-500 to-fuchsia-500"))),
		x.C(x.Span(x.T("Blox"), x.Class("text-zinc-900 font-semibold text-lg dark:text-zinc-100"))),
	)

	navigation := x.Nav(
		x.C(navLinkWithIcon("Home", "/", activeRoute == "/", lucide.House(append(lucide.Size("16"), lucide.StrokeWidth("1.5"))...))),
		x.C(navLink("Features", "/features", activeRoute == "/features")),
		x.C(navLink("Docs", "/docs", activeRoute == "/docs")),
		x.C(navLink("Contact", "/contact", activeRoute == "/contact")),
		x.C(navLink("Modal", "/modal", activeRoute == "/modal")),
		x.C(navLink("Tabs", "/tabs", activeRoute == "/tabs")),
	)

	return x.Header(
		x.Class("sticky top-0 z-40 border-b border-zinc-900/5 bg-white/70 backdrop-blur-sm dark:bg-zinc-900/60 dark:border-white/10"),
		x.C(
			x.Div(
				x.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"),
				x.C(
					x.Div(
						x.Class("h-16 flex items-center justify-between"),
						x.C(logo),
						x.C(navigation),
					),
				),
			),
		),
	)
}

// buildFooter creates the site footer
func Footer() x.Component {
	return x.Footer(
		x.Class("bg-zinc-50 border-t border-zinc-900/5 dark:bg-zinc-900/60 dark:border-white/10"),
		x.C(
			x.Div(
				x.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8"),
				x.C(
					x.Div(
						x.Class("text-center text-sm text-zinc-600 dark:text-zinc-300"),
						x.C(x.P(x.T("Made with Blox – the functional HTML component library for Go"))),
					),
				),
			),
		),
	)
}

// navLink creates a standard navigation link
func navLink(text, href string, active bool) x.Component {
	base := "inline-flex h-9 items-center rounded-md px-3 text-sm font-medium transition-colors"
	idle := "text-zinc-700 hover:text-zinc-900 hover:bg-zinc-100 dark:text-zinc-300 dark:hover:text-white dark:hover:bg-zinc-800"
	activeClass := "text-violet-700 bg-violet-50 dark:text-violet-300 dark:bg-violet-900/30"

	cls := base + " " + idle
	if active {
		cls = base + " " + activeClass
	}

	return x.A(x.Href(href), x.T(text), x.Class(cls))
}

// navLinkWithIcon creates a navigation link with an icon
func navLinkWithIcon(text, href string, active bool, icon x.Component) x.Component {
	base := "inline-flex h-9 items-center gap-2 rounded-md px-3 text-sm font-medium transition-colors"
	idle := "text-zinc-700 hover:text-zinc-900 hover:bg-zinc-100 dark:text-zinc-300 dark:hover:text-white dark:hover:bg-zinc-800"
	activeClass := "text-violet-700 bg-violet-50 dark:text-violet-300 dark:bg-violet-900/30"

	cls := base + " " + idle
	if active {
		cls = base + " " + activeClass
	}

	return x.A(x.Href(href), x.Class(cls), x.C(icon), x.T(text))
}

// container wraps content with consistent max-width and padding
func container(children ...x.Component) x.Component {
	var args []x.DivArg
	args = append(args, x.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"))
	for _, child := range children {
		args = append(args, x.C(child))
	}
	return x.Div(args...)
}
