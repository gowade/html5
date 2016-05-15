package html5

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

// OnAbort sets the element's "onabort" attribute
func (e *HTMLElement) OnAbort(v Action) *HTMLElement {
	e.a["onabort"] = v
	return e
}

// OnBlur sets the element's "onblur" attribute
func (e *HTMLElement) OnBlur(v Action) *HTMLElement {
	e.a["onblur"] = v
	return e
}

// OnCancel sets the element's "oncancel" attribute
func (e *HTMLElement) OnCancel(v Action) *HTMLElement {
	e.a["oncancel"] = v
	return e
}

// OnCanPlay sets the element's "oncanplay" attribute
func (e *HTMLElement) OnCanPlay(v Action) *HTMLElement {
	e.a["oncanplay"] = v
	return e
}

// OnCanPlayThrough sets the element's "oncanplaythrough" attribute
func (e *HTMLElement) OnCanPlayThrough(v Action) *HTMLElement {
	e.a["oncanplaythrough"] = v
	return e
}

// OnChange sets the element's "onchange" attribute
func (e *HTMLElement) OnChange(v Action) *HTMLElement {
	e.a["onchange"] = v
	return e
}

// OnClick sets the element's "onclick" attribute
func (e *HTMLElement) OnClick(v Action) *HTMLElement {
	e.a["onclick"] = v
	return e
}

// OnClose sets the element's "onclose" attribute
func (e *HTMLElement) OnClose(v Action) *HTMLElement {
	e.a["onclose"] = v
	return e
}

// OnContextMenu sets the element's "oncontextmenu" attribute
func (e *HTMLElement) OnContextMenu(v Action) *HTMLElement {
	e.a["oncontextmenu"] = v
	return e
}

// OnCueChange sets the element's "oncuechange" attribute
func (e *HTMLElement) OnCueChange(v Action) *HTMLElement {
	e.a["oncuechange"] = v
	return e
}

// OnDblClick sets the element's "ondblclick" attribute
func (e *HTMLElement) OnDblClick(v Action) *HTMLElement {
	e.a["ondblclick"] = v
	return e
}

// OnDrag sets the element's "ondrag" attribute
func (e *HTMLElement) OnDrag(v Action) *HTMLElement {
	e.a["ondrag"] = v
	return e
}

// OnDragEnd sets the element's "ondragend" attribute
func (e *HTMLElement) OnDragEnd(v Action) *HTMLElement {
	e.a["ondragend"] = v
	return e
}

// OnDragEnter sets the element's "ondragenter" attribute
func (e *HTMLElement) OnDragEnter(v Action) *HTMLElement {
	e.a["ondragenter"] = v
	return e
}

// OnDragExit sets the element's "ondragexit" attribute
func (e *HTMLElement) OnDragExit(v Action) *HTMLElement {
	e.a["ondragexit"] = v
	return e
}

// OnDragLeave sets the element's "ondragleave" attribute
func (e *HTMLElement) OnDragLeave(v Action) *HTMLElement {
	e.a["ondragleave"] = v
	return e
}

// OnDragOver sets the element's "ondragover" attribute
func (e *HTMLElement) OnDragOver(v Action) *HTMLElement {
	e.a["ondragover"] = v
	return e
}

// OnDragStart sets the element's "ondragstart" attribute
func (e *HTMLElement) OnDragStart(v Action) *HTMLElement {
	e.a["ondragstart"] = v
	return e
}

// OnDrop sets the element's "ondrop" attribute
func (e *HTMLElement) OnDrop(v Action) *HTMLElement {
	e.a["ondrop"] = v
	return e
}

// OnDurationChange sets the element's "ondurationchange" attribute
func (e *HTMLElement) OnDurationChange(v Action) *HTMLElement {
	e.a["ondurationchange"] = v
	return e
}

// OnEmptied sets the element's "onemptied" attribute
func (e *HTMLElement) OnEmptied(v Action) *HTMLElement {
	e.a["onemptied"] = v
	return e
}

