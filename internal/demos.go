package internal

import (
	x "github.com/bloxui/blox"
	"github.com/bloxui/ui"
)

// buildModalDemo creates the modal demonstration page content
func ModalDemo() []x.Component {
	return []x.Component{
		container(
			pageHeader("Modal Components", "CSS-only modal dialogs with smooth animations and keyboard navigation. No JavaScript required."),
			x.Div(
				x.Class("mt-4 space-y-8 mb-18"),
				x.C(ModalExample(
					"Basic Modal",
					"Click the button below to open a basic modal dialog:",
					"Open Basic Modal",
					"basic-modal",
					"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2",
				)),
				x.C(ModalExample(
					"Confirmation Dialog",
					"A confirmation dialog with action buttons:",
					"Delete Account",
					"confirm-modal",
					"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-destructive text-destructive-foreground hover:bg-destructive/90 h-10 px-4 py-2",
				)),
			),
		),
		BasicModal(),
		ConfirmModal(),
	}
}

// buildModalExample creates a modal trigger example section
func ModalExample(title, description, buttonText, target, buttonClass string) x.Component {
	return x.Div(
		x.Class("space-y-4"),
		x.C(x.H2(x.T(title), x.Class("text-2xl font-semibold"))),
		x.C(x.P(x.T(description), x.Class("text-zinc-600 dark:text-zinc-300"))),
		x.C(ui.ModalTrigger(
			ui.ModalTriggerText(buttonText),
			ui.ModalTriggerTarget(target),
			x.Class(buttonClass),
		)),
	)
}

// buildBasicModal creates the basic modal dialog
func BasicModal() x.Component {
	return ui.Modal(
		ui.ModalId("basic-modal"),
		x.C(ui.ModalContent(
			x.C(ui.ModalHeader(
				x.C(ui.ModalTitle(ui.ModalTitleText("Welcome to Blox"))),
				x.C(ui.ModalDescription(ui.ModalDescriptionText("This is a CSS-only modal dialog built with Blox UI components."))),
			)),
			x.C(x.Div(
				x.Class("space-y-4"),
				x.C(x.P(x.T("This modal demonstrates the power of CSS-only interactions using the :target pseudo-class. No JavaScript required!"), x.Class("text-sm"))),
				x.C(FeatureList()),
			)),
			x.C(ui.ModalFooter(
				x.C(CloseButton()),
			)),
		)),
	)
}

// buildConfirmModal creates the confirmation modal dialog
func ConfirmModal() x.Component {
	return ui.Modal(
		ui.ModalId("confirm-modal"),
		x.C(ui.ModalContent(
			x.C(ui.ModalHeader(
				x.C(ui.ModalTitle(ui.ModalTitleText("Confirm Account Deletion"))),
				x.C(ui.ModalDescription(ui.ModalDescriptionText("This action cannot be undone. This will permanently delete your account and remove your data from our servers."))),
			)),
			x.C(ui.ModalFooter(
				x.C(CloseButton()),
				x.C(x.A(
					x.Href("#"),
					x.T("Delete Account"),
					x.Class("inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-destructive text-destructive-foreground hover:bg-destructive/90 h-10 px-4 py-2"),
				)),
			)),
		)),
	)
}

// buildFeatureList creates the modal feature list
func FeatureList() x.Component {
	features := []string{
		"Smooth fade and scale animations",
		"Click outside or press × to close",
		"Keyboard accessible (Tab navigation)",
		"Fully keyboard accessible",
		"No JavaScript dependencies",
	}

	var listItems []x.Component
	for _, feature := range features {
		listItems = append(listItems, x.Li(x.T(feature)))
	}

	var listArgs []x.UlArg
	listArgs = append(listArgs, x.Class("text-sm text-muted-foreground space-y-1 list-disc pl-4"))
	for _, item := range listItems {
		listArgs = append(listArgs, x.C(item))
	}

	return x.Ul(listArgs...)
}

// buildCloseButton creates a modal close button
func CloseButton() x.Component {
	return x.A(
		x.Href("#"),
		x.T("Close"),
		x.Class("inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2"),
	)
}

