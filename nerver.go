package nerver

import (
	"encoding/json"
	"strings"
	"github.com/beevik/etree"
	"errors"
)

type prop struct {
	Name     string
	Required bool
}

type child struct {
	Element  node
	Required bool
}

type node struct {
	Name     string
	Props    []prop
	Children []child
}

type Token struct {
	Name string
	Value string
}

type Validation struct {
	Xml    string
	Schema string
	Tokens []Token
}

func (v *Validation) Ok() (bool, error) {
	var node node
	var tokens []Token
	v.Tokens = tokens
	err := json.NewDecoder(strings.NewReader(v.Schema)).Decode(&node)
	if err != nil {
		return false, err
	}
	doc := etree.NewDocument()
	doc.ReadFromString(v.Xml)
	return v.valid(node, doc.SelectElement(node.Name), node.Name)
}

func (v *Validation) valid(node node, dom *etree.Element, namespace string) (bool, error) {
	_, err := v.validateNode(node, dom)
	if err != nil {
		return false, err
	}
	_, err = v.validateProps(node, dom, namespace)
	if err != nil {
		return false, err
	}
	_, err = v.validateChildren(node, dom, namespace)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (v *Validation) validateNode(node node, dom *etree.Element) (bool, error) {
	if dom == nil {
		return false, errors.New(node.Name + " expected, but not found")
	}
	if dom.Tag != node.Name {
		return false, errors.New(node.Name + " expected, " + dom.Tag + "found")
	}
	return true, nil
}

func (v *Validation) validateProps(node node, dom *etree.Element, namespace string) (bool, error) {
	for _, prop := range node.Props {
		attr := dom.SelectAttr(prop.Name)
		value := ""
		if attr != nil {
			value = attr.Value
		}
		token := Token{
			Value: value,
			Name: namespace + "_" + node.Name + "__" + prop.Name,
		}
		v.Tokens = append(v.Tokens, token)
		if !prop.Required {
			continue
		}
		// validate all props.
		if attr == nil {
			return false, errors.New(prop.Name + " required in " + namespace + "_" + dom.Tag)
		}
	}
	return true, nil
}

func (v *Validation) validateChildren(node node, dom *etree.Element, namespace string) (bool, error) {
	flag := true
	var err error = nil
	for _, schema := range node.Children {
		if !schema.Required {
			continue
		}
		v.Tokens = append(v.Tokens, Token{
			Value: dom.SelectElement(schema.Element.Name).Text(),
			Name: namespace + "_" + schema.Element.Name,
		})
		ok, e := v.valid(schema.Element, dom.SelectElement(schema.Element.Name), namespace + "_" + schema.Element.Name)
		if !ok {
			flag = false
			err = e
		}
	}
	return flag, err
}
