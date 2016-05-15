package h

// HTMLHtml represents HTML <html> tag
type HTMLHtml struct{ HTMLElement }

// Html creates an HTML <html> tag element
func Html() *HTMLHtml {
	return Element("html")
}

// HTMLHead represents HTML <head> tag
type HTMLHead struct{ HTMLElement }

// Head creates an HTML <head> tag element
func Head() *HTMLHead {
	return Element("head")
}

// HTMLBody represents HTML <body> tag
type HTMLBody struct{ HTMLElement }

// Body creates an HTML <body> tag element
func Body() *HTMLBody {
	return Element("body")
}

// HTMLHeading represents HTML <heading> tag
type HTMLHeading struct{ HTMLElement }

// Heading creates an HTML <heading> tag element
func Heading() *HTMLHeading {
	return Element("heading")
}

// HTMLParagraph represents HTML <paragraph> tag
type HTMLParagraph struct{ HTMLElement }

// Paragraph creates an HTML <paragraph> tag element
func Paragraph() *HTMLParagraph {
	return Element("paragraph")
}

// HTMLHR represents HTML <hr> tag
type HTMLHR struct{ HTMLElement }

// HR creates an HTML <hr> tag element
func HR() *HTMLHR {
	return Element("hr")
}

// HTMLPre represents HTML <pre> tag
type HTMLPre struct{ HTMLElement }

// Pre creates an HTML <pre> tag element
func Pre() *HTMLPre {
	return Element("pre")
}

// HTMLUL represents HTML <ul> tag
type HTMLUL struct{ HTMLElement }

// UL creates an HTML <ul> tag element
func UL() *HTMLUL {
	return Element("ul")
}

// HTMLDList represents HTML <dlist> tag
type HTMLDList struct{ HTMLElement }

// DList creates an HTML <dlist> tag element
func DList() *HTMLDList {
	return Element("dlist")
}

// HTMLDiv represents HTML <div> tag
type HTMLDiv struct{ HTMLElement }

// Div creates an HTML <div> tag element
func Div() *HTMLDiv {
	return Element("div")
}

// HTMLSpan represents HTML <span> tag
type HTMLSpan struct{ HTMLElement }

// Span creates an HTML <span> tag element
func Span() *HTMLSpan {
	return Element("span")
}

// HTMLBR represents HTML <br> tag
type HTMLBR struct{ HTMLElement }

// BR creates an HTML <br> tag element
func BR() *HTMLBR {
	return Element("br")
}

// HTMLPicture represents HTML <picture> tag
type HTMLPicture struct{ HTMLElement }

// Picture creates an HTML <picture> tag element
func Picture() *HTMLPicture {
	return Element("picture")
}

// HTMLTableCaption represents HTML <tablecaption> tag
type HTMLTableCaption struct{ HTMLElement }

// TableCaption creates an HTML <tablecaption> tag element
func TableCaption() *HTMLTableCaption {
	return Element("tablecaption")
}
