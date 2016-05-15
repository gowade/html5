package h

// HTMLSlot represents HTML <slot> tag
type HTMLSlot struct {
	HTMLElement
}

// Slot creates an HTML <slot> tag element
func Slot() *HTMLSlot {
	e := &HTMLSlot{}
	e.a = make(map[string]interface{})
	e.tagName = "slot"
	return e
}

// S sets the element's CSS properties
func (e *HTMLSlot) S(style StyleMap) *HTMLSlot {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLSlot) Key(key interface{}) *HTMLSlot {
	e.key = F(key)
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLSlot) Name(v string) *HTMLSlot {
	e.a["name"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLSlot) ID(v string) *HTMLSlot {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLSlot) Class(v string) *HTMLSlot {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLSlot) Title(v string) *HTMLSlot {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLSlot) Lang(v string) *HTMLSlot {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLSlot) Translate(v bool) *HTMLSlot {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLSlot) Dir(v string) *HTMLSlot {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLSlot) Hidden(v bool) *HTMLSlot {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLSlot) TabIndex(v int) *HTMLSlot {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLSlot) AccessKey(v string) *HTMLSlot {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLSlot) Draggable(v bool) *HTMLSlot {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLSlot) Spellcheck(v bool) *HTMLSlot {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
