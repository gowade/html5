package h

// HTMLBR represents HTML <br> tag
type HTMLBR struct {
	HTMLElement
}

// BR creates a HTML <br> tag
func BR() *HTMLBR {
	e := &HTMLBR{}
	e.a = make(map[string]interface{})
	e.tagName = "br"
	return e
}

// S sets the element's CSS properties
func (e *HTMLBR) S(style StyleMap) *HTMLBR {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLBR) Key(key interface{}) *HTMLBR {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLBR) ID(v string) *HTMLBR {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLBR) Class(v string) *HTMLBR {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLBR) Title(v string) *HTMLBR {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLBR) Lang(v string) *HTMLBR {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLBR) Translate(v bool) *HTMLBR {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLBR) Dir(v string) *HTMLBR {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLBR) Hidden(v bool) *HTMLBR {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLBR) TabIndex(v int) *HTMLBR {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLBR) AccessKey(v string) *HTMLBR {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLBR) Draggable(v bool) *HTMLBR {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLBR) Spellcheck(v bool) *HTMLBR {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
