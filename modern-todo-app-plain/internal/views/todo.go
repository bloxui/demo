package views

import (
	"fmt"
	"strings"

	"modern_todo_plain/internal/store"

	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
)

type PageData struct {
	Todos  []store.Todo
	Filter store.Filter
	Stats  store.Stats
}

func TodoPage(data PageData) Component {
	return Layout("Modern Todo", appShell(data))
}

func appShell(data PageData) Node {
	return Div(
		Id("todo-app"),
		Class("flex w-full min-h-screen bg-background"),
		todoSidebar(data),
		Div(
			Class("flex-1 flex flex-col"),
			AppHeader(),
			Main(
				Class("flex-1 bg-background p-6"),
				TodoListSection(data),
			),
		),
		AddTodoDialog(data.Filter),
		EditTodoDialog(data.Filter),
	)
}

func todoSidebar(data PageData) Node {
	filters := []struct {
		key   store.Filter
		label string
		icon  Node
		count int
	}{
		{store.FilterAll, "All Tasks", icons.List(icons.Size("18")), data.Stats.Total},
		{store.FilterActive, "Active", icons.Circle(icons.Size("18")), data.Stats.Active},
		{store.FilterCompleted, "Completed", icons.CircleCheck(icons.Size("18")), data.Stats.Completed},
	}

	items := make([]ChildOpt, 0, len(filters))
	for _, f := range filters {
		isActive := f.key == data.Filter
		buttonClass := "filter-button"
		if isActive {
			buttonClass += " is-active"
		}

		items = append(items,
			Child(
				Button(
					ButtonType("button"),
					Class(buttonClass),
					Aria("pressed", fmt.Sprintf("%t", isActive)),
					Custom("hx-get", fmt.Sprintf("/?filter=%s&partial=app", f.key)),
					Custom("hx-target", "#todo-app"),
					Custom("hx-swap", "outerHTML"),
					Custom("hx-push-url", fmt.Sprintf("/?filter=%s", f.key)),
					Div(
						Class("flex items-center gap-3"),
						Span(Class("flex h-9 w-9 items-center justify-center rounded-lg bg-sidebar-muted"), Child(f.icon)),
						Span(Class("font-medium"), T(f.label)),
					),
					Span(
						Id(fmt.Sprintf("count-%s", f.key)),
						Class("filter-count"),
						T(fmt.Sprintf("%d", f.count)),
					),
				),
			),
		)
	}

	percent := completionPercent(data.Stats)

	buttonArgs := make([]DivArg, len(items)+1)
	buttonArgs[0] = Class("space-y-2")
	for i, item := range items {
		buttonArgs[i+1] = item
	}

	return Aside(
		Class("hidden md:block w-72 border-r border-border bg-sidebar text-sidebar-foreground"),
		Child(
			Div(
				Class("p-6 space-y-6"),
				Div(buttonArgs...),
				Div(
					Class("space-y-2 border-t border-sidebar-muted pt-4"),
					Div(
						Class("flex items-center justify-between text-sm"),
						Span(Class("text-sidebar-foreground"), T("Completion")),
						Span(
							Id("completion-label"),
							Class("text-sidebar-accent font-semibold"),
							T(fmt.Sprintf("%d%%", percent)),
						),
					),
					Div(
						Class("h-2 w-full rounded-full bg-sidebar-muted"),
						Div(
							Id("progress-bar"),
							Class("h-2 rounded-full bg-sidebar-accent transition-all"),
							Style(fmt.Sprintf("wdth: %d%%", percent)),
						),
					),
				),
			),
		),
	)
}

func TodoListSection(data PageData) Node {
	listChildren := []ChildOpt{
		Child(Input(InputType("hidden"), Id("todo-current-filter"), InputName("filter"), InputValue(string(data.Filter)))),
	}

	if len(data.Todos) == 0 {
		listChildren = append(listChildren,
			Child(
				Div(
					Class("flex flex-col items-center justify-center rounded-2xl border border-dashed border-border py-16 text-center"),
					Div(
						Class("mb-4 flex h-24 w-24 items-center justify-center rounded-full bg-muted"),
						icons.ShieldCheck(icons.Size("48"), Class("text-muted-foreground")),
					),
					H3(Class("text-lg font-semibold"), T("No tasks found")),
					P(Class("max-w-sm text-sm text-muted-foreground"), T("You're all caught up! Add a new task to get started.")),
				),
			),
		)
	} else {
		for _, todo := range data.Todos {
			listChildren = append(listChildren, Child(todoCard(todo)))
		}
	}

	listArgs := make([]DivArg, len(listChildren)+2)
	listArgs[0] = Id("todo-list")
	listArgs[1] = Class("space-y-3")
	for i, child := range listChildren {
		listArgs[i+2] = child
	}

	return Section(
		Div(listArgs...),
	)
}

