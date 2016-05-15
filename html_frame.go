package html5

// HTMLFrame represents HTML <frame> tag
type HTMLFrame struct {
	HTMLElement
}

// Frame creates an HTML <frame> tag element
func Frame() *HTMLFrame {
	e := &HTMLFrame{}
	e.a = make(map[string]interface{})
	e.tagName = "frame"
	return e
}

// S sets the element's CSS properties
func (e *HTMLFrame) S(style StyleMap) *HTMLFrame {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLFrame) Key(key interface{}) *HTMLFrame {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLFrame) Ref(dest *DOMElement) *HTMLFrame {
	e.ref = dest
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLFrame) Name(v string) *HTMLFrame {
	e.a["name"] = v
	return e
}

// Scrolling sets the element's "scrolling" attribute
func (e *HTMLFrame) Scrolling(v string) *HTMLFrame {
	e.a["scrolling"] = v
	return e
}

// Src sets the element's "src" attribute
func (e *HTMLFrame) Src(v string) *HTMLFrame {
	e.a["src"] = v
	return e
}

// FrameBorder sets the element's "frameborder" attribute
func (e *HTMLFrame) FrameBorder(v string) *HTMLFrame {
	e.a["frameborder"] = v
	return e
}

// LongDesc sets the element's "longdesc" attribute
func (e *HTMLFrame) LongDesc(v string) *HTMLFrame {
	e.a["longdesc"] = v
	return e
}

// NoResize sets the element's "noresize" attribute
func (e *HTMLFrame) NoResize(v bool) *HTMLFrame {
	if v {
		e.a["noresize"] = ""
	} else {
		delete(e.a, "noresize")
	}
	return e
}

// MarginHeight sets the element's "marginheight" attribute
func (e *HTMLFrame) MarginHeight(v string) *HTMLFrame {
	e.a["marginheight"] = v
	return e
}

// MarginWidth sets the element's "marginwidth" attribute
func (e *HTMLFrame) MarginWidth(v string) *HTMLFrame {
	e.a["marginwidth"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLFrame) ID(v string) *HTMLFrame {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLFrame) Class(v string) *HTMLFrame {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLFrame) Title(v string) *HTMLFrame {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLFrame) Lang(v string) *HTMLFrame {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLFrame) Translate(v bool) *HTMLFrame {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLFrame) Dir(v string) *HTMLFrame {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLFrame) Hidden(v bool) *HTMLFrame {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLFrame) TabIndex(v int) *HTMLFrame {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLFrame) AccessKey(v string) *HTMLFrame {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLFrame) Draggable(v bool) *HTMLFrame {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLFrame) Spellcheck(v bool) *HTMLFrame {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
