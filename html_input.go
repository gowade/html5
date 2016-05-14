package h

type htmlInput struct {
	htmlElement
}

func Input() *htmlInput {
	e := &htmlInput{}
	e.a = make(map[string]interface{})
	e.tagName = "input"
	return e
}

func (e *htmlInput) S(style StyleMap) *htmlInput {
	e.htmlElement.S(style)
	return e
}

func (e *htmlInput) Key(key interface{}) *htmlInput {
	e.key = F(key)
	return e
}

func (e *htmlInput) Accept(v string) *htmlInput {
	e.a["accept"] = v
	return e
}

func (e *htmlInput) Alt(v string) *htmlInput {
	e.a["alt"] = v
	return e
}

func (e *htmlInput) Autocomplete(v string) *htmlInput {
	e.a["autocomplete"] = v
	return e
}

func (e *htmlInput) Autofocus(v bool) *htmlInput {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

func (e *htmlInput) DefaultChecked(v bool) *htmlInput {
	if v {
		e.a["defaultchecked"] = ""
	} else {
		delete(e.a, "defaultchecked")
	}
	return e
}

func (e *htmlInput) Checked(v bool) *htmlInput {
	if v {
		e.a["checked"] = ""
	} else {
		delete(e.a, "checked")
	}
	return e
}

func (e *htmlInput) DirName(v string) *htmlInput {
	e.a["dirname"] = v
	return e
}

func (e *htmlInput) Disabled(v bool) *htmlInput {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlInput) FormAction(v string) *htmlInput {
	e.a["formaction"] = v
	return e
}

func (e *htmlInput) FormEnctype(v string) *htmlInput {
	e.a["formenctype"] = v
	return e
}

func (e *htmlInput) FormMethod(v string) *htmlInput {
	e.a["formmethod"] = v
	return e
}

func (e *htmlInput) FormNoValidate(v bool) *htmlInput {
	if v {
		e.a["formnovalidate"] = ""
	} else {
		delete(e.a, "formnovalidate")
	}
	return e
}

func (e *htmlInput) FormTarget(v string) *htmlInput {
	e.a["formtarget"] = v
	return e
}

func (e *htmlInput) Height(v int) *htmlInput {
	e.a["height"] = v
	return e
}

func (e *htmlInput) Indeterminate(v bool) *htmlInput {
	if v {
		e.a["indeterminate"] = ""
	} else {
		delete(e.a, "indeterminate")
	}
	return e
}

func (e *htmlInput) InputMode(v string) *htmlInput {
	e.a["inputmode"] = v
	return e
}

func (e *htmlInput) Max(v string) *htmlInput {
	e.a["max"] = v
	return e
}

func (e *htmlInput) MaxLength(v int) *htmlInput {
	e.a["maxlength"] = v
	return e
}

func (e *htmlInput) Min(v string) *htmlInput {
	e.a["min"] = v
	return e
}

func (e *htmlInput) MinLength(v int) *htmlInput {
	e.a["minlength"] = v
	return e
}

func (e *htmlInput) Multiple(v bool) *htmlInput {
	if v {
		e.a["multiple"] = ""
	} else {
		delete(e.a, "multiple")
	}
	return e
}

func (e *htmlInput) Name(v string) *htmlInput {
	e.a["name"] = v
	return e
}

func (e *htmlInput) Pattern(v string) *htmlInput {
	e.a["pattern"] = v
	return e
}

func (e *htmlInput) Placeholder(v string) *htmlInput {
	e.a["placeholder"] = v
	return e
}

func (e *htmlInput) ReadOnly(v bool) *htmlInput {
	if v {
		e.a["readonly"] = ""
	} else {
		delete(e.a, "readonly")
	}
	return e
}

func (e *htmlInput) Required(v bool) *htmlInput {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

func (e *htmlInput) Size(v int) *htmlInput {
	e.a["size"] = v
	return e
}

func (e *htmlInput) Src(v string) *htmlInput {
	e.a["src"] = v
	return e
}

func (e *htmlInput) Step(v string) *htmlInput {
	e.a["step"] = v
	return e
}

func (e *htmlInput) Type(v string) *htmlInput {
	e.a["type"] = v
	return e
}

func (e *htmlInput) DefaultValue(v string) *htmlInput {
	e.a["defaultvalue"] = v
	return e
}

func (e *htmlInput) Value(v string) *htmlInput {
	e.a["value"] = v
	return e
}

func (e *htmlInput) Width(v int) *htmlInput {
	e.a["width"] = v
	return e
}

func (e *htmlInput) CheckValidity(v bool) *htmlInput {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlInput) ReportValidity(v bool) *htmlInput {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlInput) SelectionStart(v int) *htmlInput {
	e.a["selectionstart"] = v
	return e
}

func (e *htmlInput) SelectionEnd(v int) *htmlInput {
	e.a["selectionend"] = v
	return e
}

func (e *htmlInput) SelectionDirection(v string) *htmlInput {
	e.a["selectiondirection"] = v
	return e
}

func (e *htmlInput) ID(v string) *htmlInput {
	e.a["id"] = v
	return e
}

func (e *htmlInput) Class(v string) *htmlInput {
	e.a["class"] = v
	return e
}

func (e *htmlInput) Title(v string) *htmlInput {
	e.a["title"] = v
	return e
}

func (e *htmlInput) Lang(v string) *htmlInput {
	e.a["lang"] = v
	return e
}

func (e *htmlInput) Translate(v bool) *htmlInput {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlInput) Dir(v string) *htmlInput {
	e.a["dir"] = v
	return e
}

func (e *htmlInput) Hidden(v bool) *htmlInput {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlInput) TabIndex(v int) *htmlInput {
	e.a["tabindex"] = v
	return e
}

func (e *htmlInput) AccessKey(v string) *htmlInput {
	e.a["accesskey"] = v
	return e
}

func (e *htmlInput) Draggable(v bool) *htmlInput {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlInput) Spellcheck(v bool) *htmlInput {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
