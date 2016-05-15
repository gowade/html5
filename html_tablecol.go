package html5

// HTMLTableCol represents HTML <tablecol> tag
type HTMLTableCol struct {
	HTMLElement
}

// TableCol creates an HTML <tablecol> tag element
func TableCol() *HTMLTableCol {
	e := &HTMLTableCol{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecol"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTableCol) S(style StyleMap) *HTMLTableCol {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTableCol) Key(key interface{}) *HTMLTableCol {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLTableCol) Ref(dest *DOMElement) *HTMLTableCol {
	e.ref = dest
	return e
}

// Span sets the element's "span" attribute
func (e *HTMLTableCol) Span(v int) *HTMLTableCol {
	e.a["span"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTableCol) ID(v string) *HTMLTableCol {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTableCol) Class(v string) *HTMLTableCol {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTableCol) Title(v string) *HTMLTableCol {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTableCol) Lang(v string) *HTMLTableCol {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTableCol) Translate(v bool) *HTMLTableCol {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTableCol) Dir(v string) *HTMLTableCol {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTableCol) Hidden(v bool) *HTMLTableCol {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTableCol) TabIndex(v int) *HTMLTableCol {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTableCol) AccessKey(v string) *HTMLTableCol {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTableCol) Draggable(v bool) *HTMLTableCol {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTableCol) Spellcheck(v bool) *HTMLTableCol {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
