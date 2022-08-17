package global

import (
	"fmt"
)

type Operation struct {
	Name         string
	Options      []opOption
	Arguments    []string
	SuppliedArgs []string
	Script       string
	Description  string
}

func (o Operation) String() string {
	return fmt.Sprintf("Operation %s: %v, %v, %v, %d, %s;", o.Name, o.Options, o.Arguments, o.SuppliedArgs, len(o.Script), o.Description)
}

func FindOperation(name string) (*Operation, error) {
	for _, op := range Ops {
		if op.Name == name {
			return &op, nil
		}
	}
	return nil, fmt.Errorf("Operation %s not found", name)
}
