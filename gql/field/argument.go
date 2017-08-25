package field

import (
	"github.com/graphql-go/graphql"
	"github.com/samuelngs/hyper/gql/interfaces"
	"github.com/samuelngs/hyper/router"
)

type argument struct {
	name, description string
	typ               graphql.Input
	obj               interfaces.Object
	format            int
	def               interface{}
	require           bool
	compiledArgConf   *graphql.ArgumentConfig
	compiledObjConf   *graphql.InputObjectFieldConfig
}

func (v *argument) Name(s string) interfaces.Argument {
	v.name = s
	return v
}

func (v *argument) Description(s string) interfaces.Argument {
	v.description = s
	return v
}

func (v *argument) Type(o interface{}) interfaces.Argument {
	switch t := o.(type) {
	case graphql.Input:
		switch t {
		case graphql.Int:
			v.format = router.Int
		case graphql.Float:
			v.format = router.F64
		case graphql.String:
			v.format = router.Text
		case graphql.Boolean:
			v.format = router.Bool
		case graphql.ID:
			v.format = router.Text
		case graphql.DateTime:
			v.format = router.DateTimeRFC3339
		default:
			v.format = router.Any
		}
		v.typ = t
		v.obj = nil
	case interfaces.Object:
		v.typ = t.ToInputObject()
		v.obj = t
	}
	return v
}

func (v *argument) Default(o []byte) interfaces.Argument {
	v.def = o
	return v
}

func (v *argument) Require(b bool) interfaces.Argument {
	v.require = b
	return v
}

func (v *argument) InputObject() interfaces.Object {
	return v.obj
}

func (v *argument) ToArgumentConfig() (string, *graphql.ArgumentConfig) {
	if v.compiledArgConf == nil {
		var typ = v.typ
		if v.require {
			typ = graphql.NewNonNull(typ)
		}
		v.compiledArgConf = &graphql.ArgumentConfig{
			Type:         typ,
			DefaultValue: v.def,
			Description:  v.description,
		}
	}
	return v.name, v.compiledArgConf
}

func (v *argument) ToInputObjectFieldConfig() (string, *graphql.InputObjectFieldConfig) {
	if v.compiledObjConf == nil {
		var typ = v.typ
		if v.require {
			typ = graphql.NewNonNull(typ)
		}
		v.compiledObjConf = &graphql.InputObjectFieldConfig{
			Type:         typ,
			DefaultValue: v.def,
			Description:  v.description,
		}
	}
	return v.name, v.compiledObjConf
}

// NewArgument creates new argument instance
func NewArgument(s string) interfaces.Argument {
	return &argument{name: s}
}
