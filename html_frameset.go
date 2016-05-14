package h

type htmlFrameSet struct {
	htmlElement
}

func FrameSet() *htmlFrameSet {
	e := &htmlFrameSet{}
	e.a = make(map[string]interface{})
	e.tagName = "frameset"
	return e
}

func (e *htmlFrameSet) S(style StyleMap) *htmlFrameSet {
	e.htmlElement.S(style)
	return e
}

func (e *htmlFrameSet) Key(key interface{}) *htmlFrameSet {
	e.key = F(key)
	return e
}

func (e *htmlFrameSet) Cols(v string) *htmlFrameSet {
	e.a["cols"] = v
	return e
}

func (e *htmlFrameSet) Rows(v string) *htmlFrameSet {
	e.a["rows"] = v
	return e
}

func (e *htmlFrameSet) ID(v string) *htmlFrameSet {
	e.a["id"] = v
	return e
}

func (e *htmlFrameSet) Class(v string) *htmlFrameSet {
	e.a["class"] = v
	return e
}

func (e *htmlFrameSet) Title(v string) *htmlFrameSet {
	e.a["title"] = v
	return e
}

func (e *htmlFrameSet) Lang(v string) *htmlFrameSet {
	e.a["lang"] = v
	return e
}

func (e *htmlFrameSet) Translate(v bool) *htmlFrameSet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlFrameSet) Dir(v string) *htmlFrameSet {
	e.a["dir"] = v
	return e
}

func (e *htmlFrameSet) Hidden(v bool) *htmlFrameSet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlFrameSet) TabIndex(v int) *htmlFrameSet {
	e.a["tabindex"] = v
	return e
}

func (e *htmlFrameSet) AccessKey(v string) *htmlFrameSet {
	e.a["accesskey"] = v
	return e
}

func (e *htmlFrameSet) Draggable(v bool) *htmlFrameSet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlFrameSet) Spellcheck(v bool) *htmlFrameSet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
