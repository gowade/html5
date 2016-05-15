package h

// HTMLDirectory represents HTML <directory> tag
type HTMLDirectory struct {
	HTMLElement
}

// Directory creates an HTML <directory> tag element
func Directory() *HTMLDirectory {
	e := &HTMLDirectory{}
	e.a = make(map[string]interface{})
	e.tagName = "directory"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDirectory) S(style StyleMap) *HTMLDirectory {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDirectory) Key(key interface{}) *HTMLDirectory {
	e.key = F(key)
	return e
}

// Compact sets the element's "compact" attribute
func (e *HTMLDirectory) Compact(v bool) *HTMLDirectory {
	if v {
		e.a["compact"] = ""
	} else {
		delete(e.a, "compact")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDirectory) ID(v string) *HTMLDirectory {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDirectory) Class(v string) *HTMLDirectory {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDirectory) Title(v string) *HTMLDirectory {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDirectory) Lang(v string) *HTMLDirectory {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDirectory) Translate(v bool) *HTMLDirectory {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDirectory) Dir(v string) *HTMLDirectory {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDirectory) Hidden(v bool) *HTMLDirectory {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDirectory) TabIndex(v int) *HTMLDirectory {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDirectory) AccessKey(v string) *HTMLDirectory {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDirectory) Draggable(v bool) *HTMLDirectory {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDirectory) Spellcheck(v bool) *HTMLDirectory {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
