package h

// HTMLFieldSet represents HTML <fieldset> tag
type HTMLFieldSet struct {
	HTMLElement
}

// FieldSet creates a HTML <fieldset> tag
func FieldSet() *HTMLFieldSet {
	e := &HTMLFieldSet{}
	e.a = make(map[string]interface{})
	e.tagName = "fieldset"
	return e
}

// S sets the element's CSS properties
func (e *HTMLFieldSet) S(style StyleMap) *HTMLFieldSet {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLFieldSet) Key(key interface{}) *HTMLFieldSet {
	e.key = F(key)
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLFieldSet) Disabled(v bool) *HTMLFieldSet {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLFieldSet) Name(v string) *HTMLFieldSet {
	e.a["name"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLFieldSet) CheckValidity(v bool) *HTMLFieldSet {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLFieldSet) ReportValidity(v bool) *HTMLFieldSet {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLFieldSet) ID(v string) *HTMLFieldSet {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLFieldSet) Class(v string) *HTMLFieldSet {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLFieldSet) Title(v string) *HTMLFieldSet {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLFieldSet) Lang(v string) *HTMLFieldSet {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLFieldSet) Translate(v bool) *HTMLFieldSet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLFieldSet) Dir(v string) *HTMLFieldSet {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLFieldSet) Hidden(v bool) *HTMLFieldSet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLFieldSet) TabIndex(v int) *HTMLFieldSet {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLFieldSet) AccessKey(v string) *HTMLFieldSet {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLFieldSet) Draggable(v bool) *HTMLFieldSet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLFieldSet) Spellcheck(v bool) *HTMLFieldSet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
