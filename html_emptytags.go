package html5

// Html creates an HTML <html> tag element
func Html() *HTMLElement {
	return Element("html")
}

// Head creates an HTML <head> tag element
func Head() *HTMLElement {
	return Element("head")
}

// Body creates an HTML <body> tag element
func Body() *HTMLElement {
	return Element("body")
}

// Heading creates an HTML <heading> tag element
func Heading() *HTMLElement {
	return Element("heading")
}

// Paragraph creates an HTML <paragraph> tag element
func Paragraph() *HTMLElement {
	return Element("paragraph")
}

// HR creates an HTML <hr> tag element
func HR() *HTMLElement {
	return Element("hr")
}

// Pre creates an HTML <pre> tag element
func Pre() *HTMLElement {
	return Element("pre")
}

// UL creates an HTML <ul> tag element
func UL() *HTMLElement {
	return Element("ul")
}

// DList creates an HTML <dlist> tag element
func DList() *HTMLElement {
	return Element("dlist")
}

// Div creates an HTML <div> tag element
func Div() *HTMLElement {
	return Element("div")
}

// Span creates an HTML <span> tag element
func Span() *HTMLElement {
	return Element("span")
}

// BR creates an HTML <br> tag element
func BR() *HTMLElement {
	return Element("br")
}

// Picture creates an HTML <picture> tag element
func Picture() *HTMLElement {
	return Element("picture")
}

// TableCaption creates an HTML <tablecaption> tag element
func TableCaption() *HTMLElement {
	return Element("tablecaption")
}
