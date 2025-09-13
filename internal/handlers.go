package internal

import (
	"fmt"
	"net/http"

	x "github.com/bloxui/blox"
	"github.com/bloxui/ui/css"
)

// HomeHandler handles the homepage route
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	config := DefaultPageConfig()
	config.Title = "Blox – Build HTML in Go"
	config.ActiveRoute = "/"
	config.Body = []x.Component{
		heroSection(),
		featuresGrid(),
	}

	render(w, BuildPage(config))
}

// FeaturesHandler handles the features page route
func FeaturesHandler(w http.ResponseWriter, _ *http.Request) {
	config := DefaultPageConfig()
	config.Title = "Features – Blox"
	config.ActiveRoute = "/features"
	config.Body = []x.Component{
		container(
			pageHeader("Features", "Everything you need to build modern HTML components in Go."),
		),
		featuresGrid(),
	}

	render(w, BuildPage(config))
}

// DocsHandler handles the documentation page route
func DocsHandler(w http.ResponseWriter, _ *http.Request) {
	config := DefaultPageConfig()
	config.Title = "Docs – Blox"
	config.ActiveRoute = "/docs"
	config.Body = []x.Component{
		container(
			x.Div(
				x.Class("py-12"),
				x.C(
					proseContent(
						x.H1(x.T("Documentation")),
						x.P(x.T("Blox is a modern HTML component library for Go that uses functions instead of structs for a more expressive and type-safe experience.")),
						x.H2(x.T("Quickstart")),
						x.Pre(
							x.Class("mt-0"),
							x.C(x.Code(x.Class("language-go"), x.T("page := x.Div(\n    x.Class(\"p-6\"),\n    x.Child(x.H1(x.Text(\"Hello, World!\")))\n)\nhtml := x.Render(page)"))),
						),
						x.H2(x.T("Component Functions")),
						x.P(x.T("Use fluent, chainable functions with compile-time type safety to build components. Each HTML element has its own dedicated options.")),
						x.H2(x.T("Routing")),
						x.P(x.T("Works seamlessly with Go's standard library net/http and any Go web framework.")),
					),
				),
			),
		),
	}

	render(w, BuildPage(config))
}

// ContactHandler handles the contact form page route
func ContactHandler(w http.ResponseWriter, _ *http.Request) {
	config := DefaultPageConfig()
	config.Title = "Contact – Blox"
	config.ActiveRoute = "/contact"
	config.Body = []x.Component{
		container(ContactForm()),
	}

	render(w, BuildPage(config))
}

// ModalHandler handles the modal demo page route
func ModalHandler(w http.ResponseWriter, _ *http.Request) {
	// Collect assets from modal components
	assets := x.NewAssets()
	modalContent := ModalDemo()
	assets.Collect(modalContent...)

	config := DefaultPageConfig()
	config.Title = "Modals – Blox"
	config.ActiveRoute = "/modal"
	config.Assets = assets
	config.Body = modalContent

	render(w, BuildPage(config))
}

// TabsHandler handles the tabs demo page route
func TabsHandler(w http.ResponseWriter, _ *http.Request) {
	// Collect assets from tab components
	assets := x.NewAssets()
	tabsContent := TabsDemo()
	assets.Collect(tabsContent...)

	config := DefaultPageConfig()
	config.Title = "Tabs – Blox"
	config.ActiveRoute = "/tabs"
	config.Assets = assets
	config.Body = tabsContent

	render(w, BuildPage(config))
}

// render writes HTML response to the client
func render(w http.ResponseWriter, htmlOut string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(htmlOut))
}

// CssHandler serves embedded Tailwind CSS
func CssHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	_, _ = fmt.Fprint(w, css.TailwindCSS)
}
