package h

// HTMLTable represents HTML <table> tag
type HTMLTable struct {
	HTMLElement
}

// Table creates an HTML <table> tag element
func Table() *HTMLTable {
	e := &HTMLTable{}
	e.a = make(map[string]interface{})
	e.tagName = "table"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTable) S(style StyleMap) *HTMLTable {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTable) Key(key interface{}) *HTMLTable {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTable) ID(v string) *HTMLTable {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTable) Class(v string) *HTMLTable {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTable) Title(v string) *HTMLTable {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTable) Lang(v string) *HTMLTable {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTable) Translate(v bool) *HTMLTable {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTable) Dir(v string) *HTMLTable {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTable) Hidden(v bool) *HTMLTable {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTable) TabIndex(v int) *HTMLTable {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTable) AccessKey(v string) *HTMLTable {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTable) Draggable(v bool) *HTMLTable {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTable) Spellcheck(v bool) *HTMLTable {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
