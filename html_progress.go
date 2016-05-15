package h

// HTMLProgress represents HTML <progress> tag
type HTMLProgress struct {
	HTMLElement
}

// Progress creates a HTML <progress> tag
func Progress() *HTMLProgress {
	e := &HTMLProgress{}
	e.a = make(map[string]interface{})
	e.tagName = "progress"
	return e
}

// S sets the element's CSS properties
func (e *HTMLProgress) S(style StyleMap) *HTMLProgress {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLProgress) Key(key interface{}) *HTMLProgress {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLProgress) ID(v string) *HTMLProgress {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLProgress) Class(v string) *HTMLProgress {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLProgress) Title(v string) *HTMLProgress {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLProgress) Lang(v string) *HTMLProgress {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLProgress) Translate(v bool) *HTMLProgress {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLProgress) Dir(v string) *HTMLProgress {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLProgress) Hidden(v bool) *HTMLProgress {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLProgress) TabIndex(v int) *HTMLProgress {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLProgress) AccessKey(v string) *HTMLProgress {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLProgress) Draggable(v bool) *HTMLProgress {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLProgress) Spellcheck(v bool) *HTMLProgress {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
