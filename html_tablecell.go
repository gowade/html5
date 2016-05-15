package html5

// HTMLTableCell represents HTML <tablecell> tag
type HTMLTableCell struct {
	HTMLElement
}

// TableCell creates an HTML <tablecell> tag element
func TableCell() *HTMLTableCell {
	e := &HTMLTableCell{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecell"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTableCell) S(style StyleMap) *HTMLTableCell {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTableCell) Key(key interface{}) *HTMLTableCell {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLTableCell) Ref(dest *DOMElement) *HTMLTableCell {
	e.ref = dest
	return e
}

// ColSpan sets the element's "colspan" attribute
func (e *HTMLTableCell) ColSpan(v int) *HTMLTableCell {
	e.a["colspan"] = v
	return e
}

// RowSpan sets the element's "rowspan" attribute
func (e *HTMLTableCell) RowSpan(v int) *HTMLTableCell {
	e.a["rowspan"] = v
	return e
}

// Scope sets the element's "scope" attribute
func (e *HTMLTableCell) Scope(v string) *HTMLTableCell {
	e.a["scope"] = v
	return e
}

// Abbr sets the element's "abbr" attribute
func (e *HTMLTableCell) Abbr(v string) *HTMLTableCell {
	e.a["abbr"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTableCell) ID(v string) *HTMLTableCell {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTableCell) Class(v string) *HTMLTableCell {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTableCell) Title(v string) *HTMLTableCell {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTableCell) Lang(v string) *HTMLTableCell {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTableCell) Translate(v bool) *HTMLTableCell {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTableCell) Dir(v string) *HTMLTableCell {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTableCell) Hidden(v bool) *HTMLTableCell {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTableCell) TabIndex(v int) *HTMLTableCell {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTableCell) AccessKey(v string) *HTMLTableCell {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTableCell) Draggable(v bool) *HTMLTableCell {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTableCell) Spellcheck(v bool) *HTMLTableCell {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
