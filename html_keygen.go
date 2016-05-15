package h

// HTMLKeygen represents HTML <keygen> tag
type HTMLKeygen struct {
	HTMLElement
}

// Keygen creates a HTML <keygen> tag
func Keygen() *HTMLKeygen {
	e := &HTMLKeygen{}
	e.a = make(map[string]interface{})
	e.tagName = "keygen"
	return e
}

// S sets the element's CSS properties
func (e *HTMLKeygen) S(style StyleMap) *HTMLKeygen {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLKeygen) Key(key interface{}) *HTMLKeygen {
	e.key = F(key)
	return e
}

// Autofocus sets the element's "autofocus" attribute
func (e *HTMLKeygen) Autofocus(v bool) *HTMLKeygen {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

// Challenge sets the element's "challenge" attribute
func (e *HTMLKeygen) Challenge(v string) *HTMLKeygen {
	e.a["challenge"] = v
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLKeygen) Disabled(v bool) *HTMLKeygen {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Keytype sets the element's "keytype" attribute
func (e *HTMLKeygen) Keytype(v string) *HTMLKeygen {
	e.a["keytype"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLKeygen) Name(v string) *HTMLKeygen {
	e.a["name"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLKeygen) CheckValidity(v bool) *HTMLKeygen {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLKeygen) ReportValidity(v bool) *HTMLKeygen {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLKeygen) ID(v string) *HTMLKeygen {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLKeygen) Class(v string) *HTMLKeygen {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLKeygen) Title(v string) *HTMLKeygen {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLKeygen) Lang(v string) *HTMLKeygen {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLKeygen) Translate(v bool) *HTMLKeygen {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLKeygen) Dir(v string) *HTMLKeygen {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLKeygen) Hidden(v bool) *HTMLKeygen {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLKeygen) TabIndex(v int) *HTMLKeygen {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLKeygen) AccessKey(v string) *HTMLKeygen {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLKeygen) Draggable(v bool) *HTMLKeygen {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLKeygen) Spellcheck(v bool) *HTMLKeygen {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
