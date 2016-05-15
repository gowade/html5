package h

// HTMLPre represents HTML <pre> tag
type HTMLPre struct {
	HTMLElement
}

// Pre creates a HTML <pre> tag
func Pre() *HTMLPre {
	e := &HTMLPre{}
	e.a = make(map[string]interface{})
	e.tagName = "pre"
	return e
}

// S sets the element's CSS properties
func (e *HTMLPre) S(style StyleMap) *HTMLPre {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLPre) Key(key interface{}) *HTMLPre {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLPre) ID(v string) *HTMLPre {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLPre) Class(v string) *HTMLPre {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLPre) Title(v string) *HTMLPre {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLPre) Lang(v string) *HTMLPre {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLPre) Translate(v bool) *HTMLPre {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLPre) Dir(v string) *HTMLPre {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLPre) Hidden(v bool) *HTMLPre {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLPre) TabIndex(v int) *HTMLPre {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLPre) AccessKey(v string) *HTMLPre {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLPre) Draggable(v bool) *HTMLPre {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLPre) Spellcheck(v bool) *HTMLPre {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
