package h

// ID sets the element's "id" attribute
func (e *HTMLElement) ID(v string) *HTMLElement {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLElement) Class(v string) *HTMLElement {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLElement) Title(v string) *HTMLElement {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLElement) Lang(v string) *HTMLElement {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLElement) Translate(v bool) *HTMLElement {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLElement) Dir(v string) *HTMLElement {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLElement) Hidden(v bool) *HTMLElement {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLElement) TabIndex(v int) *HTMLElement {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLElement) AccessKey(v string) *HTMLElement {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLElement) Draggable(v bool) *HTMLElement {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLElement) Spellcheck(v bool) *HTMLElement {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
