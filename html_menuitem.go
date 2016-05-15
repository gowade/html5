package h

// HTMLMenuItem represents HTML <menuitem> tag
type HTMLMenuItem struct {
	HTMLElement
}

// MenuItem creates a HTML <menuitem> tag
func MenuItem() *HTMLMenuItem {
	e := &HTMLMenuItem{}
	e.a = make(map[string]interface{})
	e.tagName = "menuitem"
	return e
}

// S sets the element's CSS properties
func (e *HTMLMenuItem) S(style StyleMap) *HTMLMenuItem {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLMenuItem) Key(key interface{}) *HTMLMenuItem {
	e.key = F(key)
	return e
}

// Type sets the element's "type" attribute
func (e *HTMLMenuItem) Type(v string) *HTMLMenuItem {
	e.a["type"] = v
	return e
}

// Label sets the element's "label" attribute
func (e *HTMLMenuItem) Label(v string) *HTMLMenuItem {
	e.a["label"] = v
	return e
}

// Icon sets the element's "icon" attribute
func (e *HTMLMenuItem) Icon(v string) *HTMLMenuItem {
	e.a["icon"] = v
	return e
}

// Disabled sets the element's "disabled" attribute
func (e *HTMLMenuItem) Disabled(v bool) *HTMLMenuItem {
	if v {
		e.a["disabled"] = ""
	} else {
		delete(e.a, "disabled")
	}
	return e
}

// Checked sets the element's "checked" attribute
func (e *HTMLMenuItem) Checked(v bool) *HTMLMenuItem {
	if v {
		e.a["checked"] = ""
	} else {
		delete(e.a, "checked")
	}
	return e
}

// Radiogroup sets the element's "radiogroup" attribute
func (e *HTMLMenuItem) Radiogroup(v string) *HTMLMenuItem {
	e.a["radiogroup"] = v
	return e
}

// Default sets the element's "default" attribute
func (e *HTMLMenuItem) Default(v bool) *HTMLMenuItem {
	if v {
		e.a["default"] = ""
	} else {
		delete(e.a, "default")
	}
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLMenuItem) ID(v string) *HTMLMenuItem {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLMenuItem) Class(v string) *HTMLMenuItem {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLMenuItem) Title(v string) *HTMLMenuItem {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLMenuItem) Lang(v string) *HTMLMenuItem {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLMenuItem) Translate(v bool) *HTMLMenuItem {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLMenuItem) Dir(v string) *HTMLMenuItem {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLMenuItem) Hidden(v bool) *HTMLMenuItem {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLMenuItem) TabIndex(v int) *HTMLMenuItem {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLMenuItem) AccessKey(v string) *HTMLMenuItem {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLMenuItem) Draggable(v bool) *HTMLMenuItem {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLMenuItem) Spellcheck(v bool) *HTMLMenuItem {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
