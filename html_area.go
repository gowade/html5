package h

type htmlArea struct {
	htmlElement
}

func Area() *htmlArea {
	e := &htmlArea{}
	e.a = make(map[string]interface{})
	e.tagName = "area"
	return e
}

func (e *htmlArea) S(style StyleMap) *htmlArea {
	e.htmlElement.S(style)
	return e
}

func (e *htmlArea) Key(key interface{}) *htmlArea {
	e.key = F(key)
	return e
}

func (e *htmlArea) Alt(v string) *htmlArea {
	e.a["alt"] = v
	return e
}

func (e *htmlArea) Coords(v string) *htmlArea {
	e.a["coords"] = v
	return e
}

func (e *htmlArea) Shape(v string) *htmlArea {
	e.a["shape"] = v
	return e
}

func (e *htmlArea) Target(v string) *htmlArea {
	e.a["target"] = v
	return e
}

func (e *htmlArea) Download(v string) *htmlArea {
	e.a["download"] = v
	return e
}

func (e *htmlArea) Rel(v string) *htmlArea {
	e.a["rel"] = v
	return e
}

func (e *htmlArea) ReferrerPolicy(v string) *htmlArea {
	e.a["referrerpolicy"] = v
	return e
}

func (e *htmlArea) ID(v string) *htmlArea {
	e.a["id"] = v
	return e
}

func (e *htmlArea) Class(v string) *htmlArea {
	e.a["class"] = v
	return e
}

func (e *htmlArea) Title(v string) *htmlArea {
	e.a["title"] = v
	return e
}

func (e *htmlArea) Lang(v string) *htmlArea {
	e.a["lang"] = v
	return e
}

func (e *htmlArea) Translate(v bool) *htmlArea {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlArea) Dir(v string) *htmlArea {
	e.a["dir"] = v
	return e
}

func (e *htmlArea) Hidden(v bool) *htmlArea {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlArea) TabIndex(v int) *htmlArea {
	e.a["tabindex"] = v
	return e
}

func (e *htmlArea) AccessKey(v string) *htmlArea {
	e.a["accesskey"] = v
	return e
}

func (e *htmlArea) Draggable(v bool) *htmlArea {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlArea) Spellcheck(v bool) *htmlArea {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
