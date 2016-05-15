package h

// HTMLPicture represents HTML <picture> tag
type HTMLPicture struct {
	HTMLElement
}

// Picture creates a HTML <picture> tag
func Picture() *HTMLPicture {
	e := &HTMLPicture{}
	e.a = make(map[string]interface{})
	e.tagName = "picture"
	return e
}

// S sets the element's CSS properties
func (e *HTMLPicture) S(style StyleMap) *HTMLPicture {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLPicture) Key(key interface{}) *HTMLPicture {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLPicture) ID(v string) *HTMLPicture {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLPicture) Class(v string) *HTMLPicture {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLPicture) Title(v string) *HTMLPicture {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLPicture) Lang(v string) *HTMLPicture {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLPicture) Translate(v bool) *HTMLPicture {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLPicture) Dir(v string) *HTMLPicture {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLPicture) Hidden(v bool) *HTMLPicture {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLPicture) TabIndex(v int) *HTMLPicture {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLPicture) AccessKey(v string) *HTMLPicture {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLPicture) Draggable(v bool) *HTMLPicture {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLPicture) Spellcheck(v bool) *HTMLPicture {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
