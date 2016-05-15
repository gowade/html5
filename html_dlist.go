package h

// HTMLDList represents HTML <dlist> tag
type HTMLDList struct {
	HTMLElement
}

// DList creates a HTML <dlist> tag
func DList() *HTMLDList {
	e := &HTMLDList{}
	e.a = make(map[string]interface{})
	e.tagName = "dlist"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDList) S(style StyleMap) *HTMLDList {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDList) Key(key interface{}) *HTMLDList {
	e.key = F(key)
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDList) ID(v string) *HTMLDList {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDList) Class(v string) *HTMLDList {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDList) Title(v string) *HTMLDList {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDList) Lang(v string) *HTMLDList {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDList) Translate(v bool) *HTMLDList {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDList) Dir(v string) *HTMLDList {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDList) Hidden(v bool) *HTMLDList {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDList) TabIndex(v int) *HTMLDList {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDList) AccessKey(v string) *HTMLDList {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDList) Draggable(v bool) *HTMLDList {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDList) Spellcheck(v bool) *HTMLDList {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
