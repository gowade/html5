package h

type htmlBase struct {
	htmlElement
}

func Base() *htmlBase {
	e := &htmlBase{}
	e.a = make(map[string]interface{})
	e.tagName = "base"
	return e
}

func (e *htmlBase) S(style StyleMap) *htmlBase {
	e.htmlElement.S(style)
	return e
}

func (e *htmlBase) Key(key interface{}) *htmlBase {
	e.key = F(key)
	return e
}

func (e *htmlBase) Href(v string) *htmlBase {
	e.a["href"] = v
	return e
}

func (e *htmlBase) Target(v string) *htmlBase {
	e.a["target"] = v
	return e
}

func (e *htmlBase) ID(v string) *htmlBase {
	e.a["id"] = v
	return e
}

func (e *htmlBase) Class(v string) *htmlBase {
	e.a["class"] = v
	return e
}

func (e *htmlBase) Title(v string) *htmlBase {
	e.a["title"] = v
	return e
}

func (e *htmlBase) Lang(v string) *htmlBase {
	e.a["lang"] = v
	return e
}

func (e *htmlBase) Translate(v bool) *htmlBase {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlBase) Dir(v string) *htmlBase {
	e.a["dir"] = v
	return e
}

func (e *htmlBase) Hidden(v bool) *htmlBase {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlBase) TabIndex(v int) *htmlBase {
	e.a["tabindex"] = v
	return e
}

func (e *htmlBase) AccessKey(v string) *htmlBase {
	e.a["accesskey"] = v
	return e
}

func (e *htmlBase) Draggable(v bool) *htmlBase {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlBase) Spellcheck(v bool) *htmlBase {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
