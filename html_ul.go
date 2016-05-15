package h

// HTMLUL represents HTML <ul> tag
type HTMLUL struct {
	HTMLElement
}

// UL creates a HTML <ul> tag
func UL() *HTMLUL {
	e := &HTMLUL{}
	e.a = make(map[string]interface{})
	e.tagName = "ul"
	return e
}

// S sets the element's CSS properties
func (e *HTMLUL) S(style StyleMap) *HTMLUL {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLUL) Key(key interface{}) *HTMLUL {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLUL) ID(v string) *HTMLUL {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLUL) Class(v string) *HTMLUL {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLUL) Title(v string) *HTMLUL {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLUL) Lang(v string) *HTMLUL {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLUL) Translate(v bool) *HTMLUL {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLUL) Dir(v string) *HTMLUL {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLUL) Hidden(v bool) *HTMLUL {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLUL) TabIndex(v int) *HTMLUL {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLUL) AccessKey(v string) *HTMLUL {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLUL) Draggable(v bool) *HTMLUL {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLUL) Spellcheck(v bool) *HTMLUL {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
