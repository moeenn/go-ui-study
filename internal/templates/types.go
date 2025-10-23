package templates

type TemplateName string

const (
	TemplateHomePage           TemplateName = "home.page.html"
	TemplateLoginPage          TemplateName = "login.page.html"
	TemplateForgotPasswordPage TemplateName = "forgot_password.page.html"
	TemplateHeadingSection     TemplateName = "heading_section.partial"
	TemplateErrorPage          TemplateName = "error.page.html"
)

// ---------------------------------------------------------------------------
//
// Page args.
//
// ---------------------------------------------------------------------------
type ErrorPageArgs struct {
	StatusCode int
	Error      string
}

// ---------------------------------------------------------------------------
//
// Components / partials args.
//
// ---------------------------------------------------------------------------
