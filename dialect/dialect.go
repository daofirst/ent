// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dialect

import (
	"context"
	"database/sql/driver"
	"log"

	"time"

	"github.com/google/uuid"
)

// Dialect names for external usage.
const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
	Gremlin  = "gremlin"
)

// ExecQuerier wraps the 2 database operations.
type ExecQuerier interface {
	// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
	// It scans the result into the pointer v. In SQL, you it's usually sql.Result.
	Exec(ctx context.Context, query string, args, v interface{}) error
	// Query executes a query that returns rows, typically a SELECT in SQL.
	// It scans the result into the pointer v. In SQL, you it's usually *sql.Rows.
	Query(ctx context.Context, query string, args, v interface{}) error
}

// Driver is the interface that wraps all necessary operations for ent clients.
type Driver interface {
	ExecQuerier
	// Tx starts and returns a new transaction.
	// The provided context is used until the transaction is committed or rolled back.
	Tx(context.Context) (Tx, error)
	// Close closes the underlying connection.
	Close() error
	// Dialect returns the dialect name of the driver.
	Dialect() string
}

// Tx wraps the Exec and Query operations in transaction.
type Tx interface {
	ExecQuerier
	driver.Tx
}

type nopTx struct {
	Driver
}

func (nopTx) Commit() error   { return nil }
func (nopTx) Rollback() error { return nil }

// NopTx returns a Tx with a no-op Commit / Rollback methods wrapping
// the provided Driver d.
func NopTx(d Driver) Tx {
	return nopTx{d}
}

// DebugDriver is a driver that logs all driver operations.
type DebugDriver struct {
	Driver                                       // underlying driver.
	log    func(context.Context, ...interface{}) // log function. defaults to log.Println.
}

// Debug gets a driver and an optional logging function, and returns
// a new debugged-driver that prints all outgoing operations.
func Debug(d Driver, logger ...func(...interface{})) Driver {
	logf := log.Println
	if len(logger) == 1 {
		logf = logger[0]
	}
	drv := &DebugDriver{
		d,
		func(ctx context.Context, v ...interface{}) {
			var args []interface{}
			args = append(args, ctx)
			args = append(args, v...)
			logf(args...)
		},
	}
	return drv
}

// DebugWithContext gets a driver and a logging function, and returns
// a new debugged-driver that prints all outgoing operations with context.
func DebugWithContext(d Driver, logger func(context.Context, ...interface{})) Driver {
	drv := &DebugDriver{d, logger}
	return drv
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	var err error
	err = d.Driver.Exec(ctx, query, args, v)
	logInfo := map[string]interface{}{
		"driver": "driver.Exec",
		"query":  query,
		"args":   args,
		"cost":   time.Since(start),
		"err":    err,
	}

	d.log(ctx, logInfo)

	return err
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	var err error
	err = d.Driver.Query(ctx, query, args, v)
	logInfo := map[string]interface{}{
		"driver": "driver.Query",
		"query":  query,
		"args":   args,
		"cost":   time.Since(start),
		"err":    err,
	}

	d.log(ctx, logInfo)

	return err
}

// Tx adds an log-id for the transaction and calls the underlying driver Tx command.
func (d *DebugDriver) Tx(ctx context.Context) (Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()

	start := time.Now()
	logInfo := map[string]interface{}{
		"driver": "Tx.started",
		"cost":   time.Since(start),
		"txID":   id,
	}

	d.log(ctx, logInfo)

	return &DebugTx{tx, id, d.log, ctx}, nil
}

// DebugTx is a transaction implementation that logs all transaction operations.
type DebugTx struct {
	Tx                                        // underlying transaction.
	id  string                                // transaction logging id.
	log func(context.Context, ...interface{}) // log function. defaults to fmt.Println.
	ctx context.Context                       // underlying transaction context.
}

// Exec logs its params and calls the underlying transaction Exec method.
func (d *DebugTx) Exec(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	var err error
	err = d.Tx.Exec(ctx, query, args, v)
	logInfo := map[string]interface{}{
		"driver": "Tx.Exec",
		"query":  query,
		"args":   args,
		"cost":   time.Since(start),
		"err":    err,
		"txID":   d.id,
	}

	d.log(ctx, logInfo)

	return err
}

// Query logs its params and calls the underlying transaction Query method.
func (d *DebugTx) Query(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	var err error
	err = d.Tx.Query(ctx, query, args, v)
	logInfo := map[string]interface{}{
		"driver": "Tx.Query",
		"query":  query,
		"args":   args,
		"cost":   time.Since(start),
		"err":    err,
		"txID":   d.id,
	}

	d.log(ctx, logInfo)

	return err
}

// Commit logs this step and calls the underlying transaction Commit method.
func (d *DebugTx) Commit() error {
	start := time.Now()
	var err error
	err = d.Tx.Commit()
	logInfo := map[string]interface{}{
		"driver": "Tx.Commit",
		"cost":   time.Since(start),
		"err":    err,
		"txID":   d.id,
	}

	d.log(d.ctx, logInfo)

	return err

}

// Rollback logs this step and calls the underlying transaction Rollback method.
func (d *DebugTx) Rollback() error {
	start := time.Now()
	var err error
	err = d.Tx.Rollback()
	logInfo := map[string]interface{}{
		"driver": "Tx.Rollback",
		"cost":   time.Since(start),
		"err":    err,
		"txID":   d.id,
	}

	d.log(d.ctx, logInfo)

	return err
}