func todoCard(todo store.Todo) Node {
	cardClass := "group rounded-xl border border-border bg-card/80 p-5 shadow-sm transition hover:shadow-md"
	if todo.Completed {
		cardClass += " opacity-80"
	}

	metaContent := []ChildOpt{
		Child(
			Div(
				Class("inline-flex items-center gap-1"),
				icons.Calendar(icons.Size("14"), Class("text-muted-foreground")),
				T(todo.CreatedAt.Format("Jan 02")),
			),
		),
		Child(
			Span(
				Class(priorityBadgeClass(todo.Priority)),
				T(capitalize(string(todo.Priority))),
			),
		),
	}

	textContent := []ChildOpt{
		Child(H3(Class(titleClasses(todo.Completed)), T(todo.Title))),
	}
	if todo.Description != "" {
		textContent = append(textContent,
			Child(P(Class(descriptionClasses(todo.Completed)), T(todo.Description))),
		)
	}
	metaArgs := make([]DivArg, len(metaContent)+1)
	metaArgs[0] = Class("flex items-center gap-3 text-xs text-muted-foreground")
	for i, child := range metaContent {
		metaArgs[i+1] = child
	}

	textContent = append(textContent,
		Child(Div(metaArgs...)),
	)

	textArgs := make([]DivArg, len(textContent)+1)
	textArgs[0] = Class("space-y-2")
	for i, child := range textContent {
		textArgs[i+1] = child
	}

	return Article(
		Class(cardClass),
		Div(
			Class("flex items-start gap-4"),
			Button(
				ButtonType("button"),
				Class(toggleButtonClasses(todo.Completed)),
				Custom("hx-post", fmt.Sprintf("/todos/toggle?id=%s", todo.ID)),
				Custom("hx-target", "#todo-app"),
				Custom("hx-swap", "outerHTML"),
				Custom("hx-include", "#todo-current-filter"),
				icons.Check(icons.Size("16")),
			),
			Div(
				Class("flex-1 space-y-3"),
				Div(
					Class("flex items-start justify-between gap-3"),
					Div(textArgs...),
					Div(
						Class("flex items-center gap-1"),
						Button(
							ButtonType("button"),
							Class("inline-flex h-8 w-8 items-center justify-center rounded-lg text-muted-foreground hover:bg-muted"),
							Data("dialog-target", "edit-dialog"),
							Data("edit-id", todo.ID),
							Data("edit-title", todo.Title),
							Data("edit-description", todo.Description),
							Data("edit-priority", string(todo.Priority)),
							icons.PencilLine(icons.Size("16")),
						),
						Button(
							ButtonType("button"),
							Class("inline-flex h-8 w-8 items-center justify-center rounded-lg text-muted-foreground hover:bg-destructive/10 hover:text-destructive"),
							Custom("hx-post", fmt.Sprintf("/todos/delete?id=%s", todo.ID)),
							Custom("hx-target", "#todo-app"),
							Custom("hx-swap", "outerHTML"),
							Custom("hx-confirm", "Delete this task?"),
							Custom("hx-include", "#todo-current-filter"),
							icons.Trash2(icons.Size("16")),
						),
					),
				),
			),
		),
	)
}

func toggleButtonClasses(completed bool) string {
	base := "mt-1 flex h-5 w-5 items-center justify-center rounded border"
	if completed {
		return base + " border-primary bg-primary text-primary-foreground"
	}
	return base + " border-muted-foreground/40 text-muted-foreground"
}

func titleClasses(completed bool) string {
	base := "text-base font-semibold text-card-foreground"
	if completed {
		return base + " line-through text-muted-foreground"
	}
	return base
}

func descriptionClasses(completed bool) string {
	base := "text-sm text-muted-foreground"
	if completed {
		return base + " line-through"
	}
	return base
}

func priorityBadgeClass(priority store.Priority) string {
	switch priority {
	case store.PriorityLow:
		return "priority-low inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	case store.PriorityHigh:
		return "priority-high inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	default:
		return "priority-medium inline-flex items-center rounded-full px-3 py-1 text-xs font-medium"
	}
}

