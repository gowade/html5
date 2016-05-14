package h

type htmlMod struct {
	htmlElement
}

func Mod() *htmlMod {
	e := &htmlMod{}
	e.a = make(map[string]interface{})
	e.tagName = "mod"
	return e
}

func (e *htmlMod) S(style StyleMap) *htmlMod {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMod) Key(key interface{}) *htmlMod {
	e.key = F(key)
	return e
}

func (e *htmlMod) Cite(v string) *htmlMod {
	e.a["cite"] = v
	return e
}

func (e *htmlMod) DateTime(v string) *htmlMod {
	e.a["datetime"] = v
	return e
}

func (e *htmlMod) ID(v string) *htmlMod {
	e.a["id"] = v
	return e
}

func (e *htmlMod) Class(v string) *htmlMod {
	e.a["class"] = v
	return e
}

func (e *htmlMod) Title(v string) *htmlMod {
	e.a["title"] = v
	return e
}

func (e *htmlMod) Lang(v string) *htmlMod {
	e.a["lang"] = v
	return e
}

func (e *htmlMod) Translate(v bool) *htmlMod {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMod) Dir(v string) *htmlMod {
	e.a["dir"] = v
	return e
}

func (e *htmlMod) Hidden(v bool) *htmlMod {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMod) TabIndex(v int) *htmlMod {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMod) AccessKey(v string) *htmlMod {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMod) Draggable(v bool) *htmlMod {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMod) Spellcheck(v bool) *htmlMod {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
