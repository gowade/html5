package h

type htmlTextArea struct {
	htmlElement
}

func TextArea() *htmlTextArea {
	e := &htmlTextArea{}
	e.a = make(map[string]interface{})
	e.tagName = "textarea"
	return e
}

func (e *htmlTextArea) S(style StyleMap) *htmlTextArea {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTextArea) Key(key interface{}) *htmlTextArea {
	e.key = F(key)
	return e
}

func (e *htmlTextArea) Autocomplete(v string) *htmlTextArea {
	e.a["autocomplete"] = v
	return e
}

func (e *htmlTextArea) Autofocus(v bool) *htmlTextArea {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

func (e *htmlTextArea) Cols(v int) *htmlTextArea {
	e.a["cols"] = v
	return e
}

func (e *htmlTextArea) DirName(v string) *htmlTextArea {
	e.a["dirname"] = v
	return e
}

func (e *htmlTextArea) Disabled(v bool) *htmlTextArea {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlTextArea) InputMode(v string) *htmlTextArea {
	e.a["inputmode"] = v
	return e
}

func (e *htmlTextArea) MaxLength(v int) *htmlTextArea {
	e.a["maxlength"] = v
	return e
}

func (e *htmlTextArea) MinLength(v int) *htmlTextArea {
	e.a["minlength"] = v
	return e
}

func (e *htmlTextArea) Name(v string) *htmlTextArea {
	e.a["name"] = v
	return e
}

func (e *htmlTextArea) Placeholder(v string) *htmlTextArea {
	e.a["placeholder"] = v
	return e
}

func (e *htmlTextArea) ReadOnly(v bool) *htmlTextArea {
	if v {
		e.a["readonly"] = ""
	} else {
		delete(e.a, "readonly")
	}
	return e
}

func (e *htmlTextArea) Required(v bool) *htmlTextArea {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

func (e *htmlTextArea) Rows(v int) *htmlTextArea {
	e.a["rows"] = v
	return e
}

func (e *htmlTextArea) Wrap(v string) *htmlTextArea {
	e.a["wrap"] = v
	return e
}

func (e *htmlTextArea) DefaultValue(v string) *htmlTextArea {
	e.a["defaultvalue"] = v
	return e
}

func (e *htmlTextArea) Value(v string) *htmlTextArea {
	e.a["value"] = v
	return e
}

func (e *htmlTextArea) CheckValidity(v bool) *htmlTextArea {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlTextArea) ReportValidity(v bool) *htmlTextArea {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlTextArea) SelectionStart(v int) *htmlTextArea {
	e.a["selectionstart"] = v
	return e
}

func (e *htmlTextArea) SelectionEnd(v int) *htmlTextArea {
	e.a["selectionend"] = v
	return e
}

func (e *htmlTextArea) SelectionDirection(v string) *htmlTextArea {
	e.a["selectiondirection"] = v
	return e
}

func (e *htmlTextArea) ID(v string) *htmlTextArea {
	e.a["id"] = v
	return e
}

func (e *htmlTextArea) Class(v string) *htmlTextArea {
	e.a["class"] = v
	return e
}

func (e *htmlTextArea) Title(v string) *htmlTextArea {
	e.a["title"] = v
	return e
}

func (e *htmlTextArea) Lang(v string) *htmlTextArea {
	e.a["lang"] = v
	return e
}

func (e *htmlTextArea) Translate(v bool) *htmlTextArea {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTextArea) Dir(v string) *htmlTextArea {
	e.a["dir"] = v
	return e
}

func (e *htmlTextArea) Hidden(v bool) *htmlTextArea {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTextArea) TabIndex(v int) *htmlTextArea {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTextArea) AccessKey(v string) *htmlTextArea {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTextArea) Draggable(v bool) *htmlTextArea {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTextArea) Spellcheck(v bool) *htmlTextArea {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
