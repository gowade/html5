package html5

// HTMLDataList represents HTML <datalist> tag
type HTMLDataList struct {
	HTMLElement
}

// DataList creates an HTML <datalist> tag element
func DataList() *HTMLDataList {
	e := &HTMLDataList{}
	e.a = make(map[string]interface{})
	e.tagName = "datalist"
	return e
}

// S sets the element's CSS properties
func (e *HTMLDataList) S(style StyleMap) *HTMLDataList {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLDataList) Key(key interface{}) *HTMLDataList {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLDataList) Ref(dest *DOMElement) *HTMLDataList {
	e.ref = dest
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLDataList) ID(v string) *HTMLDataList {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLDataList) Class(v string) *HTMLDataList {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLDataList) Title(v string) *HTMLDataList {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLDataList) Lang(v string) *HTMLDataList {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLDataList) Translate(v bool) *HTMLDataList {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLDataList) Dir(v string) *HTMLDataList {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLDataList) Hidden(v bool) *HTMLDataList {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLDataList) TabIndex(v int) *HTMLDataList {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLDataList) AccessKey(v string) *HTMLDataList {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLDataList) Draggable(v bool) *HTMLDataList {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLDataList) Spellcheck(v bool) *HTMLDataList {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
