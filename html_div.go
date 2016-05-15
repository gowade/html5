package h

// HTMLDiv represents HTML <div> tag
type HTMLDiv struct {
	HTMLElement
}

// Div creates a HTML <div> tag
func Div() *HTMLDiv {
	e := &HTMLDiv{}
	e.a = make(map[string]interface{})
	e.tagName = "div"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDiv) S(style StyleMap) *HTMLDiv {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDiv) Key(key interface{}) *HTMLDiv {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDiv) ID(v string) *HTMLDiv {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDiv) Class(v string) *HTMLDiv {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDiv) Title(v string) *HTMLDiv {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDiv) Lang(v string) *HTMLDiv {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDiv) Translate(v bool) *HTMLDiv {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDiv) Dir(v string) *HTMLDiv {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDiv) Hidden(v bool) *HTMLDiv {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDiv) TabIndex(v int) *HTMLDiv {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDiv) AccessKey(v string) *HTMLDiv {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDiv) Draggable(v bool) *HTMLDiv {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDiv) Spellcheck(v bool) *HTMLDiv {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
