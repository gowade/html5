package h

// HTMLLI represents HTML <li> tag
type HTMLLI struct {
	HTMLElement
}

// LI creates an HTML <li> tag element
func LI() *HTMLLI {
	e := &HTMLLI{}
	e.a = make(map[string]interface{})
	e.tagName = "li"
	return e
}

// S sets the element's CSS properties
func (e *HTMLLI) S(style StyleMap) *HTMLLI {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLLI) Key(key interface{}) *HTMLLI {
	e.key = F(key)
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLLI) Value(v int) *HTMLLI {
	e.a["value"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLLI) ID(v string) *HTMLLI {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLLI) Class(v string) *HTMLLI {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLLI) Title(v string) *HTMLLI {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLLI) Lang(v string) *HTMLLI {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLLI) Translate(v bool) *HTMLLI {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLLI) Dir(v string) *HTMLLI {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLLI) Hidden(v bool) *HTMLLI {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLLI) TabIndex(v int) *HTMLLI {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLLI) AccessKey(v string) *HTMLLI {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLLI) Draggable(v bool) *HTMLLI {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLLI) Spellcheck(v bool) *HTMLLI {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
