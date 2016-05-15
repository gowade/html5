package h

// HTMLOption represents HTML <option> tag
type HTMLOption struct {
	HTMLElement
}

// Option creates a HTML <option> tag
func Option() *HTMLOption {
	e := &HTMLOption{}
	e.a = make(map[string]interface{})
	e.tagName = "option"
	return e
}

// S sets the element's CSS properties
func (e *HTMLOption) S(style StyleMap) *HTMLOption {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLOption) Key(key interface{}) *HTMLOption {
	e.key = F(key)
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLOption) Disabled(v bool) *HTMLOption {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Label sets the element's "label" attribute
func (e *HTMLOption) Label(v string) *HTMLOption {
	e.a["label"] = v
	return e
}

// DefaultSelected sets the element's "defaultselected" attribute
func (e *HTMLOption) DefaultSelected(v bool) *HTMLOption {
	if v {
		e.a["defaultselected"] = ""
	} else {
		delete(e.a, "defaultselected")
	}
	return e
}

// Selected sets the element's "selected" attribute
func (e *HTMLOption) Selected(v bool) *HTMLOption {
	if v {
		e.a["selected"] = ""
	} else {
		delete(e.a, "selected")
	}
	return e
}

// Value sets the element's "value" attribute
func (e *HTMLOption) Value(v string) *HTMLOption {
	e.a["value"] = v
	return e
}

// Text sets the element's "text" attribute
func (e *HTMLOption) Text(v string) *HTMLOption {
	e.a["text"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLOption) ID(v string) *HTMLOption {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLOption) Class(v string) *HTMLOption {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLOption) Title(v string) *HTMLOption {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLOption) Lang(v string) *HTMLOption {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLOption) Translate(v bool) *HTMLOption {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLOption) Dir(v string) *HTMLOption {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLOption) Hidden(v bool) *HTMLOption {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLOption) TabIndex(v int) *HTMLOption {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLOption) AccessKey(v string) *HTMLOption {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLOption) Draggable(v bool) *HTMLOption {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLOption) Spellcheck(v bool) *HTMLOption {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
