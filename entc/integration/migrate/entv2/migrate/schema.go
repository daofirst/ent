// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_car", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_users_car",
				Columns:    []*schema.Column{CarsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConversionsColumns holds the columns for the "conversions" table.
	ConversionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "int8_to_string", Type: field.TypeString, Nullable: true, Size: 6},
		{Name: "uint8_to_string", Type: field.TypeString, Nullable: true, Size: 6},
		{Name: "int16_to_string", Type: field.TypeString, Nullable: true, Size: 6},
		{Name: "uint16_to_string", Type: field.TypeString, Nullable: true, Size: 6},
		{Name: "int32_to_string", Type: field.TypeString, Nullable: true, Size: 12},
		{Name: "uint32_to_string", Type: field.TypeString, Nullable: true, Size: 12},
		{Name: "int64_to_string", Type: field.TypeString, Nullable: true, Size: 21},
		{Name: "uint64_to_string", Type: field.TypeString, Nullable: true, Size: 21},
	}
	// ConversionsTable holds the schema information for the "conversions" table.
	ConversionsTable = &schema.Table{
		Name:        "conversions",
		Columns:     ConversionsColumns,
		PrimaryKey:  []*schema.Column{ConversionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomTypesColumns holds the columns for the "custom_types" table.
	CustomTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "custom", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "customtype"}},
	}
	// CustomTypesTable holds the schema information for the "custom_types" table.
	CustomTypesTable = &schema.Table{
		Name:        "custom_types",
		Columns:     CustomTypesColumns,
		PrimaryKey:  []*schema.Column{CustomTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "source", Type: field.TypeString, Nullable: true},
		{Name: "source_uri", Type: field.TypeString, Nullable: true},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:        "media",
		Columns:     MediaColumns,
		PrimaryKey:  []*schema.Column{MediaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "media_source_source_uri",
				Unique:  true,
				Columns: []*schema.Column{MediaColumns[1], MediaColumns[2]},
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_users_pets",
				Columns:    []*schema.Column{PetsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeInt, Increment: true},
		{Name: "mixed_string", Type: field.TypeString, Default: "default"},
		{Name: "mixed_enum", Type: field.TypeEnum, Enums: []string{"on", "off"}, Default: "on"},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "nickname", Type: field.TypeString, Size: 255},
		{Name: "phone", Type: field.TypeString, Default: "unknown"},
		{Name: "buffer", Type: field.TypeBytes, Nullable: true},
		{Name: "title", Type: field.TypeString, Default: "SWE"},
		{Name: "renamed", Type: field.TypeString, Nullable: true},
		{Name: "blob", Type: field.TypeBytes, Nullable: true, Size: 1000},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"logged_in", "logged_out", "online"}},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"done", "pending"}},
		{Name: "workplace", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_phone_age",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[6], UsersColumns[3]},
			},
		},
	}
	// FriendsColumns holds the columns for the "friends" table.
	FriendsColumns = []*schema.Column{
		{Name: "user", Type: field.TypeInt},
		{Name: "friend", Type: field.TypeInt},
	}
	// FriendsTable holds the schema information for the "friends" table.
	FriendsTable = &schema.Table{
		Name:       "friends",
		Columns:    FriendsColumns,
		PrimaryKey: []*schema.Column{FriendsColumns[0], FriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friends_user",
				Columns:    []*schema.Column{FriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "friends_friend",
				Columns:    []*schema.Column{FriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		ConversionsTable,
		CustomTypesTable,
		GroupsTable,
		MediaTable,
		PetsTable,
		UsersTable,
		FriendsTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	FriendsTable.ForeignKeys[0].RefTable = UsersTable
	FriendsTable.ForeignKeys[1].RefTable = UsersTable
}
