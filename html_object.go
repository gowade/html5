package h

// HTMLObject represents HTML <object> tag
type HTMLObject struct {
	HTMLElement
}

// Object creates an HTML <object> tag element
func Object() *HTMLObject {
	e := &HTMLObject{}
	e.a = make(map[string]interface{})
	e.tagName = "object"
	return e
}

// S sets the element's CSS properties
func (e *HTMLObject) S(style StyleMap) *HTMLObject {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLObject) Key(key interface{}) *HTMLObject {
	e.key = F(key)
	return e
}

// Data sets the element's "data" attribute
func (e *HTMLObject) Data(v string) *HTMLObject {
	e.a["data"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLObject) Type(v string) *HTMLObject {
	e.a["type"] = v
	return e
}

// TypeMustMatch sets the element's "typemustmatch" attribute
func (e *HTMLObject) TypeMustMatch(v bool) *HTMLObject {
	if v {
		e.a["typemustmatch"] = ""
	} else {
		delete(e.a, "typemustmatch")
	}
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLObject) Name(v string) *HTMLObject {
	e.a["name"] = v
	return e
}

// UseMap sets the element's "usemap" attribute
func (e *HTMLObject) UseMap(v string) *HTMLObject {
	e.a["usemap"] = v
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLObject) Width(v string) *HTMLObject {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLObject) Height(v string) *HTMLObject {
	e.a["height"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLObject) CheckValidity(v bool) *HTMLObject {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLObject) ReportValidity(v bool) *HTMLObject {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLObject) ID(v string) *HTMLObject {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLObject) Class(v string) *HTMLObject {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLObject) Title(v string) *HTMLObject {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLObject) Lang(v string) *HTMLObject {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLObject) Translate(v bool) *HTMLObject {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLObject) Dir(v string) *HTMLObject {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLObject) Hidden(v bool) *HTMLObject {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLObject) TabIndex(v int) *HTMLObject {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLObject) AccessKey(v string) *HTMLObject {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLObject) Draggable(v bool) *HTMLObject {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLObject) Spellcheck(v bool) *HTMLObject {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
