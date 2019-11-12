package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaParametrosEstandar_20191112_121414 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaParametrosEstandar_20191112_121414{}
	m.Created = "20191112_121414"

	migration.Register("CrearSchemaParametrosEstandar_20191112_121414", m)
}

// Run the migrations
func (m *CrearSchemaParametrosEstandar_20191112_121414) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS parametros_estandar;")
	m.SQL("ALTER SCHEMA parametros_estandar OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,parametros_estandar;")
}

// Reverse the migrations
func (m *CrearSchemaParametrosEstandar_20191112_121414) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS parametros_estandar");
}
