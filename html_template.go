package html5

// HTMLTemplate represents HTML <template> tag
type HTMLTemplate struct {
	HTMLElement
}

// Template creates an HTML <template> tag element
func Template() *HTMLTemplate {
	e := &HTMLTemplate{}
	e.a = make(map[string]interface{})
	e.tagName = "template"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTemplate) S(style StyleMap) *HTMLTemplate {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTemplate) Key(key interface{}) *HTMLTemplate {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLTemplate) Ref(dest *DOMElement) *HTMLTemplate {
	e.ref = dest
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTemplate) ID(v string) *HTMLTemplate {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTemplate) Class(v string) *HTMLTemplate {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTemplate) Title(v string) *HTMLTemplate {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTemplate) Lang(v string) *HTMLTemplate {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTemplate) Translate(v bool) *HTMLTemplate {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTemplate) Dir(v string) *HTMLTemplate {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTemplate) Hidden(v bool) *HTMLTemplate {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTemplate) TabIndex(v int) *HTMLTemplate {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTemplate) AccessKey(v string) *HTMLTemplate {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTemplate) Draggable(v bool) *HTMLTemplate {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTemplate) Spellcheck(v bool) *HTMLTemplate {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
