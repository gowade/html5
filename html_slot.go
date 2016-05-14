package h

type htmlSlot struct {
	htmlElement
}

func Slot() *htmlSlot {
	e := &htmlSlot{}
	e.a = make(map[string]interface{})
	e.tagName = "slot"
	return e
}

func (e *htmlSlot) S(style StyleMap) *htmlSlot {
	e.htmlElement.S(style)
	return e
}

func (e *htmlSlot) Key(key interface{}) *htmlSlot {
	e.key = F(key)
	return e
}

func (e *htmlSlot) Name(v string) *htmlSlot {
	e.a["name"] = v
	return e
}

func (e *htmlSlot) ID(v string) *htmlSlot {
	e.a["id"] = v
	return e
}

func (e *htmlSlot) Class(v string) *htmlSlot {
	e.a["class"] = v
	return e
}

func (e *htmlSlot) Title(v string) *htmlSlot {
	e.a["title"] = v
	return e
}

func (e *htmlSlot) Lang(v string) *htmlSlot {
	e.a["lang"] = v
	return e
}

func (e *htmlSlot) Translate(v bool) *htmlSlot {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlSlot) Dir(v string) *htmlSlot {
	e.a["dir"] = v
	return e
}

func (e *htmlSlot) Hidden(v bool) *htmlSlot {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlSlot) TabIndex(v int) *htmlSlot {
	e.a["tabindex"] = v
	return e
}

func (e *htmlSlot) AccessKey(v string) *htmlSlot {
	e.a["accesskey"] = v
	return e
}

func (e *htmlSlot) Draggable(v bool) *htmlSlot {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlSlot) Spellcheck(v bool) *htmlSlot {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
