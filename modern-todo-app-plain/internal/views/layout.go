package views

import (
	. "github.com/plainkit/html"
	icons "github.com/plainkit/icons/lucide"
)

const dialogController = `(() => {
  const getFilterValue = () => {
    const filter = document.getElementById('todo-current-filter');
    if (!filter) return '';
    return filter.value || filter.getAttribute('value') || '';
  };

  const applyFilterToForms = (value) => {
    const addFilter = document.querySelector('#add-form input[name="filter"]');
    if (addFilter) addFilter.value = value;
    const editFilter = document.getElementById('edit-filter');
    if (editFilter) editFilter.value = value;
  };

  const prefillEditForm = (dataset, filterValue) => {
    const form = document.getElementById('edit-form');
    if (!form) return;
    const idField = document.getElementById('edit-id');
    if (idField) idField.value = dataset.editId || '';
    const titleField = document.getElementById('edit-title');
    if (titleField) titleField.value = dataset.editTitle || '';
    const descField = document.getElementById('edit-description');
    if (descField) descField.value = dataset.editDescription || '';
    const priorityField = document.getElementById('edit-priority');
    if (priorityField) priorityField.value = dataset.editPriority || 'medium';
    const filterField = document.getElementById('edit-filter');
    if (filterField) filterField.value = filterValue;
  };

  const setupDialogControls = () => {
    const openButtons = document.querySelectorAll('[data-dialog-target]');
    openButtons.forEach(btn => {
      if (btn.dataset.dialogBound === 'true') return;
      btn.dataset.dialogBound = 'true';
      btn.addEventListener('click', () => {
        const target = btn.getAttribute('data-dialog-target');
        if (!target) return;
        const dialog = document.getElementById(target);
        if (!dialog || typeof dialog.showModal !== 'function') return;

        const value = getFilterValue();
        applyFilterToForms(value);

        if (btn.dataset && btn.dataset.editId) {
          prefillEditForm(btn.dataset, value);
        }

        dialog.showModal();
      });
    });

    document.querySelectorAll('[data-close-dialog]').forEach(btn => {
      if (btn.dataset.dialogBound === 'true') return;
      btn.dataset.dialogBound = 'true';
      btn.addEventListener('click', () => {
        const target = btn.getAttribute('data-close-dialog');
        const dialog = document.getElementById(target);
        if (dialog && typeof dialog.close === 'function') {
          dialog.close();
        }
      });
    });
  };

  const closeDialog = (id) => {
    const dialog = document.getElementById(id);
    if (dialog && typeof dialog.close === 'function') {
      dialog.close();
    }
  };

  window.todoDialogs = {
    closeDialog,
  };

  window.addEventListener('htmx:afterSwap', (event) => {
    if (event.target.id === 'todo-app') {
      requestAnimationFrame(setupDialogControls);
    }
  });

  window.addEventListener('load', setupDialogControls);
})();`

func Layout(title string, content Node) Component {
	assets := NewAssets()
	assets.Collect(content)

	return Html(
		Lang("en"),
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Meta(Name("description"), Content("Modern Todo App built with Plain.")),
			HeadTitle(T(title)),
			Link(LinkRel("preload"), LinkHref("/assets/styles.css"), LinkType("text/css")),
			Link(LinkRel("stylesheet"), LinkHref("/assets/styles.css")),
			assets.CSS(),
		),
		Body(
			Class("bg-background text-foreground min-h-screen"),
			Div(
				Class("min-h-screen bg-background"),
				content,
			),
			Script(ScriptSrc("https://unpkg.com/htmx.org@1.9.12"), Defer()),
			Script(UnsafeText(dialogController)),
			assets.JS(),
		),
	)
}

func AppHeader() Node {
	return Header(
		Class("border-b border-border bg-card/80 backdrop-blur sticky top-0 z-10"),
		Div(
			Class("px-6 py-4 flex items-center justify-between"),
			Div(
				Class("flex items-center gap-3"),
				Div(
					Class("flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10"),
					icons.ListChecks(icons.Size("22"), Class("text-primary")),
				),
				Div(
					H1(Class("text-2xl font-semibold tracking-tight"), T("Tasks")),
					P(Class("text-sm text-muted-foreground"), T("Organize your day, achieve your goals")),
				),
			),
			Button(
				Id("open-add-dialog"),
				Class("inline-flex items-center gap-2 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-primary-foreground shadow transition hover:bg-primary/90"),
				Data("dialog-target", "add-dialog"),
				Span(Class("inline-flex h-5 w-5 items-center justify-center"),
					icons.Plus(icons.Size("16"), Class("text-primary-foreground")),
				),
				T("Add Task"),
			),
		),
	)
}
