package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoPeriodo_20191112_122230 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoPeriodo_20191112_122230{}
	m.Created = "20191112_122230"

	migration.Register("CrearTablaTipoPeriodo_20191112_122230", m)
}

// Run the migrations
func (m *CrearTablaTipoPeriodo_20191112_122230) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_estandar.tipo_periodo( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_periodo PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE parametros_estandar.tipo_periodo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_estandar.tipo_periodo IS 'Tabla que parametriza diferentes tipos de periodos de tiempo, Ejemplo: Bimestre, Trimestre, Cuatrimestre, Semetre entre otros.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.id IS 'Identificador unico de la tabla tipo_periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_periodo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoPeriodo_20191112_122230) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_estandar.tipo_periodo")
}
