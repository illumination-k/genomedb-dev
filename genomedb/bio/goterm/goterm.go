package goterm

type Relationship struct {
	Name   string
	Target string
}

type GoTerm struct {
	Id            string
	Name          string
	Namespace     string
	Def           string
	IsObsolete    bool
	Xrefs         []string
	Synonyms      []string
	IsAs          []string
	Relationships []Relationship
}

func InitGoterm() *GoTerm {
	term := new(GoTerm)
	term.IsObsolete = false
	return term
}
