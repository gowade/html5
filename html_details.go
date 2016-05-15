package html5

// HTMLDetails represents HTML <details> tag
type HTMLDetails struct {
	HTMLElement
}

// Details creates an HTML <details> tag element
func Details() *HTMLDetails {
	e := &HTMLDetails{}
	e.a = make(map[string]interface{})
	e.tagName = "details"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDetails) S(style StyleMap) *HTMLDetails {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDetails) Key(key interface{}) *HTMLDetails {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLDetails) Ref(dest *DOMElement) *HTMLDetails {
	e.ref = dest
	return e
}

// Open sets the element's "open" attribute
func (e *HTMLDetails) Open(v bool) *HTMLDetails {
	if v {
		e.a["open"] = ""
	} else {
		delete(e.a, "open")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDetails) ID(v string) *HTMLDetails {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDetails) Class(v string) *HTMLDetails {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDetails) Title(v string) *HTMLDetails {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDetails) Lang(v string) *HTMLDetails {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDetails) Translate(v bool) *HTMLDetails {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDetails) Dir(v string) *HTMLDetails {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDetails) Hidden(v bool) *HTMLDetails {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDetails) TabIndex(v int) *HTMLDetails {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDetails) AccessKey(v string) *HTMLDetails {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDetails) Draggable(v bool) *HTMLDetails {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDetails) Spellcheck(v bool) *HTMLDetails {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
