package h

// HTMLMod represents HTML <mod> tag
type HTMLMod struct {
	HTMLElement
}

// Mod creates a HTML <mod> tag
func Mod() *HTMLMod {
	e := &HTMLMod{}
	e.a = make(map[string]interface{})
	e.tagName = "mod"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMod) S(style StyleMap) *HTMLMod {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMod) Key(key interface{}) *HTMLMod {
	e.key = F(key)
	return e
}

// Cite sets the element's "cite" attribute
func (e *HTMLMod) Cite(v string) *HTMLMod {
	e.a["cite"] = v
	return e
}

// DateTime sets the element's "datetime" attribute
func (e *HTMLMod) DateTime(v string) *HTMLMod {
	e.a["datetime"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMod) ID(v string) *HTMLMod {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMod) Class(v string) *HTMLMod {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMod) Title(v string) *HTMLMod {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMod) Lang(v string) *HTMLMod {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMod) Translate(v bool) *HTMLMod {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMod) Dir(v string) *HTMLMod {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMod) Hidden(v bool) *HTMLMod {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMod) TabIndex(v int) *HTMLMod {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMod) AccessKey(v string) *HTMLMod {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMod) Draggable(v bool) *HTMLMod {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMod) Spellcheck(v bool) *HTMLMod {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
