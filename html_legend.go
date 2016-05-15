package h

// HTMLLegend represents HTML <legend> tag
type HTMLLegend struct {
	HTMLElement
}

// Legend creates a HTML <legend> tag
func Legend() *HTMLLegend {
	e := &HTMLLegend{}
	e.a = make(map[string]interface{})
	e.tagName = "legend"
	return e
}

// S sets the element's CSS properties
func (e *HTMLLegend) S(style StyleMap) *HTMLLegend {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLLegend) Key(key interface{}) *HTMLLegend {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLLegend) ID(v string) *HTMLLegend {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLLegend) Class(v string) *HTMLLegend {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLLegend) Title(v string) *HTMLLegend {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLLegend) Lang(v string) *HTMLLegend {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLLegend) Translate(v bool) *HTMLLegend {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLLegend) Dir(v string) *HTMLLegend {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLLegend) Hidden(v bool) *HTMLLegend {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLLegend) TabIndex(v int) *HTMLLegend {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLLegend) AccessKey(v string) *HTMLLegend {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLLegend) Draggable(v bool) *HTMLLegend {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLLegend) Spellcheck(v bool) *HTMLLegend {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
