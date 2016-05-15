package h

// HTMLHeading represents HTML <heading> tag
type HTMLHeading struct {
	HTMLElement
}

// Heading creates a HTML <heading> tag
func Heading() *HTMLHeading {
	e := &HTMLHeading{}
	e.a = make(map[string]interface{})
	e.tagName = "heading"
	return e
}

// S sets the element's CSS properties
func (e *HTMLHeading) S(style StyleMap) *HTMLHeading {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLHeading) Key(key interface{}) *HTMLHeading {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLHeading) ID(v string) *HTMLHeading {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLHeading) Class(v string) *HTMLHeading {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLHeading) Title(v string) *HTMLHeading {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLHeading) Lang(v string) *HTMLHeading {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLHeading) Translate(v bool) *HTMLHeading {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLHeading) Dir(v string) *HTMLHeading {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLHeading) Hidden(v bool) *HTMLHeading {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLHeading) TabIndex(v int) *HTMLHeading {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLHeading) AccessKey(v string) *HTMLHeading {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLHeading) Draggable(v bool) *HTMLHeading {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLHeading) Spellcheck(v bool) *HTMLHeading {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
