package views

import (
	. "github.com/plainkit/blox"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/ui"
)

func HomePage() Node {
	return Div(
		Class("grid gap-8"),

		// Hero Section
		Div(
			Class("text-center py-12"),
			Div(
				Class("flex items-center justify-center w-16 h-16 bg-gradient-to-br from-primary to-chart-4 rounded-xl mx-auto mb-6"),
				icons.Zap(icons.Size("32"), Class("text-primary-foreground")),
			),
			H1(
				Class("text-4xl font-bold text-foreground mb-4"),
				T("Welcome to Plain"),
			),
			P(
				Class("text-xl text-muted-foreground max-w-2xl mx-auto mb-8"),
				T("A modern, type-safe HTML component library for Go. This demo showcases beautiful interfaces built with compile-time guarantees."),
			),
			Div(
				Class("flex items-center justify-center gap-4"),
				A(
					Href("/users"),
					ui.ButtonClass(),
					Class("flex items-center gap-2"),
					icons.Users(icons.Size("16")),
					T("View Demo"),
				),
				A(
					Href("https://github.com/plainkit/blox"),
					Target("_blank"),
					ui.ButtonClass(ui.ButtonSecondary()),
					Class("flex items-center gap-2"),
					icons.Github(icons.Size("16")),
					T("GitHub"),
				),
			),
		),

		// Stats Section
		Div(
			Class("text-center mb-8"),
			H2(Class("text-2xl font-bold mb-6"), T("Key Features")),
		),
		Div(
			Class("grid grid-cols-1 md:grid-cols-4 gap-4"),
			ui.Card(
				Class("p-6 text-center"),
				Div(
					Class("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg mx-auto mb-3"),
					icons.Code(icons.Size("24"), Class("text-primary")),
				),
				H3(Class("text-2xl font-bold"), T("100%")),
				P(Class("text-muted-foreground text-sm"), T("Type Safe")),
			),
			ui.Card(
				Class("p-6 text-center"),
				Div(
					Class("flex items-center justify-center w-12 h-12 bg-chart-1/10 rounded-lg mx-auto mb-3"),
					icons.Zap(icons.Size("24"), Class("text-chart-1")),
				),
				H3(Class("text-2xl font-bold"), T("0ms")),
				P(Class("text-muted-foreground text-sm"), T("Runtime Overhead")),
			),
			ui.Card(
				Class("p-6 text-center"),
				Div(
					Class("flex items-center justify-center w-12 h-12 bg-chart-3/10 rounded-lg mx-auto mb-3"),
					icons.Blocks(icons.Size("24"), Class("text-chart-3")),
				),
				H3(Class("text-2xl font-bold"), T("50+")),
				P(Class("text-muted-foreground text-sm"), T("Components")),
			),
			ui.Card(
				Class("p-6 text-center"),
				Div(
					Class("flex items-center justify-center w-12 h-12 bg-chart-5/10 rounded-lg mx-auto mb-3"),
					icons.Heart(icons.Size("24"), Class("text-chart-5")),
				),
				H3(Class("text-2xl font-bold"), T("1000+")),
				P(Class("text-muted-foreground text-sm"), T("Beautiful Icons")),
			),
		),

		// Features Grid
		Div(
			Class("text-center mb-8"),
			H2(Class("text-2xl font-bold mb-6"), T("Why Choose Plain")),
		),
		Div(
			Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"),

			// Feature 1: Type Safety
			ui.Card(
				Class("p-6 hover:shadow-lg transition-shadow"),
				Div(
					Class("flex items-center gap-4 mb-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg"),
						icons.Shield(icons.Size("24"), Class("text-primary")),
					),
					H3(Class("text-lg font-semibold"), T("Type Safe")),
				),
				P(
					Class("text-muted-foreground"),
					T("Compile-time validation ensures your HTML is always correct. Invalid combinations fail at build time."),
				),
			),

			// Feature 2: Performance
			ui.Card(
				Class("p-6 hover:shadow-lg transition-shadow"),
				Div(
					Class("flex items-center gap-4 mb-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-chart-2/10 rounded-lg"),
						icons.Zap(icons.Size("24"), Class("text-chart-2")),
					),
					H3(Class("text-lg font-semibold"), T("Lightning Fast")),
				),
				P(
					Class("text-muted-foreground"),
					T("Zero runtime overhead. Pure function calls generate HTML strings at compile time."),
				),
			),

			// Feature 3: Beautiful Design
			ui.Card(
				Class("p-6 hover:shadow-lg transition-shadow"),
				Div(
					Class("flex items-center gap-4 mb-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-chart-4/10 rounded-lg"),
						icons.Palette(icons.Size("24"), Class("text-chart-4")),
					),
					H3(Class("text-lg font-semibold"), T("Beautiful Design")),
				),
				P(
					Class("text-muted-foreground"),
					T("Modern UI components with shadcn/ui styling and 1000+ Lucide icons included."),
				),
			),
		),

		// Code Example
		ui.Card(
			Class("p-8 bg-gradient-to-br from-muted/50 to-accent/50"),
			ui.CardHeader(
				ui.CardTitle(
					Div(
						Class("flex items-center gap-2"),
						icons.Code(icons.Size("24")),
						T("Clean, Readable Code"),
					),
				),
				ui.CardDescription(T("This entire page is built with type-safe Go functions. No templates needed!")),
			),
			ui.CardContent(
				Pre(
					Class("bg-muted p-4 rounded-lg border text-sm overflow-auto font-mono"),
					T(`import (
	. "github.com/plainkit/blox"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/ui"
)

func HomePage() Node {
    return ui.Card(
        ui.CardHeader(
            ui.CardTitle(
                Div(
                    Class("flex items-center gap-2"),
                    icons.Zap(icons.Size("20")),
                    T("Plain Demo"),
                ),
            ),
        ),
        ui.CardContent(
            P(T("Type-safe HTML in Go!")),
        ),
    )
}`),
				),
			),
		),
	)
}
