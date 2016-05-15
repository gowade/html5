package html5

// HTMLForm represents HTML <form> tag
type HTMLForm struct {
	HTMLElement
}

// Form creates an HTML <form> tag element
func Form() *HTMLForm {
	e := &HTMLForm{}
	e.a = make(map[string]interface{})
	e.tagName = "form"
	return e
}

// S sets the element's CSS properties
func (e *HTMLForm) S(style StyleMap) *HTMLForm {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLForm) Key(key interface{}) *HTMLForm {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLForm) Ref(dest *DOMElement) *HTMLForm {
	e.ref = dest
	return e
}

// AcceptCharset sets the element's "acceptcharset" attribute
func (e *HTMLForm) AcceptCharset(v string) *HTMLForm {
	e.a["acceptcharset"] = v
	return e
}

// Action sets the element's "action" attribute
func (e *HTMLForm) Action(v string) *HTMLForm {
	e.a["action"] = v
	return e
}

// Autocomplete sets the element's "autocomplete" attribute
func (e *HTMLForm) Autocomplete(v string) *HTMLForm {
	e.a["autocomplete"] = v
	return e
}

// Enctype sets the element's "enctype" attribute
func (e *HTMLForm) Enctype(v string) *HTMLForm {
	e.a["enctype"] = v
	return e
}

// Encoding sets the element's "encoding" attribute
func (e *HTMLForm) Encoding(v string) *HTMLForm {
	e.a["encoding"] = v
	return e
}

// Method sets the element's "method" attribute
func (e *HTMLForm) Method(v string) *HTMLForm {
	e.a["method"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLForm) Name(v string) *HTMLForm {
	e.a["name"] = v
	return e
}

// NoValidate sets the element's "novalidate" attribute
func (e *HTMLForm) NoValidate(v bool) *HTMLForm {
	if v {
		e.a["novalidate"] = ""
	} else {
		delete(e.a, "novalidate")
	}
	return e
}

// Target sets the element's "target" attribute
func (e *HTMLForm) Target(v string) *HTMLForm {
	e.a["target"] = v
	return e
}

// CheckValidity sets the element's "checkvalidity" attribute
func (e *HTMLForm) CheckValidity(v bool) *HTMLForm {
	if v {
		e.a["checkvalidity"] = ""
	} else {
		delete(e.a, "checkvalidity")
	}
	return e
}

// ReportValidity sets the element's "reportvalidity" attribute
func (e *HTMLForm) ReportValidity(v bool) *HTMLForm {
	if v {
		e.a["reportvalidity"] = ""
	} else {
		delete(e.a, "reportvalidity")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLForm) ID(v string) *HTMLForm {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLForm) Class(v string) *HTMLForm {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLForm) Title(v string) *HTMLForm {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLForm) Lang(v string) *HTMLForm {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLForm) Translate(v bool) *HTMLForm {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLForm) Dir(v string) *HTMLForm {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLForm) Hidden(v bool) *HTMLForm {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLForm) TabIndex(v int) *HTMLForm {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLForm) AccessKey(v string) *HTMLForm {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLForm) Draggable(v bool) *HTMLForm {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLForm) Spellcheck(v bool) *HTMLForm {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
