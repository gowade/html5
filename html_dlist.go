package h

type htmlDList struct {
	htmlElement
}

func DList() *htmlDList {
	e := &htmlDList{}
	e.a = make(map[string]interface{})
	e.tagName = "dlist"
	return e
}

func (e *htmlDList) S(style StyleMap) *htmlDList {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDList) Key(key interface{}) *htmlDList {
	e.key = F(key)
	return e
}

func (e *htmlDList) ID(v string) *htmlDList {
	e.a["id"] = v
	return e
}

func (e *htmlDList) Class(v string) *htmlDList {
	e.a["class"] = v
	return e
}

func (e *htmlDList) Title(v string) *htmlDList {
	e.a["title"] = v
	return e
}

func (e *htmlDList) Lang(v string) *htmlDList {
	e.a["lang"] = v
	return e
}

func (e *htmlDList) Translate(v bool) *htmlDList {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDList) Dir(v string) *htmlDList {
	e.a["dir"] = v
	return e
}

func (e *htmlDList) Hidden(v bool) *htmlDList {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDList) TabIndex(v int) *htmlDList {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDList) AccessKey(v string) *htmlDList {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDList) Draggable(v bool) *htmlDList {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDList) Spellcheck(v bool) *htmlDList {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
