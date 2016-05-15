package html5

// HTMLData represents HTML <data> tag
type HTMLData struct {
	HTMLElement
}

// Data creates an HTML <data> tag element
func Data() *HTMLData {
	e := &HTMLData{}
	e.a = make(map[string]interface{})
	e.tagName = "data"
	return e
}

// S sets the element's CSS properties
func (e *HTMLData) S(style StyleMap) *HTMLData {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLData) Key(key interface{}) *HTMLData {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLData) Ref(dest *DOMElement) *HTMLData {
	e.ref = dest
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLData) Value(v string) *HTMLData {
	e.a["value"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLData) ID(v string) *HTMLData {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLData) Class(v string) *HTMLData {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLData) Title(v string) *HTMLData {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLData) Lang(v string) *HTMLData {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLData) Translate(v bool) *HTMLData {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLData) Dir(v string) *HTMLData {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLData) Hidden(v bool) *HTMLData {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLData) TabIndex(v int) *HTMLData {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLData) AccessKey(v string) *HTMLData {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLData) Draggable(v bool) *HTMLData {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLData) Spellcheck(v bool) *HTMLData {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
