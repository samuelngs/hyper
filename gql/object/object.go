package object

import (
	"github.com/graphql-go/graphql"
	"github.com/samuelngs/hyper/gql/interfaces"
)

type object struct {
	name, description string
	args              []interfaces.Argument
	fields            []interfaces.Field
}

func (v *object) Name(s string) interfaces.Object {
	v.name = s
	return v
}

func (v *object) Description(s string) interfaces.Object {
	v.description = s
	return v
}

func (v *object) Fields(fs ...interfaces.Field) interfaces.Object {
	for _, f := range fs {
		if f != nil {
			v.fields = append(v.fields, f)
		}
	}
	return v
}

func (v *object) Args(args ...interfaces.Argument) interfaces.Object {
	for _, arg := range args {
		if arg != nil {
			v.args = append(v.args, arg)
		}
	}
	return v
}

func (v *object) ToObject() *graphql.Object {
	fields := graphql.Fields{}
	for _, f := range v.fields {
		v := f.Compile()
		fields[v.Name] = v
	}
	c := graphql.ObjectConfig{
		Name:        v.name,
		Description: v.description,
		Fields:      fields,
	}
	return graphql.NewObject(c)
}

func (v *object) ToInputObject() *graphql.InputObject {
	args := graphql.InputObjectConfigFieldMap{}
	for _, arg := range v.args {
		k, v := arg.ToInputObjectFieldConfig()
		args[k] = v
	}
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        v.name,
		Description: v.description,
		Fields:      args,
	})
}

func (v *object) ExportFields() []interfaces.Field {
	return v.fields
}

func (v *object) ExportArgs() []interfaces.Argument {
	return v.args
}

// New creates a new object
func New(opt ...Option) interfaces.Object {
	opts := newOptions(opt...)
	return &object{
		name:        opts.Name,
		description: opts.Description,
		fields:      opts.Fields,
	}
}
