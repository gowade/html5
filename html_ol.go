package html5

// HTMLOL represents HTML <ol> tag
type HTMLOL struct {
	HTMLElement
}

// OL creates an HTML <ol> tag element
func OL() *HTMLOL {
	e := &HTMLOL{}
	e.a = make(map[string]interface{})
	e.tagName = "ol"
	return e
}

// S sets the element's CSS properties
func (e *HTMLOL) S(style StyleMap) *HTMLOL {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLOL) Key(key interface{}) *HTMLOL {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLOL) Ref(dest *DOMElement) *HTMLOL {
	e.ref = dest
	return e
}

// Reversed sets the element's "reversed" attribute
func (e *HTMLOL) Reversed(v bool) *HTMLOL {
	if v {
		e.a["reversed"] = ""
	} else {
		delete(e.a, "reversed")
	}
	return e
}

// Start sets the element's "start" attribute
func (e *HTMLOL) Start(v int) *HTMLOL {
	e.a["start"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLOL) Type(v string) *HTMLOL {
	e.a["type"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLOL) ID(v string) *HTMLOL {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLOL) Class(v string) *HTMLOL {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLOL) Title(v string) *HTMLOL {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLOL) Lang(v string) *HTMLOL {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLOL) Translate(v bool) *HTMLOL {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLOL) Dir(v string) *HTMLOL {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLOL) Hidden(v bool) *HTMLOL {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLOL) TabIndex(v int) *HTMLOL {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLOL) AccessKey(v string) *HTMLOL {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLOL) Draggable(v bool) *HTMLOL {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLOL) Spellcheck(v bool) *HTMLOL {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
