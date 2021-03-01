// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/pet"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BeforeID holds the value of the "before_id" field.
	BeforeID float64 `json:"before_id,omitempty"`
	// AfterID holds the value of the "after_id" field.
	AfterID float64 `json:"after_id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges    CarEdges `json:"edges"`
	pet_cars *string
}

// CarEdges holds the relations/edges for other nodes in the graph.
type CarEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Pet `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarEdges) OwnerOrErr() (*Pet, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldBeforeID, car.FieldAfterID:
			values[i] = &sql.NullFloat64{}
		case car.FieldID:
			values[i] = &sql.NullInt64{}
		case car.FieldModel:
			values[i] = &sql.NullString{}
		case car.ForeignKeys[0]: // pet_cars
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Car", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case car.FieldBeforeID:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field before_id", values[i])
			} else if value.Valid {
				c.BeforeID = value.Float64
			}
		case car.FieldAfterID:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field after_id", values[i])
			} else if value.Valid {
				c.AfterID = value.Float64
			}
		case car.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				c.Model = value.String
			}
		case car.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pet_cars", values[i])
			} else if value.Valid {
				c.pet_cars = new(string)
				*c.pet_cars = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Car entity.
func (c *Car) QueryOwner() *PetQuery {
	return (&CarClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", before_id=")
	builder.WriteString(fmt.Sprintf("%v", c.BeforeID))
	builder.WriteString(", after_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AfterID))
	builder.WriteString(", model=")
	builder.WriteString(c.Model)
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
