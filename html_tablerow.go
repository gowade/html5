package h

type htmlTableRow struct {
	htmlElement
}

func TableRow() *htmlTableRow {
	e := &htmlTableRow{}
	e.a = make(map[string]interface{})
	e.tagName = "tablerow"
	return e
}

func (e *htmlTableRow) S(style StyleMap) *htmlTableRow {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTableRow) Key(key interface{}) *htmlTableRow {
	e.key = F(key)
	return e
}

func (e *htmlTableRow) ID(v string) *htmlTableRow {
	e.a["id"] = v
	return e
}

func (e *htmlTableRow) Class(v string) *htmlTableRow {
	e.a["class"] = v
	return e
}

func (e *htmlTableRow) Title(v string) *htmlTableRow {
	e.a["title"] = v
	return e
}

func (e *htmlTableRow) Lang(v string) *htmlTableRow {
	e.a["lang"] = v
	return e
}

func (e *htmlTableRow) Translate(v bool) *htmlTableRow {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTableRow) Dir(v string) *htmlTableRow {
	e.a["dir"] = v
	return e
}

func (e *htmlTableRow) Hidden(v bool) *htmlTableRow {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTableRow) TabIndex(v int) *htmlTableRow {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTableRow) AccessKey(v string) *htmlTableRow {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTableRow) Draggable(v bool) *htmlTableRow {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTableRow) Spellcheck(v bool) *htmlTableRow {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
