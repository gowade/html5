package h

// HTMLHR represents HTML <hr> tag
type HTMLHR struct {
	HTMLElement
}

// HR creates a HTML <hr> tag
func HR() *HTMLHR {
	e := &HTMLHR{}
	e.a = make(map[string]interface{})
	e.tagName = "hr"
	return e
}

// S sets the element's CSS properties
func (e *HTMLHR) S(style StyleMap) *HTMLHR {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLHR) Key(key interface{}) *HTMLHR {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLHR) ID(v string) *HTMLHR {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLHR) Class(v string) *HTMLHR {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLHR) Title(v string) *HTMLHR {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLHR) Lang(v string) *HTMLHR {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLHR) Translate(v bool) *HTMLHR {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLHR) Dir(v string) *HTMLHR {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLHR) Hidden(v bool) *HTMLHR {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLHR) TabIndex(v int) *HTMLHR {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLHR) AccessKey(v string) *HTMLHR {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLHR) Draggable(v bool) *HTMLHR {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLHR) Spellcheck(v bool) *HTMLHR {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
