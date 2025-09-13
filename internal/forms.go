package internal

import (
	x "github.com/bloxui/blox"
	"github.com/bloxui/ui"
)

// ContactForm creates the comprehensive contact form
func ContactForm() x.Component {
	return x.Div(
		x.Class("mb-12"),
		x.C(pageHeader("UI Components Demo", "A comprehensive showcase of Blox UI components in a real-world contact form.")),
		x.C(
			x.Form(
				x.Action("/contact"),
				x.Method("POST"),
				x.Class("mt-6 max-w-2xl space-y-8 bg-card p-8 rounded-lg border shadow"),

				// Personal Information Section
				x.C(PersonalInfoSection()),

				// Company Information Section
				x.C(CompanyInfoSection()),

				// Inquiry Details Section
				x.C(InquirySection()),

				// Preferences Section
				x.C(PreferencesSection()),

				// Terms and Submit Section
				x.C(SubmitSection()),
			),
		),
	)
}

// buildPersonalInfoSection creates the personal information form section
func PersonalInfoSection() x.Component {
	return x.Div(
		x.Class("space-y-4"),
		x.C(x.H3(x.T("Personal Information"), x.Class("text-lg font-semibold"))),
		x.C(
			x.Div(
				x.Class("grid grid-cols-1 sm:grid-cols-2 gap-4"),
				x.C(FormField("First Name", "first_name", "text", "John", true)),
				x.C(FormField("Last Name", "last_name", "text", "Doe", true)),
			),
		),
		x.C(FormField("Email Address", "email", "email", "john@example.com", true)),
		x.C(FormField("Phone Number (Optional)", "phone", "tel", "+1 (555) 123-4567", false)),
	)
}

// buildCompanyInfoSection creates the company information form section
func CompanyInfoSection() x.Component {
	return x.Div(
		x.Class("space-y-4"),
		x.C(x.H3(x.T("Company Information"), x.Class("text-lg font-semibold"))),
		x.C(FormField("Company Name (Optional)", "company", "text", "Acme Corporation", false)),
		x.C(
			x.Div(
				x.Class("space-y-3"),
				x.C(ui.Label(ui.LabelText("Company Size"))),
				x.C(ui.RadioGroup(
					x.C(ui.Radio(ui.RadioLabel("Solo developer"), x.InputName("company_size"), x.InputValue("solo"))),
					x.C(ui.Radio(ui.RadioLabel("2-10 employees"), x.InputName("company_size"), x.InputValue("small"))),
					x.C(ui.Radio(ui.RadioLabel("11-50 employees"), x.InputName("company_size"), x.InputValue("medium"))),
					x.C(ui.Radio(ui.RadioLabel("50+ employees"), x.InputName("company_size"), x.InputValue("large"))),
				)),
			),
		),
	)
}

// buildInquirySection creates the inquiry details form section
func InquirySection() x.Component {
	return x.Div(
		x.Class("space-y-4"),
		x.C(x.H3(x.T("Inquiry Details"), x.Class("text-lg font-semibold"))),
		x.C(
			x.Div(
				x.Class("space-y-3"),
				x.C(ui.Label(ui.LabelText("What brings you here?"))),
				x.C(ui.RadioGroup(
					x.C(ui.Radio(ui.RadioLabel("General question about Blox"), x.InputName("inquiry_type"), x.InputValue("general"))),
					x.C(ui.Radio(ui.RadioLabel("Technical support needed"), x.InputName("inquiry_type"), x.InputValue("support"))),
					x.C(ui.Radio(ui.RadioLabel("Partnership opportunity"), x.InputName("inquiry_type"), x.InputValue("partnership"))),
					x.C(ui.Radio(ui.RadioLabel("Feature request or feedback"), x.InputName("inquiry_type"), x.InputValue("feature"))),
				)),
			),
		),
		x.C(
			x.Div(
				x.C(ui.Label(ui.LabelText("Message"), x.For("message"))),
				x.C(ui.Textarea(x.Id("message"), x.TextareaName("message"), x.Rows(4), x.Required(), x.Placeholder("Please describe how we can help you or what you're looking to accomplish with Blox..."))),
			),
		),
	)
}

// buildPreferencesSection creates the communication preferences form section
func PreferencesSection() x.Component {
	return x.Div(
		x.Class("space-y-4"),
		x.C(x.H3(x.T("Communication Preferences"), x.Class("text-lg font-semibold"))),
		x.C(
			x.Div(
				x.Class("space-y-3"),
				x.C(ui.Label(ui.LabelText("Preferred contact method"))),
				x.C(ui.RadioGroup(
					x.C(ui.Radio(ui.RadioLabel("Email"), x.InputName("contact_method"), x.InputValue("email"))),
					x.C(ui.Radio(ui.RadioLabel("Phone call"), x.InputName("contact_method"), x.InputValue("phone"))),
					x.C(ui.Radio(ui.RadioLabel("Video call"), x.InputName("contact_method"), x.InputValue("video"))),
				)),
			),
		),
		x.C(
			x.Div(
				x.Class("space-y-3"),
				x.C(ui.Checkbox(ui.CheckboxLabel("Send me Blox newsletter and product updates"), x.Id("newsletter"), x.InputName("newsletter"))),
				x.C(ui.Checkbox(ui.CheckboxLabel("Notify me about new releases and features"), x.Id("releases"), x.InputName("releases"))),
				x.C(ui.Checkbox(ui.CheckboxLabel("Include me in beta testing opportunities"), x.Id("beta"), x.InputName("beta"))),
			),
		),
	)
}

// buildSubmitSection creates the terms and submit form section
func SubmitSection() x.Component {
	return x.Div(
		x.Class("space-y-4 pt-6 border-t"),
		x.C(ui.Checkbox(ui.CheckboxLabel("I agree to the privacy policy and terms of service"), x.Id("privacy"), x.InputName("privacy"), x.Required())),
		x.C(
			x.Div(
				x.Class("flex gap-3"),
				x.C(ui.Button(ui.Default(), x.ButtonType("submit"), x.T("Send Message"), x.Class("flex-1"))),
				x.C(ui.Button(ui.Outline(), x.ButtonType("button"), x.T("Reset Form"))),
			),
		),
	)
}

// buildFormField creates a standard form field with label and input
func FormField(label, id, inputType, placeholder string, required bool) x.Component {
	var inputArgs []interface{}
	inputArgs = append(inputArgs, x.InputType(inputType), x.Id(id), x.InputName(id), x.Placeholder(placeholder))
	if required {
		inputArgs = append(inputArgs, x.Required())
	}

	return x.Div(
		x.C(ui.Label(ui.LabelText(label), x.For(id))),
		x.C(ui.Input(inputArgs...)),
	)
}
