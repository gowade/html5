package h

// HTMLButton represents HTML <button> tag
type HTMLButton struct {
	HTMLElement
}

// Button creates a HTML <button> tag
func Button() *HTMLButton {
	e := &HTMLButton{}
	e.a = make(map[string]interface{})
	e.tagName = "button"
	return e
}

// S sets the element's CSS properties
func (e *HTMLButton) S(style StyleMap) *HTMLButton {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLButton) Key(key interface{}) *HTMLButton {
	e.key = F(key)
	return e
}

// Autofocus sets the element's "autofocus" attribute
func (e *HTMLButton) Autofocus(v bool) *HTMLButton {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLButton) Disabled(v bool) *HTMLButton {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// FormAction sets the element's "formaction" attribute
func (e *HTMLButton) FormAction(v string) *HTMLButton {
	e.a["formaction"] = v
	return e
}

// FormEnctype sets the element's "formenctype" attribute
func (e *HTMLButton) FormEnctype(v string) *HTMLButton {
	e.a["formenctype"] = v
	return e
}

// FormMethod sets the element's "formmethod" attribute
func (e *HTMLButton) FormMethod(v string) *HTMLButton {
	e.a["formmethod"] = v
	return e
}

// FormNoValidate sets the element's "formnovalidate" attribute
func (e *HTMLButton) FormNoValidate(v bool) *HTMLButton {
	if v {
		e.a["formnovalidate"] = ""
	} else {
		delete(e.a, "formnovalidate")
	}
	return e
}

// FormTarget sets the element's "formtarget" attribute
func (e *HTMLButton) FormTarget(v string) *HTMLButton {
	e.a["formtarget"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLButton) Name(v string) *HTMLButton {
	e.a["name"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLButton) Type(v string) *HTMLButton {
	e.a["type"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLButton) Value(v string) *HTMLButton {
	e.a["value"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLButton) CheckValidity(v bool) *HTMLButton {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLButton) ReportValidity(v bool) *HTMLButton {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLButton) ID(v string) *HTMLButton {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLButton) Class(v string) *HTMLButton {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLButton) Title(v string) *HTMLButton {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLButton) Lang(v string) *HTMLButton {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLButton) Translate(v bool) *HTMLButton {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLButton) Dir(v string) *HTMLButton {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLButton) Hidden(v bool) *HTMLButton {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLButton) TabIndex(v int) *HTMLButton {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLButton) AccessKey(v string) *HTMLButton {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLButton) Draggable(v bool) *HTMLButton {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLButton) Spellcheck(v bool) *HTMLButton {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
