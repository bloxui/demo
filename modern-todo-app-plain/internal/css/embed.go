package css

import _ "embed"

// TailwindCSS contains the compiled Tailwind output for the todo example.
//
//go:embed output.css
var TailwindCSS string
