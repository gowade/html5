package h

func (e *htmlElement) ID(v string) *htmlElement {
	e.a["id"] = v
	return e
}

func (e *htmlElement) Class(v string) *htmlElement {
	e.a["class"] = v
	return e
}

func (e *htmlElement) Title(v string) *htmlElement {
	e.a["title"] = v
	return e
}

func (e *htmlElement) Lang(v string) *htmlElement {
	e.a["lang"] = v
	return e
}

func (e *htmlElement) Translate(v bool) *htmlElement {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlElement) Dir(v string) *htmlElement {
	e.a["dir"] = v
	return e
}

func (e *htmlElement) Hidden(v bool) *htmlElement {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlElement) TabIndex(v int) *htmlElement {
	e.a["tabindex"] = v
	return e
}

func (e *htmlElement) AccessKey(v string) *htmlElement {
	e.a["accesskey"] = v
	return e
}

func (e *htmlElement) Draggable(v bool) *htmlElement {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlElement) Spellcheck(v bool) *htmlElement {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
