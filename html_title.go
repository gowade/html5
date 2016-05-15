package h

// HTMLTitle represents HTML <title> tag
type HTMLTitle struct {
	HTMLElement
}

// Title creates a HTML <title> tag
func Title() *HTMLTitle {
	e := &HTMLTitle{}
	e.a = make(map[string]interface{})
	e.tagName = "title"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTitle) S(style StyleMap) *HTMLTitle {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTitle) Key(key interface{}) *HTMLTitle {
	e.key = F(key)
	return e
}

// Text sets the element's "text" attribute
func (e *HTMLTitle) Text(v string) *HTMLTitle {
	e.a["text"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTitle) ID(v string) *HTMLTitle {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTitle) Class(v string) *HTMLTitle {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTitle) Title(v string) *HTMLTitle {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTitle) Lang(v string) *HTMLTitle {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTitle) Translate(v bool) *HTMLTitle {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTitle) Dir(v string) *HTMLTitle {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTitle) Hidden(v bool) *HTMLTitle {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTitle) TabIndex(v int) *HTMLTitle {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTitle) AccessKey(v string) *HTMLTitle {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTitle) Draggable(v bool) *HTMLTitle {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTitle) Spellcheck(v bool) *HTMLTitle {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
