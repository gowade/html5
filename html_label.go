package html5

// HTMLLabel represents HTML <label> tag
type HTMLLabel struct {
	HTMLElement
}

// Label creates an HTML <label> tag element
func Label() *HTMLLabel {
	e := &HTMLLabel{}
	e.a = make(map[string]interface{})
	e.tagName = "label"
	return e
}

// S sets the element's CSS properties
func (e *HTMLLabel) S(style StyleMap) *HTMLLabel {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLLabel) Key(key interface{}) *HTMLLabel {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLLabel) Ref(dest *DOMElement) *HTMLLabel {
	e.ref = dest
	return e
}

// For sets the element's "htmlfor" attribute
func (e *HTMLLabel) For(v string) *HTMLLabel {
	e.a["htmlfor"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLLabel) ID(v string) *HTMLLabel {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLLabel) Class(v string) *HTMLLabel {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLLabel) Title(v string) *HTMLLabel {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLLabel) Lang(v string) *HTMLLabel {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLLabel) Translate(v bool) *HTMLLabel {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLLabel) Dir(v string) *HTMLLabel {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLLabel) Hidden(v bool) *HTMLLabel {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLLabel) TabIndex(v int) *HTMLLabel {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLLabel) AccessKey(v string) *HTMLLabel {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLLabel) Draggable(v bool) *HTMLLabel {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLLabel) Spellcheck(v bool) *HTMLLabel {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
