package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
	"github.com/plainkit/starter/internal/domain"
	"github.com/plainkit/starter/internal/ui"
)

// UsersPage renders the users list and includes the modal markup.
func UsersPage(users []domain.User) Node {
	// Enhanced table with better styling and action buttons
	list := Div(
		Class("overflow-auto border border-border rounded-lg"),
		Table(
			Class("w-full text-sm"),
			Thead(
				Tr(Class("border-b border-border"),
					Th(Class("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("ID")),
					Th(Class("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("User")),
					Th(Class("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("Email")),
					Th(Class("text-left p-4 font-medium text-muted-foreground border-r border-border"), T("Status")),
					Th(Class("text-right p-4 font-medium text-muted-foreground"), T("Actions")),
				),
			),
			Tbody(rows(users)...),
		),
	)

	// Enhanced modal with better UX and icons
	modal := ui.Modal(
		Id("add-user"),
		ui.ModalContent(
			ui.ModalHeader(
				ui.ModalTitle(
					Div(
						Class("flex items-center gap-2"),
						Div(
							Class("flex items-center justify-center w-10 h-10 bg-primary/10 rounded-lg"),
							icons.UserPlus(icons.Size("20"), Class("text-primary")),
						),
						T("Add New Team Member"),
					),
				),
				ui.ModalDescription(T("Fill in the details below to invite a new member to your team.")),
			),
			Form(
				Method("post"),
				Action("/users"),
				Class("grid gap-6"),
				Div(
					Class("grid gap-2"),
					ui.Label(
						For("name"),
						T("Full Name"),
					),
					ui.Input(
						Id("name"),
						InputName("name"),
						Placeholder("Enter full name"),
						Required(),
					),
				),
				Div(
					Class("grid gap-2 relative"),
					ui.Label(
						For("email"),
						T("Email Address"),
					),
					ui.Input(
						Id("email"),
						InputName("email"),
						InputType("email"),
						Placeholder("name@company.com"),
						Required(),
					),
				),
				ui.ModalFooter(
					Class("flex items-center gap-4"),
					A(
						Href("#"),
						ui.ButtonClass(ui.ButtonSecondary()),
						T("Cancel"),
					),
					Button(
						ButtonType("submit"),
						ui.ButtonClass(),
						icons.Plus(icons.Size("16")),
						T("Create User"),
					),
				),
			),
		),
	)

	return Div(
		Class("grid gap-6"),
		// Stats cards
		Div(
			Class("grid grid-cols-1 md:grid-cols-3 gap-4 mb-6"),
			ui.Card(
				Class("p-6"),
				Div(
					Class("flex items-center gap-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-primary/10 rounded-lg"),
						icons.Users(icons.Size("24"), Class("text-primary")),
					),
					Div(
						H3(Class("text-2xl font-bold"), T(itoa(len(users)))),
						P(Class("text-muted-foreground text-sm"), T("Total Users")),
					),
				),
			),
			ui.Card(
				Class("p-6"),
				Div(
					Class("flex items-center gap-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-chart-2/10 rounded-lg"),
						icons.Check(icons.Size("24"), Class("text-chart-2")),
					),
					Div(
						H3(Class("text-2xl font-bold"), T(itoa(len(users)))),
						P(Class("text-muted-foreground text-sm"), T("Active Users")),
					),
				),
			),
			ui.Card(
				Class("p-6"),
				Div(
					Class("flex items-center gap-4"),
					Div(
						Class("flex items-center justify-center w-12 h-12 bg-chart-4/10 rounded-lg"),
						icons.Star(icons.Size("24"), Class("text-chart-4")),
					),
					Div(
						H3(Class("text-2xl font-bold"), T("4.8")),
						P(Class("text-muted-foreground text-sm"), T("Avg Rating")),
					),
				),
			),
		),
		// Main users table
		ui.Card(
			ui.CardHeader(
				Div(
					Class("flex items-center justify-between"),
					Div(
						ui.CardTitle(
							Div(
								Class("flex items-center gap-2"),
								icons.Users(icons.Size("20")),
								T("User Management"),
							),
						),
						ui.CardDescription(T("Manage your team members and their permissions.")),
					),
					ui.ModalTrigger(
						Href("#add-user"),
						ui.ButtonClass(),
						icons.UserPlus(icons.Size("16")),
						T("Add User"),
					),
				),
			),
			ui.CardContent(list),
		),
		modal,
	)
}

func rows(users []domain.User) []TbodyArg {
	out := make([]TbodyArg, 0, len(users))
	for _, u := range users {
		tr := Tr(
			Class("hover:bg-muted/50 transition-colors"),
			// ID Column
			Td(
				Class("p-4 border-r border-border"),
				Span(Class("font-mono text-muted-foreground text-xs"), T("#"+itoa(u.ID))),
			),
			// User Column with Avatar
			Td(
				Class("p-4 border-r border-border"),
				Div(
					Class("flex items-center gap-3"),
					Div(
						Class("flex items-center justify-center w-8 h-8 bg-gradient-to-br from-primary/10 to-chart-4/10 rounded-full"),
						icons.User(icons.Size("16"), Class("text-primary")),
					),
					Div(
						P(Class("font-medium"), T(u.Name)),
						P(Class("text-muted-foreground text-xs"), T("Team Member")),
					),
				),
			),
			// Email Column
			Td(
				Class("p-4 border-r border-border"),
				Div(
					Class("flex items-center gap-2"),
					icons.Mail(icons.Size("14"), Class("text-muted-foreground")),
					T(u.Email),
				),
			),
			// Status Column
			Td(
				Class("p-4 border-r border-border"),
				Div(
					Class("inline-flex items-center gap-1 px-2 py-1 bg-chart-2/10 text-chart-2 text-xs rounded-full"),
					icons.Check(icons.Size("12")),
					T("Active"),
				),
			),
			// Actions Column
			Td(
				Class("p-4"),
				Div(
					Class("flex items-center justify-end gap-1"),
					Button(
						ButtonType("button"),
						Class("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-muted transition-colors"),
						Title("View user"),
						icons.Eye(Class("text-muted-foreground"), icons.Size("14")),
					),
					Button(
						ButtonType("button"),
						Class("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-muted transition-colors"),
						Title("Edit user"),
						icons.Pen(icons.Size("14"), Class("text-muted-foreground")),
					),
					Button(
						ButtonType("button"),
						Class("inline-flex items-center justify-center h-8 w-8 rounded-md hover:bg-destructive/10 hover:text-destructive transition-colors"),
						Title("Delete user"),
						icons.Trash2(icons.Size("14"), Class("text-muted-foreground")),
					),
				),
			),
		)
		out = append(out, tr)
	}

	// Add empty state if no users
	if len(users) == 0 {
		emptyRow := Tr(
			Td(
				Class("p-8 text-center"),
				Colspan(5),
				Div(
					Class("flex flex-col items-center gap-3 text-muted-foreground"),
					icons.Users(icons.Size("48"), Class("opacity-50")),
					H3(Class("text-lg font-medium"), T("No users found")),
					P(Class("text-sm"), T("Get started by adding your first team member.")),
				),
			),
		)
		out = append(out, emptyRow)
	}

	return out
}

// helper: integer to string using blox internal pattern
func itoa(i int) string {
	// small helper to avoid importing strconv all over
	// mirrored from blox's internal itoa
	// but here we simply do a minimal import if needed
	// reimplement to avoid public dependency footprint
	// This is deliberately simple.
	// In real code, prefer strconv.Itoa.
	if i == 0 {
		return "0"
	}
	n := i
	b := [20]byte{}
	bp := len(b)
	neg := n < 0
	if neg {
		n = -n
	}
	for n > 0 {
		bp--
		b[bp] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		bp--
		b[bp] = '-'
	}
	return string(b[bp:])
}
