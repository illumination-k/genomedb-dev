package goterm

import (
	"fmt"
	"strings"
)

type OboHeader struct {
	Version        string
	DataVersion    string
	SubsetDefs     []string
	Synonymtypedef []string
}

type OboParser struct {
	curTermsType string
	curGoTerm    GoTerm
	curTypedef   map[string]string
	Header       OboHeader
	Typedefs     []map[string]string
	GoTerms      []GoTerm
}

func NewOboParser() *OboParser {
	parser := new(OboParser)
	parser.curTermsType = "header"

	return parser
}

func (p *OboParser) pushTerm() {
	if p.curTermsType == "goterm" {
		p.GoTerms = append(p.GoTerms, p.curGoTerm)
	} else if p.curTermsType == "typedef" {
		p.Typedefs = append(p.Typedefs, p.curTypedef)
	}
}

func (p *OboParser) ConsumeLine(line string) (err error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	if strings.HasPrefix(line, "[Term]") {
		p.pushTerm()
		p.curTermsType = "goterm"
		p.curGoTerm = *InitGoterm()
		return nil
	}

	if strings.HasPrefix(line, "[Typedef]") {
		p.pushTerm()
		p.curTermsType = "typedef"
		p.curTypedef = map[string]string{}
		return nil
	}

	if p.curTermsType == "goterm" {
		key, value, err := getKeyValueFromLine(line)

		if err != nil {
			return err
		}

		switch key {
		case "id":
			p.curGoTerm.Id = value
		case "name":
			p.curGoTerm.Name = value
		case "namespace":
			p.curGoTerm.Namespace = value
		case "def":
			p.curGoTerm.Def = value
		case "is_obsolete":
			if value == "true" {
				p.curGoTerm.IsObsolete = true
			}
		case "synonym":
			p.curGoTerm.Synonyms = append(p.curGoTerm.Synonyms, value)
		case "xref":
			p.curGoTerm.Xrefs = append(p.curGoTerm.Xrefs, value)
		case "is_a":
			attrs := strings.Split(value, "!")
			target := strings.TrimSpace(attrs[0])
			p.curGoTerm.IsAs = append(p.curGoTerm.IsAs, target)
		case "relationship":
			//
			attrs := strings.Split(value, " ")
			name := attrs[0]
			targetInfo := strings.Split(attrs[1], "!")

			target := strings.TrimSpace(targetInfo[0])
			p.curGoTerm.Relationships = append(
				p.curGoTerm.Relationships,
				Relationship{Name: name, Target: target},
			)
		}
	} else if p.curTermsType == "typedef" {
		// ToDo
	} else if p.curTermsType == "header" {
		key, value, err := getKeyValueFromLine(line)

		if err != nil {
			return err
		}

		switch key {
		case "format-version":
			p.Header.Version = value
		case "data-version":
			p.Header.DataVersion = value
		case "subsetdef":
			p.Header.SubsetDefs = append(p.Header.SubsetDefs, value)
		case "synonymtypedef":
			p.Header.Synonymtypedef = append(p.Header.Synonymtypedef, value)
		}
	} else {
		return fmt.Errorf("Unreacheable")
	}

	return err
}

func (p *OboParser) Flush() {
	p.pushTerm()
}

func getKeyValueFromLine(line string) (key string, value string, err error) {
	attrs := strings.Split(line, ":")

	if len(attrs) < 2 {
		return key, value, fmt.Errorf("Invalid value line: %s\n", line)
	}

	key = attrs[0]
	value = strings.TrimSpace(strings.Join(attrs[1:], ":"))

	return key, value, nil
}
