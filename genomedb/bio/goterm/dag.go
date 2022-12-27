package goterm

import (
	"fmt"
	"genomedb/ds/hashset"
)

const INF int = 10000000

type GoDAGRecord struct {
	GoTermId string
	// Depth: longest distance from root node
	Depth int
	// Level: shortest distance from root node
	Level    int
	Parent   []string
	Children []string
}

func (dag GoDAG) GetAllChildren(id string) (hashset.HashSet[string], error) {
	record, found := dag[id]
	set := hashset.NewHashSet[string]()

	if !found {
		return set, fmt.Errorf("%s is Not Found in GoDAG", id)
	}

	for _, parent_id := range record.Children {
		set.Add(parent_id)
		children_ids, err := dag.GetAllChildren(parent_id)
		if err != nil {
			return set, err
		}

		for children_id := range children_ids {
			set.Add(children_id)
		}
	}

	return set, nil
}

func (r GoDAGRecord) depthIsInitValue() bool {
	return r.Depth == -INF
}

func (r GoDAGRecord) levelIsInitValue() bool {
	return r.Level == INF
}

func NewGoDAGRecord(go_term_id string) (record GoDAGRecord) {
	record.GoTermId = go_term_id
	record.Depth = -INF
	record.Level = INF
	return record
}

type GoDAG map[string]GoDAGRecord

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (dag *GoDAG) setDepth(id string) int {
	rec := (*dag)[id]

	if rec.depthIsInitValue() {
		if len(rec.Parent) != 0 {
			for _, parent_rec := range rec.Parent {
				rec.Depth = max(dag.setDepth(parent_rec), rec.Depth) + 1
			}
		} else {
			rec.Depth = 0
		}
	}

	return rec.Depth
}

func (dag *GoDAG) setLevel(id string) int {
	rec := (*dag)[id]

	if rec.levelIsInitValue() {
		if len(rec.Parent) != 0 {
			for _, parent_rec := range rec.Parent {
				rec.Level = min(dag.setLevel(parent_rec), rec.Level) + 1
			}
		} else {
			rec.Level = 0
		}
	}

	return rec.Level
}

func CreateGoDAGFromGoterms(goterms []GoTerm) (godag GoDAG) {
	godag = map[string]GoDAGRecord{}
	// initilize go term
	for _, goterm := range goterms {
		// set is_a
		// set parents
		rec, found := godag[goterm.Id]

		if !found {
			rec = NewGoDAGRecord(goterm.Id)
		}

		rec.Parent = goterm.IsAs
		godag[goterm.Id] = rec

		// set children
		for _, parent_goterm_id := range goterm.IsAs {
			parent, found := godag[parent_goterm_id]

			if !found {
				parent = NewGoDAGRecord(parent_goterm_id)
			}

			parent.Children = append(parent.Children, goterm.Id)
			godag[parent_goterm_id] = parent
		}
	}

	// Set go depth and level
	for id := range godag {
		rec := godag[id]
		if rec.depthIsInitValue() {
			rec.Depth = godag.setDepth(id)
		}

		if rec.levelIsInitValue() {
			rec.Level = godag.setLevel(id)
		}

		godag[id] = rec
	}

	return godag
}
