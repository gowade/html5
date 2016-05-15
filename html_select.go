package h

// HTMLSelect represents HTML <select> tag
type HTMLSelect struct {
	HTMLElement
}

// Select creates an HTML <select> tag element
func Select() *HTMLSelect {
	e := &HTMLSelect{}
	e.a = make(map[string]interface{})
	e.tagName = "select"
	return e
}

// S sets the element's CSS properties
func (e *HTMLSelect) S(style StyleMap) *HTMLSelect {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLSelect) Key(key interface{}) *HTMLSelect {
	e.key = F(key)
	return e
}

// Autocomplete sets the element's "autocomplete" attribute
func (e *HTMLSelect) Autocomplete(v string) *HTMLSelect {
	e.a["autocomplete"] = v
	return e
}

// Autofocus sets the element's "autofocus" attribute
func (e *HTMLSelect) Autofocus(v bool) *HTMLSelect {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLSelect) Disabled(v bool) *HTMLSelect {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Multiple sets the element's "multiple" attribute
func (e *HTMLSelect) Multiple(v bool) *HTMLSelect {
	if v {
		e.a["multiple"] = ""
	} else {
		delete(e.a, "multiple")
	}
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLSelect) Name(v string) *HTMLSelect {
	e.a["name"] = v
	return e
}

// Required sets the element's "required" attribute
func (e *HTMLSelect) Required(v bool) *HTMLSelect {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

// Size sets the element's "size" attribute
func (e *HTMLSelect) Size(v int) *HTMLSelect {
	e.a["size"] = v
	return e
}

// Length sets the element's "length" attribute
func (e *HTMLSelect) Length(v int) *HTMLSelect {
	e.a["length"] = v
	return e
}

// SelectedIndex sets the element's "selectedindex" attribute
func (e *HTMLSelect) SelectedIndex(v int) *HTMLSelect {
	e.a["selectedindex"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLSelect) Value(v string) *HTMLSelect {
	e.a["value"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLSelect) CheckValidity(v bool) *HTMLSelect {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLSelect) ReportValidity(v bool) *HTMLSelect {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLSelect) ID(v string) *HTMLSelect {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLSelect) Class(v string) *HTMLSelect {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLSelect) Title(v string) *HTMLSelect {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLSelect) Lang(v string) *HTMLSelect {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLSelect) Translate(v bool) *HTMLSelect {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLSelect) Dir(v string) *HTMLSelect {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLSelect) Hidden(v bool) *HTMLSelect {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLSelect) TabIndex(v int) *HTMLSelect {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLSelect) AccessKey(v string) *HTMLSelect {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLSelect) Draggable(v bool) *HTMLSelect {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLSelect) Spellcheck(v bool) *HTMLSelect {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
