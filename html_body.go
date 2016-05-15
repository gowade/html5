package h

// HTMLBody represents HTML <body> tag
type HTMLBody struct {
	HTMLElement
}

// Body creates a HTML <body> tag
func Body() *HTMLBody {
	e := &HTMLBody{}
	e.a = make(map[string]interface{})
	e.tagName = "body"
	return e
}

// S sets the element's CSS properties
func (e *HTMLBody) S(style StyleMap) *HTMLBody {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLBody) Key(key interface{}) *HTMLBody {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLBody) ID(v string) *HTMLBody {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLBody) Class(v string) *HTMLBody {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLBody) Title(v string) *HTMLBody {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLBody) Lang(v string) *HTMLBody {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLBody) Translate(v bool) *HTMLBody {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLBody) Dir(v string) *HTMLBody {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLBody) Hidden(v bool) *HTMLBody {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLBody) TabIndex(v int) *HTMLBody {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLBody) AccessKey(v string) *HTMLBody {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLBody) Draggable(v bool) *HTMLBody {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLBody) Spellcheck(v bool) *HTMLBody {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
