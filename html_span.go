package h

// HTMLSpan represents HTML <span> tag
type HTMLSpan struct {
	HTMLElement
}

// Span creates a HTML <span> tag
func Span() *HTMLSpan {
	e := &HTMLSpan{}
	e.a = make(map[string]interface{})
	e.tagName = "span"
	return e
}

// S sets the element's CSS properties
func (e *HTMLSpan) S(style StyleMap) *HTMLSpan {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLSpan) Key(key interface{}) *HTMLSpan {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLSpan) ID(v string) *HTMLSpan {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLSpan) Class(v string) *HTMLSpan {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLSpan) Title(v string) *HTMLSpan {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLSpan) Lang(v string) *HTMLSpan {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLSpan) Translate(v bool) *HTMLSpan {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLSpan) Dir(v string) *HTMLSpan {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLSpan) Hidden(v bool) *HTMLSpan {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLSpan) TabIndex(v int) *HTMLSpan {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLSpan) AccessKey(v string) *HTMLSpan {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLSpan) Draggable(v bool) *HTMLSpan {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLSpan) Spellcheck(v bool) *HTMLSpan {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