func completionPercent(stats store.Stats) int {
	if stats.Total == 0 {
		return 0
	}
	return int((float64(stats.Completed)/float64(stats.Total))*100 + 0.5)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func AddTodoDialog(filter store.Filter) Node {
	return Dialog(
		Id("add-dialog"),
		Class("modal"),
		Child(
			Form(
				Id("add-form"),
				Class("space-y-4 p-6"),
				Custom("hx-post", "/todos/create"),
				Custom("hx-target", "#todo-app"),
				Custom("hx-swap", "outerHTML"),
				Custom("hx-include", "#todo-current-filter"),
				Custom("hx-on::afterRequest", "if(event.detail.successful){ todoDialogs.closeDialog('add-dialog'); this.reset(); }"),
				H2(Class("text-xl font-semibold"), T("Add New Task")),
				FormLabel(
					Class("block space-y-2"),
					For("add-title"),
					Span(Class("text-sm font-medium"), T("Title")),
					Input(
						Id("add-title"),
						InputName("title"),
						Required(),
						Placeholder("What needs to be done?"),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				FormLabel(
					Class("block space-y-2"),
					For("add-description"),
					Span(Class("text-sm font-medium"), T("Description")),
					Textarea(
						Id("add-description"),
						TextareaName("description"),
						Rows(3),
						Placeholder("Add more details... (optional)"),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				FormLabel(
					Class("block space-y-2"),
					For("add-priority"),
					Span(Class("text-sm font-medium"), T("Priority")),
					Select(
						Id("add-priority"),
						Custom("name", "priority"),
						Child(Option(Custom("value", string(store.PriorityLow)), T("Low"))),
						Child(Option(Custom("value", string(store.PriorityMedium)), Selected(), T("Medium"))),
						Child(Option(Custom("value", string(store.PriorityHigh)), T("High"))),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm"),
					),
				),
				Input(InputType("hidden"), InputName("filter"), InputValue(string(filter))),
				Div(
					Class("flex gap-2 pt-2"),
					Button(
						ButtonType("button"),
						Class("flex-1 rounded-lg border border-border px-4 py-2 text-sm font-medium hover:bg-muted"),
						Data("close-dialog", "add-dialog"),
						T("Cancel"),
					),
					Button(
						ButtonType("submit"),
						Class("flex-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"),
						T("Add Task"),
					),
				),
			),
		),
	)
}

func EditTodoDialog(filter store.Filter) Node {
	return Dialog(
		Id("edit-dialog"),
		Class("modal"),
		Child(
			Form(
				Id("edit-form"),
				Class("space-y-4 p-6"),
				Custom("hx-post", "/todos/update"),
				Custom("hx-target", "#todo-app"),
				Custom("hx-swap", "outerHTML"),
				Custom("hx-include", "#todo-current-filter"),
				Custom("hx-on::afterRequest", "if(event.detail.successful){ todoDialogs.closeDialog('edit-dialog'); }"),
				Input(InputType("hidden"), Id("edit-id"), InputName("id")),
				Input(InputType("hidden"), Id("edit-filter"), InputName("filter"), InputValue(string(filter))),
				H2(Class("text-xl font-semibold"), T("Edit Task")),
				FormLabel(
					Class("block space-y-2"),
					For("edit-title"),
					Span(Class("text-sm font-medium"), T("Title")),
					Input(
						Id("edit-title"),
						InputName("title"),
						Required(),
						Placeholder("What needs to be done?"),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				FormLabel(
					Class("block space-y-2"),
					For("edit-description"),
					Span(Class("text-sm font-medium"), T("Description")),
					Textarea(
						Id("edit-description"),
						TextareaName("description"),
						Rows(3),
						Placeholder("Add more details... (optional)"),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"),
					),
				),
				FormLabel(
					Class("block space-y-2"),
					For("edit-priority"),
					Span(Class("text-sm font-medium"), T("Priority")),
					Select(
						Id("edit-priority"),
						Custom("name", "priority"),
						Child(Option(Custom("value", string(store.PriorityLow)), T("Low"))),
						Child(Option(Custom("value", string(store.PriorityMedium)), T("Medium"))),
						Child(Option(Custom("value", string(store.PriorityHigh)), T("High"))),
						Class("w-full rounded-lg border bg-background px-3 py-2 text-sm"),
					),
				),
				Div(
					Class("flex gap-2 pt-2"),
					Button(
						ButtonType("button"),
						Class("flex-1 rounded-lg border border-border px-4 py-2 text-sm font-medium hover:bg-muted"),
						Data("close-dialog", "edit-dialog"),
						T("Cancel"),
					),
					Button(
						ButtonType("submit"),
						Class("flex-1 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"),
						T("Save Changes"),
					),
				),
			),
		),
	)
}

func RenderFullPage(data PageData) string {
	return Render(TodoPage(data))
}

func RenderAppShell(data PageData) string {
	return Render(appShell(data))
}
