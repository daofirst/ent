// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// TeamInverseLabel holds the string label denoting the team inverse edge type in the database.
	TeamInverseLabel = "user_team"
	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_pets"
)

// comment from another template.