// buildTabsDemo creates the tabs demonstration page content
func TabsDemo() []x.Component {
	return []x.Component{
		container(
			pageHeader("Tab Components", "Pure CSS tab interface with smooth transitions and keyboard navigation. Built with radio buttons for state management."),
			x.H2(x.T("Basic Tabs"), x.Class("mt-12 mb-4 text-2xl font-semibold")),
			x.P(x.T("A simple tab interface with three panels:"), x.Class("mb-6 text-zinc-600 dark:text-zinc-300")),
			MainTabsDemo(),
			x.H2(x.T("Multiple Tab Groups"), x.Class("mt-12 mb-4 text-2xl font-semibold")),
			x.P(x.T("You can have multiple independent tab groups on the same page:"), x.Class("mb-6 text-zinc-600 dark:text-zinc-300")),
			MultipleTabsDemo(),
			TabsDocumentation(),
		),
	}
}

// buildMainTabsDemo creates the main tabs demonstration
func MainTabsDemo() x.Component {
	return ui.Tabs(
		ui.TabsName("demo-tabs"),
		x.Class("w-full"),
		x.C(ui.TabsList(
			x.Class("w-fit"),
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("account"),
				ui.TabsTriggerGroup("demo-tabs"),
				ui.TabsTriggerLabel("Account"),
				ui.TabsTriggerDefault(),
			)),
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("password"),
				ui.TabsTriggerGroup("demo-tabs"),
				ui.TabsTriggerLabel("Password"),
			)),
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("settings"),
				ui.TabsTriggerGroup("demo-tabs"),
				ui.TabsTriggerLabel("Settings"),
			)),
		)),
		x.C(AccountTabContent()),
		x.C(PasswordTabContent()),
		x.C(SettingsTabContent()),
	)
}

// buildAccountTabContent creates the account tab panel content
func AccountTabContent() x.Component {
	return ui.TabsContent(
		ui.TabsContentValue("account"),
		ui.TabsContentGroup("demo-tabs"),
		x.Class("space-y-2"),
		x.C(x.H3(x.T("Account Settings"), x.Class("text-lg font-semibold"))),
		x.C(x.P(x.T("Manage your account information and preferences."), x.Class("text-zinc-600 dark:text-zinc-300"))),
		x.C(ui.Card(
			x.C(ui.CardHeader(
				x.C(ui.CardTitle(ui.CardTitleText("Profile Information"))),
				x.C(ui.CardDescription(ui.CardDescriptionText("Update your profile details"))),
			)),
			x.C(ui.CardContent(
				x.C(x.Div(
					x.Class("space-y-4"),
					x.C(FormField("Username", "username", "text", "johndoe", false)),
					x.C(FormField("Email", "email", "email", "john@example.com", false)),
				)),
			)),
		)),
	)
}

// buildPasswordTabContent creates the password tab panel content
func PasswordTabContent() x.Component {
	return ui.TabsContent(
		ui.TabsContentValue("password"),
		ui.TabsContentGroup("demo-tabs"),
		x.Class("space-y-2"),
		x.C(x.H3(x.T("Password Settings"), x.Class("text-lg font-semibold"))),
		x.C(x.P(x.T("Change your password and security settings."), x.Class("text-zinc-600 dark:text-zinc-300"))),
		x.C(ui.Card(
			x.C(ui.CardHeader(
				x.C(ui.CardTitle(ui.CardTitleText("Change Password"))),
				x.C(ui.CardDescription(ui.CardDescriptionText("Ensure your account is secure"))),
			)),
			x.C(ui.CardContent(
				x.C(x.Div(
					x.Class("space-y-4"),
					x.C(FormField("Current Password", "current", "password", "Enter current password", false)),
					x.C(FormField("New Password", "new", "password", "Enter new password", false)),
					x.C(FormField("Confirm Password", "confirm", "password", "Confirm new password", false)),
				)),
			)),
			x.C(ui.CardFooter(
				x.C(ui.Button(ui.Text("Update Password"))),
			)),
		)),
	)
}

// buildSettingsTabContent creates the settings tab panel content
func SettingsTabContent() x.Component {
	return ui.TabsContent(
		ui.TabsContentValue("settings"),
		ui.TabsContentGroup("demo-tabs"),
		x.Class("space-y-2"),
		x.C(x.H3(x.T("Application Settings"), x.Class("text-lg font-semibold"))),
		x.C(x.P(x.T("Configure your application preferences."), x.Class("text-zinc-600 dark:text-zinc-300"))),
		x.C(ui.Card(
			x.C(ui.CardHeader(
				x.C(ui.CardTitle(ui.CardTitleText("Preferences"))),
				x.C(ui.CardDescription(ui.CardDescriptionText("Customize your experience"))),
			)),
			x.C(ui.CardContent(
				x.C(x.Div(
					x.Class("space-y-4"),
					x.C(ui.Checkbox(ui.CheckboxLabel("Enable notifications"), x.Id("notifications"))),
					x.C(ui.Checkbox(ui.CheckboxLabel("Auto-save drafts"), x.Id("autosave"))),
					x.C(ui.Checkbox(ui.CheckboxLabel("Show advanced options"), x.Id("advanced"))),
				)),
			)),
		)),
	)
}

