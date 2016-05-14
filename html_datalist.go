package h

type htmlDataList struct {
	htmlElement
}

func DataList() *htmlDataList {
	e := &htmlDataList{}
	e.a = make(map[string]interface{})
	e.tagName = "datalist"
	return e
}

func (e *htmlDataList) S(style StyleMap) *htmlDataList {
	e.htmlElement.S(style)
	return e
}

func (e *htmlDataList) Key(key interface{}) *htmlDataList {
	e.key = F(key)
	return e
}

func (e *htmlDataList) ID(v string) *htmlDataList {
	e.a["id"] = v
	return e
}

func (e *htmlDataList) Class(v string) *htmlDataList {
	e.a["class"] = v
	return e
}

func (e *htmlDataList) Title(v string) *htmlDataList {
	e.a["title"] = v
	return e
}

func (e *htmlDataList) Lang(v string) *htmlDataList {
	e.a["lang"] = v
	return e
}

func (e *htmlDataList) Translate(v bool) *htmlDataList {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlDataList) Dir(v string) *htmlDataList {
	e.a["dir"] = v
	return e
}

func (e *htmlDataList) Hidden(v bool) *htmlDataList {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlDataList) TabIndex(v int) *htmlDataList {
	e.a["tabindex"] = v
	return e
}

func (e *htmlDataList) AccessKey(v string) *htmlDataList {
	e.a["accesskey"] = v
	return e
}

func (e *htmlDataList) Draggable(v bool) *htmlDataList {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlDataList) Spellcheck(v bool) *htmlDataList {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
