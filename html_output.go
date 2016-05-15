package h

// HTMLOutput represents HTML <output> tag
type HTMLOutput struct {
	HTMLElement
}

// Output creates an HTML <output> tag element
func Output() *HTMLOutput {
	e := &HTMLOutput{}
	e.a = make(map[string]interface{})
	e.tagName = "output"
	return e
}

// S sets the element's CSS properties
func (e *HTMLOutput) S(style StyleMap) *HTMLOutput {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLOutput) Key(key interface{}) *HTMLOutput {
	e.key = F(key)
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLOutput) Name(v string) *HTMLOutput {
	e.a["name"] = v
	return e
}

// DefaultValue sets the element's "defaultvalue" attribute
func (e *HTMLOutput) DefaultValue(v string) *HTMLOutput {
	e.a["defaultvalue"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLOutput) Value(v string) *HTMLOutput {
	e.a["value"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLOutput) CheckValidity(v bool) *HTMLOutput {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLOutput) ReportValidity(v bool) *HTMLOutput {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLOutput) ID(v string) *HTMLOutput {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLOutput) Class(v string) *HTMLOutput {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLOutput) Title(v string) *HTMLOutput {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLOutput) Lang(v string) *HTMLOutput {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLOutput) Translate(v bool) *HTMLOutput {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLOutput) Dir(v string) *HTMLOutput {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLOutput) Hidden(v bool) *HTMLOutput {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLOutput) TabIndex(v int) *HTMLOutput {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLOutput) AccessKey(v string) *HTMLOutput {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLOutput) Draggable(v bool) *HTMLOutput {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLOutput) Spellcheck(v bool) *HTMLOutput {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
