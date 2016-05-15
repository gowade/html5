package h

// HTMLTableSection represents HTML <tablesection> tag
type HTMLTableSection struct {
	HTMLElement
}

// TableSection creates an HTML <tablesection> tag element
func TableSection() *HTMLTableSection {
	e := &HTMLTableSection{}
	e.a = make(map[string]interface{})
	e.tagName = "tablesection"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTableSection) S(style StyleMap) *HTMLTableSection {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTableSection) Key(key interface{}) *HTMLTableSection {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTableSection) ID(v string) *HTMLTableSection {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTableSection) Class(v string) *HTMLTableSection {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTableSection) Title(v string) *HTMLTableSection {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTableSection) Lang(v string) *HTMLTableSection {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTableSection) Translate(v bool) *HTMLTableSection {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTableSection) Dir(v string) *HTMLTableSection {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTableSection) Hidden(v bool) *HTMLTableSection {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTableSection) TabIndex(v int) *HTMLTableSection {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTableSection) AccessKey(v string) *HTMLTableSection {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTableSection) Draggable(v bool) *HTMLTableSection {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTableSection) Spellcheck(v bool) *HTMLTableSection {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
