package h

// HTMLCanvas represents HTML <canvas> tag
type HTMLCanvas struct {
	HTMLElement
}

// Canvas creates a HTML <canvas> tag
func Canvas() *HTMLCanvas {
	e := &HTMLCanvas{}
	e.a = make(map[string]interface{})
	e.tagName = "canvas"
	return e
}

// S sets the element's CSS properties
func (e *HTMLCanvas) S(style StyleMap) *HTMLCanvas {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLCanvas) Key(key interface{}) *HTMLCanvas {
	e.key = F(key)
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLCanvas) Width(v int) *HTMLCanvas {
	e.a["width"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLCanvas) Height(v int) *HTMLCanvas {
	e.a["height"] = v
	return e
}

// ProbablySupportsContext sets the element's "probablysupportscontext" attribute
func (e *HTMLCanvas) ProbablySupportsContext(v bool) *HTMLCanvas {
	if v {
		e.a["probablysupportscontext"] = ""
	} else {
		delete(e.a, "probablysupportscontext")
	}
	return e
}

// ToDataURL sets the element's "todataurl" attribute
func (e *HTMLCanvas) ToDataURL(v string) *HTMLCanvas {
	e.a["todataurl"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLCanvas) ID(v string) *HTMLCanvas {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLCanvas) Class(v string) *HTMLCanvas {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLCanvas) Title(v string) *HTMLCanvas {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLCanvas) Lang(v string) *HTMLCanvas {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLCanvas) Translate(v bool) *HTMLCanvas {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLCanvas) Dir(v string) *HTMLCanvas {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLCanvas) Hidden(v bool) *HTMLCanvas {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLCanvas) TabIndex(v int) *HTMLCanvas {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLCanvas) AccessKey(v string) *HTMLCanvas {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLCanvas) Draggable(v bool) *HTMLCanvas {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLCanvas) Spellcheck(v bool) *HTMLCanvas {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
