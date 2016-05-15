package h

// HTMLQuote represents HTML <quote> tag
type HTMLQuote struct {
	HTMLElement
}

// Quote creates an HTML <quote> tag element
func Quote() *HTMLQuote {
	e := &HTMLQuote{}
	e.a = make(map[string]interface{})
	e.tagName = "quote"
	return e
}

// S sets the element's CSS properties
func (e *HTMLQuote) S(style StyleMap) *HTMLQuote {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLQuote) Key(key interface{}) *HTMLQuote {
	e.key = F(key)
	return e
}

// Cite sets the element's "cite" attribute
func (e *HTMLQuote) Cite(v string) *HTMLQuote {
	e.a["cite"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLQuote) ID(v string) *HTMLQuote {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLQuote) Class(v string) *HTMLQuote {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLQuote) Title(v string) *HTMLQuote {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLQuote) Lang(v string) *HTMLQuote {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLQuote) Translate(v bool) *HTMLQuote {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLQuote) Dir(v string) *HTMLQuote {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLQuote) Hidden(v bool) *HTMLQuote {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLQuote) TabIndex(v int) *HTMLQuote {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLQuote) AccessKey(v string) *HTMLQuote {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLQuote) Draggable(v bool) *HTMLQuote {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLQuote) Spellcheck(v bool) *HTMLQuote {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
