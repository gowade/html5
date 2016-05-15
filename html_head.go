package h

// HTMLHead represents HTML <head> tag
type HTMLHead struct {
	HTMLElement
}

// Head creates a HTML <head> tag
func Head() *HTMLHead {
	e := &HTMLHead{}
	e.a = make(map[string]interface{})
	e.tagName = "head"
	return e
}

// S sets the element's CSS properties
func (e *HTMLHead) S(style StyleMap) *HTMLHead {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLHead) Key(key interface{}) *HTMLHead {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLHead) ID(v string) *HTMLHead {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLHead) Class(v string) *HTMLHead {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLHead) Title(v string) *HTMLHead {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLHead) Lang(v string) *HTMLHead {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLHead) Translate(v bool) *HTMLHead {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLHead) Dir(v string) *HTMLHead {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLHead) Hidden(v bool) *HTMLHead {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLHead) TabIndex(v int) *HTMLHead {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLHead) AccessKey(v string) *HTMLHead {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLHead) Draggable(v bool) *HTMLHead {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLHead) Spellcheck(v bool) *HTMLHead {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
