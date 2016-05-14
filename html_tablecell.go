package h

type htmlTableCell struct {
	htmlElement
}

func TableCell() *htmlTableCell {
	e := &htmlTableCell{}
	e.a = make(map[string]interface{})
	e.tagName = "tablecell"
	return e
}

func (e *htmlTableCell) S(style StyleMap) *htmlTableCell {
	e.htmlElement.S(style)
	return e
}

func (e *htmlTableCell) Key(key interface{}) *htmlTableCell {
	e.key = F(key)
	return e
}

func (e *htmlTableCell) ColSpan(v int) *htmlTableCell {
	e.a["colspan"] = v
	return e
}

func (e *htmlTableCell) RowSpan(v int) *htmlTableCell {
	e.a["rowspan"] = v
	return e
}

func (e *htmlTableCell) Scope(v string) *htmlTableCell {
	e.a["scope"] = v
	return e
}

func (e *htmlTableCell) Abbr(v string) *htmlTableCell {
	e.a["abbr"] = v
	return e
}

func (e *htmlTableCell) ID(v string) *htmlTableCell {
	e.a["id"] = v
	return e
}

func (e *htmlTableCell) Class(v string) *htmlTableCell {
	e.a["class"] = v
	return e
}

func (e *htmlTableCell) Title(v string) *htmlTableCell {
	e.a["title"] = v
	return e
}

func (e *htmlTableCell) Lang(v string) *htmlTableCell {
	e.a["lang"] = v
	return e
}

func (e *htmlTableCell) Translate(v bool) *htmlTableCell {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlTableCell) Dir(v string) *htmlTableCell {
	e.a["dir"] = v
	return e
}

func (e *htmlTableCell) Hidden(v bool) *htmlTableCell {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlTableCell) TabIndex(v int) *htmlTableCell {
	e.a["tabindex"] = v
	return e
}

func (e *htmlTableCell) AccessKey(v string) *htmlTableCell {
	e.a["accesskey"] = v
	return e
}

func (e *htmlTableCell) Draggable(v bool) *htmlTableCell {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlTableCell) Spellcheck(v bool) *htmlTableCell {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
