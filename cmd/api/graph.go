package main

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

/*
===========================================================
GraphQL Object Struct:
  - MyDBinfo: Pointer to List of all DB objets
  - QueryString: Query recieved
  - GQSchemaConfig: GQL SchemaConfig
  - gqfields: GQL Filed (functions)
  - gqobj: GQL Objects (variables)

===========================================================
*/
type Graph struct {
	MyDBinfo       *[]DBinfo
	QueryString    string
	GQSchemaConfig graphql.SchemaConfig
	gqfields       graphql.Fields
	gqobj          *graphql.Object
}

// func New() creates a pointer to a new Graph Object, using Pointer to List of all DB objets
func New(dbinfo *[]DBinfo) *Graph {
	// Define GQL objects

	var ComplexType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Complex",
			Fields: graphql.Fields{
				"C1param": &graphql.Field{
					Type: graphql.String,
				},
				"C2param": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	var gqobj = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "DBinfo",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Complex": &graphql.Field{
					Type: ComplexType,
				},
			},
		},
	)

	// Define GQL Fields
	var gqfields = graphql.Fields{
		"getDB": &graphql.Field{
			Type:        graphql.NewList(gqobj),
			Description: "Get DB Info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return []*DBinfo{&(*dbinfo)[0]}, nil
			},
		},
		"setDB": &graphql.Field{
			Type:        graphql.NewList(gqobj),
			Description: "Set DB Info",
			Args: graphql.FieldConfigArgument{
				"ID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"Name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"ComplexArg": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
						Name: "ComplexArg",
						Fields: graphql.InputObjectConfigFieldMap{
							"C1param": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"C2param": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
					)),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				(*dbinfo)[0].ID, _ = p.Args["ID"].(int)
				(*dbinfo)[0].Name, _ = p.Args["Name"].(string)
				// fmt.Printf("params: %+v\n", p.Args["ComplexArg"])
				a := p.Args["ComplexArg"]
				complexag := a.(map[string]interface{})
				(*dbinfo)[0].Complex.C1param, _ = complexag["C1param"].(string)
				(*dbinfo)[0].Complex.C2param, _ = complexag["C2param"].(string)

				return []*DBinfo{&(*dbinfo)[0]}, nil
			},
		},
	}

	return &Graph{
		MyDBinfo: dbinfo,
		gqfields: gqfields,
		gqobj:    gqobj,
	}
}

func (g *Graph) Query() (*graphql.Result, error) {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: g.gqfields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	params := graphql.Params{Schema: schema, RequestString: g.QueryString}
	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		fmt.Printf("Error Print %+v", resp.Errors)
		return nil, errors.New("error executing query")
	}
	return resp, nil
}
