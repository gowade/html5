package h

// HTMLDialog represents HTML <dialog> tag
type HTMLDialog struct {
	HTMLElement
}

// Dialog creates a HTML <dialog> tag
func Dialog() *HTMLDialog {
	e := &HTMLDialog{}
	e.a = make(map[string]interface{})
	e.tagName = "dialog"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDialog) S(style StyleMap) *HTMLDialog {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDialog) Key(key interface{}) *HTMLDialog {
	e.key = F(key)
	return e
}

// Open sets the element's "open" attribute
func (e *HTMLDialog) Open(v bool) *HTMLDialog {
	if v {
		e.a["open"] = ""
	} else {
		delete(e.a, "open")
	}
	return e
}

// ReturnValue sets the element's "returnvalue" attribute
func (e *HTMLDialog) ReturnValue(v string) *HTMLDialog {
	e.a["returnvalue"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDialog) ID(v string) *HTMLDialog {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDialog) Class(v string) *HTMLDialog {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDialog) Title(v string) *HTMLDialog {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDialog) Lang(v string) *HTMLDialog {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDialog) Translate(v bool) *HTMLDialog {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDialog) Dir(v string) *HTMLDialog {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDialog) Hidden(v bool) *HTMLDialog {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDialog) TabIndex(v int) *HTMLDialog {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDialog) AccessKey(v string) *HTMLDialog {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDialog) Draggable(v bool) *HTMLDialog {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDialog) Spellcheck(v bool) *HTMLDialog {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
