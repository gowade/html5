package h

type htmlMarquee struct {
	htmlElement
}

func Marquee() *htmlMarquee {
	e := &htmlMarquee{}
	e.a = make(map[string]interface{})
	e.tagName = "marquee"
	return e
}

func (e *htmlMarquee) S(style StyleMap) *htmlMarquee {
	e.htmlElement.S(style)
	return e
}

func (e *htmlMarquee) Key(key interface{}) *htmlMarquee {
	e.key = F(key)
	return e
}

func (e *htmlMarquee) Behavior(v string) *htmlMarquee {
	e.a["behavior"] = v
	return e
}

func (e *htmlMarquee) BgColor(v string) *htmlMarquee {
	e.a["bgcolor"] = v
	return e
}

func (e *htmlMarquee) Direction(v string) *htmlMarquee {
	e.a["direction"] = v
	return e
}

func (e *htmlMarquee) Height(v string) *htmlMarquee {
	e.a["height"] = v
	return e
}

func (e *htmlMarquee) Hspace(v int) *htmlMarquee {
	e.a["hspace"] = v
	return e
}

func (e *htmlMarquee) Loop(v int) *htmlMarquee {
	e.a["loop"] = v
	return e
}

func (e *htmlMarquee) ScrollAmount(v int) *htmlMarquee {
	e.a["scrollamount"] = v
	return e
}

func (e *htmlMarquee) ScrollDelay(v int) *htmlMarquee {
	e.a["scrolldelay"] = v
	return e
}

func (e *htmlMarquee) TrueSpeed(v bool) *htmlMarquee {
	if v {
		e.a["truespeed"] = ""
	} else {
		delete(e.a, "truespeed")
	}
	return e
}

func (e *htmlMarquee) Vspace(v int) *htmlMarquee {
	e.a["vspace"] = v
	return e
}

func (e *htmlMarquee) Width(v string) *htmlMarquee {
	e.a["width"] = v
	return e
}

func (e *htmlMarquee) ID(v string) *htmlMarquee {
	e.a["id"] = v
	return e
}

func (e *htmlMarquee) Class(v string) *htmlMarquee {
	e.a["class"] = v
	return e
}

func (e *htmlMarquee) Title(v string) *htmlMarquee {
	e.a["title"] = v
	return e
}

func (e *htmlMarquee) Lang(v string) *htmlMarquee {
	e.a["lang"] = v
	return e
}

func (e *htmlMarquee) Translate(v bool) *htmlMarquee {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

func (e *htmlMarquee) Dir(v string) *htmlMarquee {
	e.a["dir"] = v
	return e
}

func (e *htmlMarquee) Hidden(v bool) *htmlMarquee {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

func (e *htmlMarquee) TabIndex(v int) *htmlMarquee {
	e.a["tabindex"] = v
	return e
}

func (e *htmlMarquee) AccessKey(v string) *htmlMarquee {
	e.a["accesskey"] = v
	return e
}

func (e *htmlMarquee) Draggable(v bool) *htmlMarquee {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

func (e *htmlMarquee) Spellcheck(v bool) *htmlMarquee {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
