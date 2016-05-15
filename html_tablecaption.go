package h

// HTMLTableCaption represents HTML <tablecaption> tag
type HTMLTableCaption struct {
	HTMLElement
}

// TableCaption creates a HTML <tablecaption> tag
func TableCaption() *HTMLTableCaption {
	e := &HTMLTableCaption{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecaption"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTableCaption) S(style StyleMap) *HTMLTableCaption {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTableCaption) Key(key interface{}) *HTMLTableCaption {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTableCaption) ID(v string) *HTMLTableCaption {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTableCaption) Class(v string) *HTMLTableCaption {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTableCaption) Title(v string) *HTMLTableCaption {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTableCaption) Lang(v string) *HTMLTableCaption {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTableCaption) Translate(v bool) *HTMLTableCaption {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTableCaption) Dir(v string) *HTMLTableCaption {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTableCaption) Hidden(v bool) *HTMLTableCaption {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTableCaption) TabIndex(v int) *HTMLTableCaption {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTableCaption) AccessKey(v string) *HTMLTableCaption {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTableCaption) Draggable(v bool) *HTMLTableCaption {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTableCaption) Spellcheck(v bool) *HTMLTableCaption {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
