package h

type htmlParam struct {
	htmlElement
}

func Param() *htmlParam {
	e := &htmlParam{}
	e.a = make(map[string]interface{})
	e.tagName = "param"
	return e
}

func (e *htmlParam) S(style StyleMap) *htmlParam {
	e.htmlElement.S(style)
	return e
}

func (e *htmlParam) Key(key interface{}) *htmlParam {
	e.key = F(key)
	return e
}

func (e *htmlParam) Name(v string) *htmlParam {
	e.a["name"] = v
	return e
}

func (e *htmlParam) Value(v string) *htmlParam {
	e.a["value"] = v
	return e
}

func (e *htmlParam) ID(v string) *htmlParam {
	e.a["id"] = v
	return e
}

func (e *htmlParam) Class(v string) *htmlParam {
	e.a["class"] = v
	return e
}

func (e *htmlParam) Title(v string) *htmlParam {
	e.a["title"] = v
	return e
}

func (e *htmlParam) Lang(v string) *htmlParam {
	e.a["lang"] = v
	return e
}

func (e *htmlParam) Translate(v bool) *htmlParam {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlParam) Dir(v string) *htmlParam {
	e.a["dir"] = v
	return e
}

func (e *htmlParam) Hidden(v bool) *htmlParam {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlParam) TabIndex(v int) *htmlParam {
	e.a["tabindex"] = v
	return e
}

func (e *htmlParam) AccessKey(v string) *htmlParam {
	e.a["accesskey"] = v
	return e
}

func (e *htmlParam) Draggable(v bool) *htmlParam {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlParam) Spellcheck(v bool) *htmlParam {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
