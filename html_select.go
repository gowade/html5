package h

type htmlSelect struct {
	htmlElement
}

func Select() *htmlSelect {
	e := &htmlSelect{}
	e.a = make(map[string]interface{})
	e.tagName = "select"
	return e
}

func (e *htmlSelect) S(style StyleMap) *htmlSelect {
	e.htmlElement.S(style)
	return e
}

func (e *htmlSelect) Key(key interface{}) *htmlSelect {
	e.key = F(key)
	return e
}

func (e *htmlSelect) Autocomplete(v string) *htmlSelect {
	e.a["autocomplete"] = v
	return e
}

func (e *htmlSelect) Autofocus(v bool) *htmlSelect {
	if v {
		e.a["autofocus"] = ""
	} else {
		delete(e.a, "autofocus")
	}
	return e
}

func (e *htmlSelect) Disabled(v bool) *htmlSelect {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

func (e *htmlSelect) Multiple(v bool) *htmlSelect {
	if v {
		e.a["multiple"] = ""
	} else {
		delete(e.a, "multiple")
	}
	return e
}

func (e *htmlSelect) Name(v string) *htmlSelect {
	e.a["name"] = v
	return e
}

func (e *htmlSelect) Required(v bool) *htmlSelect {
	if v {
		e.a["required"] = ""
	} else {
		delete(e.a, "required")
	}
	return e
}

func (e *htmlSelect) Size(v int) *htmlSelect {
	e.a["size"] = v
	return e
}

func (e *htmlSelect) Length(v int) *htmlSelect {
	e.a["length"] = v
	return e
}

func (e *htmlSelect) SelectedIndex(v int) *htmlSelect {
	e.a["selectedindex"] = v
	return e
}

func (e *htmlSelect) Value(v string) *htmlSelect {
	e.a["value"] = v
	return e
}

func (e *htmlSelect) CheckValidity(v bool) *htmlSelect {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

func (e *htmlSelect) ReportValidity(v bool) *htmlSelect {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

func (e *htmlSelect) ID(v string) *htmlSelect {
	e.a["id"] = v
	return e
}

func (e *htmlSelect) Class(v string) *htmlSelect {
	e.a["class"] = v
	return e
}

func (e *htmlSelect) Title(v string) *htmlSelect {
	e.a["title"] = v
	return e
}

func (e *htmlSelect) Lang(v string) *htmlSelect {
	e.a["lang"] = v
	return e
}

func (e *htmlSelect) Translate(v bool) *htmlSelect {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlSelect) Dir(v string) *htmlSelect {
	e.a["dir"] = v
	return e
}

func (e *htmlSelect) Hidden(v bool) *htmlSelect {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlSelect) TabIndex(v int) *htmlSelect {
	e.a["tabindex"] = v
	return e
}

func (e *htmlSelect) AccessKey(v string) *htmlSelect {
	e.a["accesskey"] = v
	return e
}

func (e *htmlSelect) Draggable(v bool) *htmlSelect {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlSelect) Spellcheck(v bool) *htmlSelect {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