// buildMultipleTabsDemo creates the multiple tabs demonstration
func MultipleTabsDemo() x.Component {
	return x.Div(
		x.Class("grid md:grid-cols-2 gap-8"),
		x.C(FirstTabGroup()),
		x.C(SecondTabGroup()),
	)
}

// buildFirstTabGroup creates the first tab group
func FirstTabGroup() x.Component {
	return ui.Tabs(
		ui.TabsName("tabs-1"),
		x.Class("w-full"),
		x.C(ui.TabsList(
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("overview"),
				ui.TabsTriggerGroup("tabs-1"),
				ui.TabsTriggerLabel("Overview"),
				ui.TabsTriggerDefault(),
			)),
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("analytics"),
				ui.TabsTriggerGroup("tabs-1"),
				ui.TabsTriggerLabel("Analytics"),
			)),
		)),
		x.C(ui.TabsContent(
			ui.TabsContentValue("overview"),
			ui.TabsContentGroup("tabs-1"),
			x.C(x.P(x.T("Welcome to your dashboard overview. Here you can see a summary of your key metrics."), x.Class("text-sm text-zinc-600 dark:text-zinc-300"))),
		)),
		x.C(ui.TabsContent(
			ui.TabsContentValue("analytics"),
			ui.TabsContentGroup("tabs-1"),
			x.C(x.P(x.T("Detailed analytics and insights about your usage patterns and performance metrics."), x.Class("text-sm text-zinc-600 dark:text-zinc-300"))),
		)),
	)
}

// buildSecondTabGroup creates the second tab group
func SecondTabGroup() x.Component {
	return ui.Tabs(
		ui.TabsName("tabs-2"),
		x.Class("w-full"),
		x.C(ui.TabsList(
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("general"),
				ui.TabsTriggerGroup("tabs-2"),
				ui.TabsTriggerLabel("General"),
				ui.TabsTriggerDefault(),
			)),
			x.C(ui.TabsTrigger(
				ui.TabsTriggerValue("advanced"),
				ui.TabsTriggerGroup("tabs-2"),
				ui.TabsTriggerLabel("Advanced"),
			)),
		)),
		x.C(ui.TabsContent(
			ui.TabsContentValue("general"),
			ui.TabsContentGroup("tabs-2"),
			x.C(x.P(x.T("General configuration options and basic settings for your application."), x.Class("text-sm text-zinc-600 dark:text-zinc-300"))),
		)),
		x.C(ui.TabsContent(
			ui.TabsContentValue("advanced"),
			ui.TabsContentGroup("tabs-2"),
			x.C(x.P(x.T("Advanced configuration options for power users and developers."), x.Class("text-sm text-zinc-600 dark:text-zinc-300"))),
		)),
	)
}

// buildTabsDocumentation creates the tabs documentation section
func TabsDocumentation() x.Component {
	benefits := []string{
		"No JavaScript required",
		"Fully keyboard accessible",
		"Works with screen readers",
		"Lightweight and fast",
		"Easy to style with CSS",
	}

	var benefitItems []x.Component
	for _, benefit := range benefits {
		benefitItems = append(benefitItems, x.Li(x.T(benefit)))
	}

	var listArgs []x.UlArg
	listArgs = append(listArgs, x.Class("list-disc list-inside space-y-1 text-zinc-600 dark:text-zinc-300 mb-4"))
	for _, item := range benefitItems {
		listArgs = append(listArgs, x.C(item))
	}

	return x.Div(
		x.C(x.H2(x.T("How It Works"), x.Class("mt-12 mb-4 text-2xl font-semibold"))),
		x.C(x.P(x.T("These tabs are implemented using pure CSS with hidden radio buttons for state management. Each tab trigger is a label that controls a hidden radio input, and CSS selectors show/hide content based on which radio is checked."), x.Class("mb-4 text-zinc-600 dark:text-zinc-300"))),
		x.C(x.P(x.T("Benefits of this approach:"), x.Class("mb-2 font-medium"))),
		x.C(x.Ul(listArgs...)),
	)
}
