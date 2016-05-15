package html5

// HTMLOptGroup represents HTML <optgroup> tag
type HTMLOptGroup struct {
	HTMLElement
}

// OptGroup creates an HTML <optgroup> tag element
func OptGroup() *HTMLOptGroup {
	e := &HTMLOptGroup{}
	e.a = make(map[string]interface{})
	e.tagName = "optgroup"
	return e
}

// S sets the element's CSS properties
func (e *HTMLOptGroup) S(style StyleMap) *HTMLOptGroup {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLOptGroup) Key(key interface{}) *HTMLOptGroup {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLOptGroup) Ref(dest *DOMElement) *HTMLOptGroup {
	e.ref = dest
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLOptGroup) Disabled(v bool) *HTMLOptGroup {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Label sets the element's "label" attribute
func (e *HTMLOptGroup) Label(v string) *HTMLOptGroup {
	e.a["label"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLOptGroup) ID(v string) *HTMLOptGroup {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLOptGroup) Class(v string) *HTMLOptGroup {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLOptGroup) Title(v string) *HTMLOptGroup {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLOptGroup) Lang(v string) *HTMLOptGroup {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLOptGroup) Translate(v bool) *HTMLOptGroup {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLOptGroup) Dir(v string) *HTMLOptGroup {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLOptGroup) Hidden(v bool) *HTMLOptGroup {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLOptGroup) TabIndex(v int) *HTMLOptGroup {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLOptGroup) AccessKey(v string) *HTMLOptGroup {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLOptGroup) Draggable(v bool) *HTMLOptGroup {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLOptGroup) Spellcheck(v bool) *HTMLOptGroup {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
