package h

// HTMLInput represents HTML <input> tag
type HTMLInput struct {
	HTMLElement
}

// Input creates a HTML <input> tag
func Input() *HTMLInput {
	e := &HTMLInput{}
	e.a = make(map[string]interface{})
	e.tagName = "input"
	return e
}

// S sets the element's CSS properties
func (e *HTMLInput) S(style StyleMap) *HTMLInput {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLInput) Key(key interface{}) *HTMLInput {
	e.key = F(key)
	return e
}

// Accept sets the element's "accept" attribute
func (e *HTMLInput) Accept(v string) *HTMLInput {
	e.a["accept"] = v
	return e
}

// Alt sets the element's "alt" attribute
func (e *HTMLInput) Alt(v string) *HTMLInput {
	e.a["alt"] = v
	return e
}

// Autocomplete sets the element's "autocomplete" attribute
func (e *HTMLInput) Autocomplete(v string) *HTMLInput {
	e.a["autocomplete"] = v
	return e
}

// Autofocus sets the element's "autofocus" attribute
func (e *HTMLInput) Autofocus(v bool) *HTMLInput {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

// DefaultChecked sets the element's "defaultchecked" attribute
func (e *HTMLInput) DefaultChecked(v bool) *HTMLInput {
	if v {
		e.a["defaultchecked"] = ""
	} else {
		delete(e.a, "defaultchecked")
	}
	return e
}

// Checked sets the element's "checked" attribute
func (e *HTMLInput) Checked(v bool) *HTMLInput {
	if v {
		e.a["checked"] = ""
	} else {
		delete(e.a, "checked")
	}
	return e
}

// DirName sets the element's "dirname" attribute
func (e *HTMLInput) DirName(v string) *HTMLInput {
	e.a["dirname"] = v
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLInput) Disabled(v bool) *HTMLInput {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// FormAction sets the element's "formaction" attribute
func (e *HTMLInput) FormAction(v string) *HTMLInput {
	e.a["formaction"] = v
	return e
}

// FormEnctype sets the element's "formenctype" attribute
func (e *HTMLInput) FormEnctype(v string) *HTMLInput {
	e.a["formenctype"] = v
	return e
}

// FormMethod sets the element's "formmethod" attribute
func (e *HTMLInput) FormMethod(v string) *HTMLInput {
	e.a["formmethod"] = v
	return e
}

// FormNoValidate sets the element's "formnovalidate" attribute
func (e *HTMLInput) FormNoValidate(v bool) *HTMLInput {
	if v {
		e.a["formnovalidate"] = ""
	} else {
		delete(e.a, "formnovalidate")
	}
	return e
}

// FormTarget sets the element's "formtarget" attribute
func (e *HTMLInput) FormTarget(v string) *HTMLInput {
	e.a["formtarget"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLInput) Height(v int) *HTMLInput {
	e.a["height"] = v
	return e
}

// Indeterminate sets the element's "indeterminate" attribute
func (e *HTMLInput) Indeterminate(v bool) *HTMLInput {
	if v {
		e.a["indeterminate"] = ""
	} else {
		delete(e.a, "indeterminate")
	}
	return e
}

// InputMode sets the element's "inputmode" attribute
func (e *HTMLInput) InputMode(v string) *HTMLInput {
	e.a["inputmode"] = v
	return e
}

// Max sets the element's "max" attribute
func (e *HTMLInput) Max(v string) *HTMLInput {
	e.a["max"] = v
	return e
}

// MaxLength sets the element's "maxlength" attribute
func (e *HTMLInput) MaxLength(v int) *HTMLInput {
	e.a["maxlength"] = v
	return e
}

// Min sets the element's "min" attribute
func (e *HTMLInput) Min(v string) *HTMLInput {
	e.a["min"] = v
	return e
}

// MinLength sets the element's "minlength" attribute
func (e *HTMLInput) MinLength(v int) *HTMLInput {
	e.a["minlength"] = v
	return e
}

// Multiple sets the element's "multiple" attribute
func (e *HTMLInput) Multiple(v bool) *HTMLInput {
	if v {
		e.a["multiple"] = ""
	} else {
		delete(e.a, "multiple")
	}
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLInput) Name(v string) *HTMLInput {
	e.a["name"] = v
	return e
}

// Pattern sets the element's "pattern" attribute
func (e *HTMLInput) Pattern(v string) *HTMLInput {
	e.a["pattern"] = v
	return e
}

// Placeholder sets the element's "placeholder" attribute
func (e *HTMLInput) Placeholder(v string) *HTMLInput {
	e.a["placeholder"] = v
	return e
}

// ReadOnly sets the element's "readonly" attribute
func (e *HTMLInput) ReadOnly(v bool) *HTMLInput {
	if v {
		e.a["readonly"] = ""
	} else {
		delete(e.a, "readonly")
	}
	return e
}

// Required sets the element's "required" attribute
func (e *HTMLInput) Required(v bool) *HTMLInput {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

// Size sets the element's "size" attribute
func (e *HTMLInput) Size(v int) *HTMLInput {
	e.a["size"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLInput) Src(v string) *HTMLInput {
	e.a["src"] = v
	return e
}

// Step sets the element's "step" attribute
func (e *HTMLInput) Step(v string) *HTMLInput {
	e.a["step"] = v
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLInput) Type(v string) *HTMLInput {
	e.a["type"] = v
	return e
}

// DefaultValue sets the element's "defaultvalue" attribute
func (e *HTMLInput) DefaultValue(v string) *HTMLInput {
	e.a["defaultvalue"] = v
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLInput) Value(v string) *HTMLInput {
	e.a["value"] = v
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLInput) Width(v int) *HTMLInput {
	e.a["width"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLInput) CheckValidity(v bool) *HTMLInput {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLInput) ReportValidity(v bool) *HTMLInput {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// SelectionStart sets the element's "selectionstart" attribute
func (e *HTMLInput) SelectionStart(v int) *HTMLInput {
	e.a["selectionstart"] = v
	return e
}

// SelectionEnd sets the element's "selectionend" attribute
func (e *HTMLInput) SelectionEnd(v int) *HTMLInput {
	e.a["selectionend"] = v
	return e
}

// SelectionDirection sets the element's "selectiondirection" attribute
func (e *HTMLInput) SelectionDirection(v string) *HTMLInput {
	e.a["selectiondirection"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLInput) ID(v string) *HTMLInput {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLInput) Class(v string) *HTMLInput {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLInput) Title(v string) *HTMLInput {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLInput) Lang(v string) *HTMLInput {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLInput) Translate(v bool) *HTMLInput {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLInput) Dir(v string) *HTMLInput {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLInput) Hidden(v bool) *HTMLInput {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLInput) TabIndex(v int) *HTMLInput {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLInput) AccessKey(v string) *HTMLInput {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLInput) Draggable(v bool) *HTMLInput {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLInput) Spellcheck(v bool) *HTMLInput {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
