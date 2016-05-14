package h

type htmlDialog struct {
	htmlElement
}

func Dialog() *htmlDialog {
	e := &htmlDialog{}
	e.a = make(map[string]interface{})
	e.tagName = "dialog"
	return e
}

func (e *htmlDialog) S(style StyleMap) *htmlDialog {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDialog) Key(key interface{}) *htmlDialog {
	e.key = F(key)
	return e
}

func (e *htmlDialog) Open(v bool) *htmlDialog {
	if v {
		e.a["open"] = ""
	} else {
		delete(e.a, "open")
	}
	return e
}

func (e *htmlDialog) ReturnValue(v string) *htmlDialog {
	e.a["returnvalue"] = v
	return e
}

func (e *htmlDialog) ID(v string) *htmlDialog {
	e.a["id"] = v
	return e
}

func (e *htmlDialog) Class(v string) *htmlDialog {
	e.a["class"] = v
	return e
}

func (e *htmlDialog) Title(v string) *htmlDialog {
	e.a["title"] = v
	return e
}

func (e *htmlDialog) Lang(v string) *htmlDialog {
	e.a["lang"] = v
	return e
}

func (e *htmlDialog) Translate(v bool) *htmlDialog {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDialog) Dir(v string) *htmlDialog {
	e.a["dir"] = v
	return e
}

func (e *htmlDialog) Hidden(v bool) *htmlDialog {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDialog) TabIndex(v int) *htmlDialog {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDialog) AccessKey(v string) *htmlDialog {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDialog) Draggable(v bool) *htmlDialog {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDialog) Spellcheck(v bool) *htmlDialog {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
