package h

// HTMLMarquee represents HTML <marquee> tag
type HTMLMarquee struct {
	HTMLElement
}

// Marquee creates an HTML <marquee> tag element
func Marquee() *HTMLMarquee {
	e := &HTMLMarquee{}
	e.a = make(map[string]interface{})
	e.tagName = "marquee"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMarquee) S(style StyleMap) *HTMLMarquee {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMarquee) Key(key interface{}) *HTMLMarquee {
	e.key = F(key)
	return e
}

// Behavior sets the element's "behavior" attribute
func (e *HTMLMarquee) Behavior(v string) *HTMLMarquee {
	e.a["behavior"] = v
	return e
}

// BgColor sets the element's "bgcolor" attribute
func (e *HTMLMarquee) BgColor(v string) *HTMLMarquee {
	e.a["bgcolor"] = v
	return e
}

// Direction sets the element's "direction" attribute
func (e *HTMLMarquee) Direction(v string) *HTMLMarquee {
	e.a["direction"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLMarquee) Height(v string) *HTMLMarquee {
	e.a["height"] = v
	return e
}

// Hspace sets the element's "hspace" attribute
func (e *HTMLMarquee) Hspace(v int) *HTMLMarquee {
	e.a["hspace"] = v
	return e
}

// Loop sets the element's "loop" attribute
func (e *HTMLMarquee) Loop(v int) *HTMLMarquee {
	e.a["loop"] = v
	return e
}

// ScrollAmount sets the element's "scrollamount" attribute
func (e *HTMLMarquee) ScrollAmount(v int) *HTMLMarquee {
	e.a["scrollamount"] = v
	return e
}

// ScrollDelay sets the element's "scrolldelay" attribute
func (e *HTMLMarquee) ScrollDelay(v int) *HTMLMarquee {
	e.a["scrolldelay"] = v
	return e
}

// TrueSpeed sets the element's "truespeed" attribute
func (e *HTMLMarquee) TrueSpeed(v bool) *HTMLMarquee {
	if v {
		e.a["truespeed"] = ""
	} else {
		delete(e.a, "truespeed")
	}
	return e
}

// Vspace sets the element's "vspace" attribute
func (e *HTMLMarquee) Vspace(v int) *HTMLMarquee {
	e.a["vspace"] = v
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLMarquee) Width(v string) *HTMLMarquee {
	e.a["width"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMarquee) ID(v string) *HTMLMarquee {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMarquee) Class(v string) *HTMLMarquee {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMarquee) Title(v string) *HTMLMarquee {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMarquee) Lang(v string) *HTMLMarquee {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMarquee) Translate(v bool) *HTMLMarquee {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMarquee) Dir(v string) *HTMLMarquee {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMarquee) Hidden(v bool) *HTMLMarquee {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMarquee) TabIndex(v int) *HTMLMarquee {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMarquee) AccessKey(v string) *HTMLMarquee {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMarquee) Draggable(v bool) *HTMLMarquee {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMarquee) Spellcheck(v bool) *HTMLMarquee {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
