package internal

import (
	x "github.com/bloxui/blox"
	"github.com/bloxui/ui"
)

// heroSection creates the main hero section for the homepage
func heroSection() x.Component {
	return x.Section(
		x.Class("relative isolate overflow-hidden"),
		x.C(
			x.Div(
				x.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-20 pb-24"),
				x.C(
					x.Div(
						x.Class("mx-auto max-w-3xl text-center"),
						x.C(x.H1(x.Class("text-4xl sm:text-5xl font-bold tracking-tight"), x.T("Build HTML components in Go"))),
						x.C(x.P(x.Class("mt-3 text-transparent bg-clip-text bg-gradient-to-r from-violet-500 to-fuchsia-500 font-semibold"), x.T("Type-safe. Functional. Expressive."))),
						x.C(x.P(x.Class("mt-6 text-lg text-zinc-600 dark:text-zinc-300"), x.T("Blox provides a fluent, function-based API for composing HTML components with compile-time safety and zero runtime overhead."))),
						x.C(
							x.Div(
								x.Class("mt-8 flex items-center justify-center gap-3"),
								x.C(ui.Button(ui.Destructive(), x.T("Read the Docs"))),
								x.C(ui.Button(ui.Outline(), x.T("See Features"))),
							),
						),
					),
				),
			),
		),
	)
}

// featuresGrid creates a grid of feature cards
func featuresGrid() x.Component {
	features := []struct {
		title       string
		description string
	}{
		{"Component Functions", "Compose HTML components with fluent, chainable functions for maximum expressiveness."},
		{"Type-Safe Components", "Catch HTML structure errors at build time with Go's type system."},
		{"Zero Runtime Overhead", "Generate clean HTML strings with no reflection or dynamic dispatch."},
		{"Modern HTML Components", "Full coverage of HTML5 elements with proper attributes and semantics."},
		{"Developer Experience", "Works seamlessly with Go tooling, testing, and refactoring."},
		{"Modular Architecture", "Each tag is self-contained with its own file for easy maintenance."},
	}

	var cards []x.Component
	for _, feature := range features {
		card := ui.Card(
			x.C(ui.CardHeader(
				x.C(ui.CardTitle(x.T(feature.title))),
			)),
			x.C(ui.CardContent(
				x.C(ui.CardDescription(x.T(feature.description))),
			)),
		)
		cards = append(cards, card)
	}

	var gridArgs []x.DivArg
	gridArgs = append(gridArgs, x.Class("grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"))
	for _, card := range cards {
		gridArgs = append(gridArgs, x.C(card))
	}

	return x.Section(
		x.Class("py-16"),
		x.C(
			x.Div(
				x.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"),
				x.C(x.Div(gridArgs...)),
			),
		),
	)
}

// pageHeader creates a standard page header with title and description
func pageHeader(title, description string) x.Component {
	return x.Div(
		x.Class("py-6"),
		x.C(x.H1(x.T(title), x.Class("text-4xl font-bold tracking-tight"))),
		x.C(x.P(x.T(description), x.Class("mt-4 text-lg text-zinc-600 dark:text-zinc-300"))),
	)
}

// proseContent creates a prose-styled content area for documentation
func proseContent(children ...x.Component) x.Component {
	var args []x.DivArg
	args = append(args, x.Class("prose prose-zinc dark:prose-invert max-w-none"))
	for _, child := range children {
		args = append(args, x.C(child))
	}
	return x.Div(args...)
}