// OnEnded sets the element's "onended" attribute
func (e *HTMLElement) OnEnded(v Action) *HTMLElement {
	e.a["onended"] = v
	return e
}

// OnError sets the element's "onerror" attribute
func (e *HTMLElement) OnError(v Action) *HTMLElement {
	e.a["onerror"] = v
	return e
}

// OnFocus sets the element's "onfocus" attribute
func (e *HTMLElement) OnFocus(v Action) *HTMLElement {
	e.a["onfocus"] = v
	return e
}

// OnInput sets the element's "oninput" attribute
func (e *HTMLElement) OnInput(v Action) *HTMLElement {
	e.a["oninput"] = v
	return e
}

// OnInvalid sets the element's "oninvalid" attribute
func (e *HTMLElement) OnInvalid(v Action) *HTMLElement {
	e.a["oninvalid"] = v
	return e
}

// OnKeyDown sets the element's "onkeydown" attribute
func (e *HTMLElement) OnKeyDown(v Action) *HTMLElement {
	e.a["onkeydown"] = v
	return e
}

// OnKeyPress sets the element's "onkeypress" attribute
func (e *HTMLElement) OnKeyPress(v Action) *HTMLElement {
	e.a["onkeypress"] = v
	return e
}

// OnKeyUp sets the element's "onkeyup" attribute
func (e *HTMLElement) OnKeyUp(v Action) *HTMLElement {
	e.a["onkeyup"] = v
	return e
}

// OnLoad sets the element's "onload" attribute
func (e *HTMLElement) OnLoad(v Action) *HTMLElement {
	e.a["onload"] = v
	return e
}

// OnLoadedData sets the element's "onloadeddata" attribute
func (e *HTMLElement) OnLoadedData(v Action) *HTMLElement {
	e.a["onloadeddata"] = v
	return e
}

// OnLoadedMetaData sets the element's "onloadedmetadata" attribute
func (e *HTMLElement) OnLoadedMetaData(v Action) *HTMLElement {
	e.a["onloadedmetadata"] = v
	return e
}

// OnLoadStart sets the element's "onloadstart" attribute
func (e *HTMLElement) OnLoadStart(v Action) *HTMLElement {
	e.a["onloadstart"] = v
	return e
}

// OnMouseDown sets the element's "onmousedown" attribute
func (e *HTMLElement) OnMouseDown(v Action) *HTMLElement {
	e.a["onmousedown"] = v
	return e
}

// OnMouseEnter sets the element's "onmouseenter" attribute
func (e *HTMLElement) OnMouseEnter(v Action) *HTMLElement {
	e.a["onmouseenter"] = v
	return e
}

// OnMouseLeave sets the element's "onmouseleave" attribute
func (e *HTMLElement) OnMouseLeave(v Action) *HTMLElement {
	e.a["onmouseleave"] = v
	return e
}

// OnMouseMove sets the element's "onmousemove" attribute
func (e *HTMLElement) OnMouseMove(v Action) *HTMLElement {
	e.a["onmousemove"] = v
	return e
}

// OnMouseOut sets the element's "onmouseout" attribute
func (e *HTMLElement) OnMouseOut(v Action) *HTMLElement {
	e.a["onmouseout"] = v
	return e
}

// OnMouseOver sets the element's "onmouseover" attribute
func (e *HTMLElement) OnMouseOver(v Action) *HTMLElement {
	e.a["onmouseover"] = v
	return e
}

// OnMouseUp sets the element's "onmouseup" attribute
func (e *HTMLElement) OnMouseUp(v Action) *HTMLElement {
	e.a["onmouseup"] = v
	return e
}

// OnWheel sets the element's "onwheel" attribute
func (e *HTMLElement) OnWheel(v Action) *HTMLElement {
	e.a["onwheel"] = v
	return e
}

// OnPause sets the element's "onpause" attribute
func (e *HTMLElement) OnPause(v Action) *HTMLElement {
	e.a["onpause"] = v
	return e
}

