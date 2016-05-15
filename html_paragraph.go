package h

// HTMLParagraph represents HTML <paragraph> tag
type HTMLParagraph struct {
	HTMLElement
}

// Paragraph creates a HTML <paragraph> tag
func Paragraph() *HTMLParagraph {
	e := &HTMLParagraph{}
	e.a = make(map[string]interface{})
	e.tagName = "paragraph"
	return e
}

// S sets the element's CSS properties
func (e *HTMLParagraph) S(style StyleMap) *HTMLParagraph {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLParagraph) Key(key interface{}) *HTMLParagraph {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLParagraph) ID(v string) *HTMLParagraph {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLParagraph) Class(v string) *HTMLParagraph {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLParagraph) Title(v string) *HTMLParagraph {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLParagraph) Lang(v string) *HTMLParagraph {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLParagraph) Translate(v bool) *HTMLParagraph {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLParagraph) Dir(v string) *HTMLParagraph {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLParagraph) Hidden(v bool) *HTMLParagraph {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLParagraph) TabIndex(v int) *HTMLParagraph {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLParagraph) AccessKey(v string) *HTMLParagraph {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLParagraph) Draggable(v bool) *HTMLParagraph {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLParagraph) Spellcheck(v bool) *HTMLParagraph {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
