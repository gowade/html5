package h

// HTMLParam represents HTML <param> tag
type HTMLParam struct {
	HTMLElement
}

// Param creates an HTML <param> tag element
func Param() *HTMLParam {
	e := &HTMLParam{}
	e.a = make(map[string]interface{})
	e.tagName = "param"
	return e
}

// S sets the element's CSS properties
func (e *HTMLParam) S(style StyleMap) *HTMLParam {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLParam) Key(key interface{}) *HTMLParam {
	e.key = F(key)
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLParam) Name(v string) *HTMLParam {
	e.a["name"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLParam) Value(v string) *HTMLParam {
	e.a["value"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLParam) ID(v string) *HTMLParam {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLParam) Class(v string) *HTMLParam {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLParam) Title(v string) *HTMLParam {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLParam) Lang(v string) *HTMLParam {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLParam) Translate(v bool) *HTMLParam {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLParam) Dir(v string) *HTMLParam {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLParam) Hidden(v bool) *HTMLParam {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLParam) TabIndex(v int) *HTMLParam {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLParam) AccessKey(v string) *HTMLParam {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLParam) Draggable(v bool) *HTMLParam {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLParam) Spellcheck(v bool) *HTMLParam {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
