package templates

type TemplateName string

const (
	TemplateHomePage       TemplateName = "home.page.html"
	TemplateLoginPage      TemplateName = "login.page.html"
	TemplateHeadingSection TemplateName = "heading_section.partial"
)

// ---------------------------------------------------------------------------
//
// Page args.
//
// ---------------------------------------------------------------------------
type HomePageArgs struct {
	Headings []HeadingSection
}

// ---------------------------------------------------------------------------
//
// Components / partials args.
//
// ---------------------------------------------------------------------------
type HeadingSection struct {
	Level       uint8
	Heading     string
	Description string
}
