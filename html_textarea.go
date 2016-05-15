package html5

// HTMLTextArea represents HTML <textarea> tag
type HTMLTextArea struct {
	HTMLElement
}

// TextArea creates an HTML <textarea> tag element
func TextArea() *HTMLTextArea {
	e := &HTMLTextArea{}
	e.a = make(map[string]interface{})
	e.tagName = "textarea"
	return e
}

// S sets the element's CSS properties
func (e *HTMLTextArea) S(style StyleMap) *HTMLTextArea {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLTextArea) Key(key interface{}) *HTMLTextArea {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLTextArea) Ref(dest *DOMElement) *HTMLTextArea {
	e.ref = dest
	return e
}

// Autocomplete sets the element's "autocomplete" attribute
func (e *HTMLTextArea) Autocomplete(v string) *HTMLTextArea {
	e.a["autocomplete"] = v
	return e
}

// Autofocus sets the element's "autofocus" attribute
func (e *HTMLTextArea) Autofocus(v bool) *HTMLTextArea {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

// Cols sets the element's "cols" attribute
func (e *HTMLTextArea) Cols(v int) *HTMLTextArea {
	e.a["cols"] = v
	return e
}

// DirName sets the element's "dirname" attribute
func (e *HTMLTextArea) DirName(v string) *HTMLTextArea {
	e.a["dirname"] = v
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLTextArea) Disabled(v bool) *HTMLTextArea {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// InputMode sets the element's "inputmode" attribute
func (e *HTMLTextArea) InputMode(v string) *HTMLTextArea {
	e.a["inputmode"] = v
	return e
}

// MaxLength sets the element's "maxlength" attribute
func (e *HTMLTextArea) MaxLength(v int) *HTMLTextArea {
	e.a["maxlength"] = v
	return e
}

// MinLength sets the element's "minlength" attribute
func (e *HTMLTextArea) MinLength(v int) *HTMLTextArea {
	e.a["minlength"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLTextArea) Name(v string) *HTMLTextArea {
	e.a["name"] = v
	return e
}

// Placeholder sets the element's "placeholder" attribute
func (e *HTMLTextArea) Placeholder(v string) *HTMLTextArea {
	e.a["placeholder"] = v
	return e
}

// ReadOnly sets the element's "readonly" attribute
func (e *HTMLTextArea) ReadOnly(v bool) *HTMLTextArea {
	if v {
		e.a["readonly"] = ""
	} else {
		delete(e.a, "readonly")
	}
	return e
}

// Required sets the element's "required" attribute
func (e *HTMLTextArea) Required(v bool) *HTMLTextArea {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

// Rows sets the element's "rows" attribute
func (e *HTMLTextArea) Rows(v int) *HTMLTextArea {
	e.a["rows"] = v
	return e
}

// Wrap sets the element's "wrap" attribute
func (e *HTMLTextArea) Wrap(v string) *HTMLTextArea {
	e.a["wrap"] = v
	return e
}

// DefaultValue sets the element's "defaultvalue" attribute
func (e *HTMLTextArea) DefaultValue(v string) *HTMLTextArea {
	e.a["defaultvalue"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLTextArea) Value(v string) *HTMLTextArea {
	e.a["value"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLTextArea) CheckValidity(v bool) *HTMLTextArea {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLTextArea) ReportValidity(v bool) *HTMLTextArea {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// SelectionStart sets the element's "selectionstart" attribute
func (e *HTMLTextArea) SelectionStart(v int) *HTMLTextArea {
	e.a["selectionstart"] = v
	return e
}

// SelectionEnd sets the element's "selectionend" attribute
func (e *HTMLTextArea) SelectionEnd(v int) *HTMLTextArea {
	e.a["selectionend"] = v
	return e
}

// SelectionDirection sets the element's "selectiondirection" attribute
func (e *HTMLTextArea) SelectionDirection(v string) *HTMLTextArea {
	e.a["selectiondirection"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLTextArea) ID(v string) *HTMLTextArea {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLTextArea) Class(v string) *HTMLTextArea {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLTextArea) Title(v string) *HTMLTextArea {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLTextArea) Lang(v string) *HTMLTextArea {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLTextArea) Translate(v bool) *HTMLTextArea {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLTextArea) Dir(v string) *HTMLTextArea {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLTextArea) Hidden(v bool) *HTMLTextArea {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLTextArea) TabIndex(v int) *HTMLTextArea {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLTextArea) AccessKey(v string) *HTMLTextArea {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLTextArea) Draggable(v bool) *HTMLTextArea {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLTextArea) Spellcheck(v bool) *HTMLTextArea {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
