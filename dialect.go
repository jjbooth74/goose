package goose

import (
	"database/sql"
	"fmt"
)

// SQLDialect abstracts the details of specific SQL dialects
// for goose's few SQL specific statements
type SQLDialect interface {
	CreateVersionTableSQL() string // sql string to create the goose_db_version table
	InsertVersionSQL() string      // sql string to insert the initial version table row
	DbVersionQuery(db *sql.DB) (*sql.Rows, error)
}

var dialect SQLDialect = &PostgresDialect{}

// GetDialect gets the SQLDialect
func GetDialect() SQLDialect {
	return dialect
}

// SetCustomDialect allows you to pass in a non-standard dialect.
func SetCustomDialect(d SQLDialect) {
	dialect = d
}

// SetDialect sets the SQLDialect
func SetDialect(d string) error {
	switch d {
	case "postgres":
		dialect = &PostgresDialect{}
	case "mysql":
		dialect = &MySQLDialect{}
	case "sqlite3":
		dialect = &Sqlite3Dialect{}
	case "redshift":
		dialect = &RedshiftDialect{}
	case "tidb":
		dialect = &TiDBDialect{}
	default:
		return fmt.Errorf("%q: unknown dialect", d)
	}

	return nil
}

////////////////////////////
// Postgres
////////////////////////////

// PostgresDialect struct.
type PostgresDialect struct{}

func (pg PostgresDialect) CreateVersionTableSQL() string {
	return `CREATE TABLE goose_db_version (
            	id serial NOT NULL,
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default now(),
                PRIMARY KEY(id)
            );`
}

func (pg PostgresDialect) InsertVersionSQL() string {
	return "INSERT INTO goose_db_version (version_id, is_applied) VALUES ($1, $2);"
}

func (pg PostgresDialect) DbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from goose_db_version ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	return rows, err
}

////////////////////////////
// MySQL
////////////////////////////

// MySQLDialect struct.
type MySQLDialect struct{}

func (m MySQLDialect) CreateVersionTableSQL() string {
	return `CREATE TABLE goose_db_version (
                id serial NOT NULL,
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default now(),
                PRIMARY KEY(id)
            );`
}

func (m MySQLDialect) InsertVersionSQL() string {
	return "INSERT INTO goose_db_version (version_id, is_applied) VALUES (?, ?);"
}

func (m MySQLDialect) DbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from goose_db_version ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	return rows, err
}

////////////////////////////
// sqlite3
////////////////////////////

// Sqlite3Dialect struct.
type Sqlite3Dialect struct{}

func (m Sqlite3Dialect) CreateVersionTableSQL() string {
	return `CREATE TABLE goose_db_version (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                version_id INTEGER NOT NULL,
                is_applied INTEGER NOT NULL,
                tstamp TIMESTAMP DEFAULT (datetime('now'))
            );`
}

func (m Sqlite3Dialect) InsertVersionSQL() string {
	return "INSERT INTO goose_db_version (version_id, is_applied) VALUES (?, ?);"
}

func (m Sqlite3Dialect) DbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from goose_db_version ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	return rows, err
}

////////////////////////////
// Redshift
////////////////////////////

// RedshiftDialect struct.
type RedshiftDialect struct{}

func (rs RedshiftDialect) CreateVersionTableSQL() string {
	return `CREATE TABLE goose_db_version (
            	id integer NOT NULL identity(1, 1),
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default sysdate,
                PRIMARY KEY(id)
            );`
}

func (rs RedshiftDialect) InsertVersionSQL() string {
	return "INSERT INTO goose_db_version (version_id, is_applied) VALUES ($1, $2);"
}

func (rs RedshiftDialect) DbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from goose_db_version ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	return rows, err
}

////////////////////////////
// TiDB
////////////////////////////

// TiDBDialect struct.
type TiDBDialect struct{}

func (m TiDBDialect) CreateVersionTableSQL() string {
	return `CREATE TABLE goose_db_version (
                id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE,
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default now(),
                PRIMARY KEY(id)
            );`
}

func (m TiDBDialect) InsertVersionSQL() string {
	return "INSERT INTO goose_db_version (version_id, is_applied) VALUES (?, ?);"
}

func (m TiDBDialect) DbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from goose_db_version ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	return rows, err
}
