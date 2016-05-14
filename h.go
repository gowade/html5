package h

import (
	"fmt"

	"github.com/gowade/vdom"
)

type StyleMap map[string]interface{}

type Node interface {
	ToVNode() vdom.VNode
}

type TextNode string

func (t TextNode) ToVNode() vdom.VNode {
	return vdom.VText(t)
}

func (e *htmlElement) ToVNode() vdom.VNode {
	ve := vdom.VElement{
		TagName:   e.tagName,
		Namespace: e.namespace,
		Key:       e.key,
		Props:     e.a,
		Children:  make([]vdom.VNode, len(e.c)),
	}

	for i := range e.c {
		ve.Children[i] = e.c[i].ToVNode()
	}

	return ve
}

func F(formatOrValue interface{}, args ...interface{}) string {
	if s, ok := formatOrValue.(string); ok {
		if len(args) == 0 {
			return s
		}
		return fmt.Sprintf(s, args...)
	}

	return fmt.Sprint(formatOrValue)
}

func T(formatOrValue interface{}, args ...interface{}) TextNode {
	return TextNode(F(formatOrValue, args...))
}

type L []Node

func (l *L) Add(node Node) {
	*l = append(*l, node)
}

type htmlElement struct {
	a         map[string]interface{}
	c         L
	style     map[string]interface{}
	tagName   string
	namespace string
	key       string
}

func (e *htmlElement) C(children ...Node) *htmlElement {
	e.c = children
	return e
}

func (e *htmlElement) L(iterateFn func(l *L)) *htmlElement {
	iterateFn(&e.c)
	return e
}

func (e *htmlElement) T(formatOrValue interface{}, args ...interface{}) *htmlElement {
	return e.C(T(formatOrValue, args...))
}

func (e *htmlElement) A(name string, value interface{}) *htmlElement {
	if valueBool, ok := value.(bool); ok {
		if valueBool {
			e.a[name] = ""
		} else {
			delete(e.a, name)
		}
		return e
	}

	e.a[name] = value
	return e
}

func (e *htmlElement) S(style StyleMap) *htmlElement {
	e.style = style
	return e
}

func Element(tagName string) *htmlElement {
	return &htmlElement{tagName: tagName, a: make(map[string]interface{})}
}

func ElementNS(tagName string, namespace string) *htmlElement {
	e := Element(tagName)
	e.namespace = namespace
	return e
}
