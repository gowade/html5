package h

// HTMLTableRow represents HTML <tablerow> tag
type HTMLTableRow struct {
	HTMLElement
}

// TableRow creates an HTML <tablerow> tag element
func TableRow() *HTMLTableRow {
	e := &HTMLTableRow{}
	e.a = make(map[string]interface{})
	e.tagName = "tablerow"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTableRow) S(style StyleMap) *HTMLTableRow {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTableRow) Key(key interface{}) *HTMLTableRow {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTableRow) ID(v string) *HTMLTableRow {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTableRow) Class(v string) *HTMLTableRow {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTableRow) Title(v string) *HTMLTableRow {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTableRow) Lang(v string) *HTMLTableRow {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTableRow) Translate(v bool) *HTMLTableRow {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTableRow) Dir(v string) *HTMLTableRow {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTableRow) Hidden(v bool) *HTMLTableRow {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTableRow) TabIndex(v int) *HTMLTableRow {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTableRow) AccessKey(v string) *HTMLTableRow {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTableRow) Draggable(v bool) *HTMLTableRow {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTableRow) Spellcheck(v bool) *HTMLTableRow {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
