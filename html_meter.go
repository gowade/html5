package html5

// HTMLMeter represents HTML <meter> tag
type HTMLMeter struct {
	HTMLElement
}

// Meter creates an HTML <meter> tag element
func Meter() *HTMLMeter {
	e := &HTMLMeter{}
	e.a = make(map[string]interface{})
	e.tagName = "meter"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMeter) S(style StyleMap) *HTMLMeter {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMeter) Key(key interface{}) *HTMLMeter {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLMeter) Ref(dest *DOMElement) *HTMLMeter {
	e.ref = dest
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMeter) ID(v string) *HTMLMeter {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMeter) Class(v string) *HTMLMeter {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMeter) Title(v string) *HTMLMeter {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMeter) Lang(v string) *HTMLMeter {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMeter) Translate(v bool) *HTMLMeter {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMeter) Dir(v string) *HTMLMeter {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMeter) Hidden(v bool) *HTMLMeter {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMeter) TabIndex(v int) *HTMLMeter {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMeter) AccessKey(v string) *HTMLMeter {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMeter) Draggable(v bool) *HTMLMeter {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMeter) Spellcheck(v bool) *HTMLMeter {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
