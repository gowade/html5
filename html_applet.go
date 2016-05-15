package h

// HTMLApplet represents HTML <applet> tag
type HTMLApplet struct {
	HTMLElement
}

// Applet creates a HTML <applet> tag
func Applet() *HTMLApplet {
	e := &HTMLApplet{}
	e.a = make(map[string]interface{})
	e.tagName = "applet"
	return e
}

// S sets the element's CSS properties
func (e *HTMLApplet) S(style StyleMap) *HTMLApplet {
	e.htmlElement.S(style)
	return
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLApplet) Key(key interface{}) *HTMLApplet {
	e.key = F(key)
	return e
}

// Align sets the element's "align" attribute
func (e *HTMLApplet) Align(v string) *HTMLApplet {
	e.a["align"] = v
	return e
}

// Alt sets the element's "alt" attribute
func (e *HTMLApplet) Alt(v string) *HTMLApplet {
	e.a["alt"] = v
	return e
}

// Archive sets the element's "archive" attribute
func (e *HTMLApplet) Archive(v string) *HTMLApplet {
	e.a["archive"] = v
	return e
}

// Code sets the element's "code" attribute
func (e *HTMLApplet) Code(v string) *HTMLApplet {
	e.a["code"] = v
	return e
}

// CodeBase sets the element's "codebase" attribute
func (e *HTMLApplet) CodeBase(v string) *HTMLApplet {
	e.a["codebase"] = v
	return e
}

// Height sets the element's "height" attribute
func (e *HTMLApplet) Height(v string) *HTMLApplet {
	e.a["height"] = v
	return e
}

// Hspace sets the element's "hspace" attribute
func (e *HTMLApplet) Hspace(v int) *HTMLApplet {
	e.a["hspace"] = v
	return e
}

// Name sets the element's "name" attribute
func (e *HTMLApplet) Name(v string) *HTMLApplet {
	e.a["name"] = v
	return e
}

// Object sets the element's "object" attribute
func (e *HTMLApplet) Object(v string) *HTMLApplet {
	e.a["object"] = v
	return e
}

// Vspace sets the element's "vspace" attribute
func (e *HTMLApplet) Vspace(v int) *HTMLApplet {
	e.a["vspace"] = v
	return e
}

// Width sets the element's "width" attribute
func (e *HTMLApplet) Width(v string) *HTMLApplet {
	e.a["width"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLApplet) ID(v string) *HTMLApplet {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLApplet) Class(v string) *HTMLApplet {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLApplet) Title(v string) *HTMLApplet {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLApplet) Lang(v string) *HTMLApplet {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLApplet) Translate(v bool) *HTMLApplet {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLApplet) Dir(v string) *HTMLApplet {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLApplet) Hidden(v bool) *HTMLApplet {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLApplet) TabIndex(v int) *HTMLApplet {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLApplet) AccessKey(v string) *HTMLApplet {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLApplet) Draggable(v bool) *HTMLApplet {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLApplet) Spellcheck(v bool) *HTMLApplet {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}
