package h

// HTMLMenu represents HTML <menu> tag
type HTMLMenu struct {
	HTMLElement
}

// Menu creates an HTML <menu> tag element
func Menu() *HTMLMenu {
	e := &HTMLMenu{}
	e.a = make(map[string]interface{})
	e.tagName = "menu"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMenu) S(style StyleMap) *HTMLMenu {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMenu) Key(key interface{}) *HTMLMenu {
	e.key = F(key)
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLMenu) Type(v string) *HTMLMenu {
	e.a["type"] = v
	return e
}

// Label sets the element's "label" attribute
func (e *HTMLMenu) Label(v string) *HTMLMenu {
	e.a["label"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMenu) ID(v string) *HTMLMenu {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMenu) Class(v string) *HTMLMenu {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMenu) Title(v string) *HTMLMenu {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMenu) Lang(v string) *HTMLMenu {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMenu) Translate(v bool) *HTMLMenu {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMenu) Dir(v string) *HTMLMenu {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMenu) Hidden(v bool) *HTMLMenu {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMenu) TabIndex(v int) *HTMLMenu {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMenu) AccessKey(v string) *HTMLMenu {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMenu) Draggable(v bool) *HTMLMenu {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMenu) Spellcheck(v bool) *HTMLMenu {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