// OnPlay sets the element's "onplay" attribute
func (e *HTMLElement) OnPlay(v Action) *HTMLElement {
	e.a["onplay"] = v
	return e
}

// OnPlaying sets the element's "onplaying" attribute
func (e *HTMLElement) OnPlaying(v Action) *HTMLElement {
	e.a["onplaying"] = v
	return e
}

// OnProgress sets the element's "onprogress" attribute
func (e *HTMLElement) OnProgress(v Action) *HTMLElement {
	e.a["onprogress"] = v
	return e
}

// OnRateChange sets the element's "onratechange" attribute
func (e *HTMLElement) OnRateChange(v Action) *HTMLElement {
	e.a["onratechange"] = v
	return e
}

// OnReset sets the element's "onreset" attribute
func (e *HTMLElement) OnReset(v Action) *HTMLElement {
	e.a["onreset"] = v
	return e
}

// OnResize sets the element's "onresize" attribute
func (e *HTMLElement) OnResize(v Action) *HTMLElement {
	e.a["onresize"] = v
	return e
}

// OnScroll sets the element's "onscroll" attribute
func (e *HTMLElement) OnScroll(v Action) *HTMLElement {
	e.a["onscroll"] = v
	return e
}

// OnSeeked sets the element's "onseeked" attribute
func (e *HTMLElement) OnSeeked(v Action) *HTMLElement {
	e.a["onseeked"] = v
	return e
}

// OnSeeking sets the element's "onseeking" attribute
func (e *HTMLElement) OnSeeking(v Action) *HTMLElement {
	e.a["onseeking"] = v
	return e
}

// OnSelect sets the element's "onselect" attribute
func (e *HTMLElement) OnSelect(v Action) *HTMLElement {
	e.a["onselect"] = v
	return e
}

// OnShow sets the element's "onshow" attribute
func (e *HTMLElement) OnShow(v Action) *HTMLElement {
	e.a["onshow"] = v
	return e
}

// OnStalled sets the element's "onstalled" attribute
func (e *HTMLElement) OnStalled(v Action) *HTMLElement {
	e.a["onstalled"] = v
	return e
}

// OnSubmit sets the element's "onsubmit" attribute
func (e *HTMLElement) OnSubmit(v Action) *HTMLElement {
	e.a["onsubmit"] = v
	return e
}

// OnSuspend sets the element's "onsuspend" attribute
func (e *HTMLElement) OnSuspend(v Action) *HTMLElement {
	e.a["onsuspend"] = v
	return e
}

// OnTimeUpdate sets the element's "ontimeupdate" attribute
func (e *HTMLElement) OnTimeUpdate(v Action) *HTMLElement {
	e.a["ontimeupdate"] = v
	return e
}

// OnToggle sets the element's "ontoggle" attribute
func (e *HTMLElement) OnToggle(v Action) *HTMLElement {
	e.a["ontoggle"] = v
	return e
}

// OnVolumeChange sets the element's "onvolumechange" attribute
func (e *HTMLElement) OnVolumeChange(v Action) *HTMLElement {
	e.a["onvolumechange"] = v
	return e
}

// OnWaiting sets the element's "onwaiting" attribute
func (e *HTMLElement) OnWaiting(v Action) *HTMLElement {
	e.a["onwaiting"] = v
	return e
}

// OnCopy sets the element's "oncopy" attribute
func (e *HTMLElement) OnCopy(v Action) *HTMLElement {
	e.a["oncopy"] = v
	return e
}

// OnCut sets the element's "oncut" attribute
func (e *HTMLElement) OnCut(v Action) *HTMLElement {
	e.a["oncut"] = v
	return e
}

// OnPaste sets the element's "onpaste" attribute
func (e *HTMLElement) OnPaste(v Action) *HTMLElement {
	e.a["onpaste"] = v
	return e
}

// ContentEditable sets the element's "contenteditable" attribute
func (e *HTMLElement) ContentEditable(v string) *HTMLElement {
	e.a["contenteditable"] = v
	return e
}
