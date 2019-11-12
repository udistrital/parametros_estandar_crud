package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPeriodo_20191112_122613 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPeriodo_20191112_122613{}
	m.Created = "20191112_122613"

	migration.Register("CrearTablaPeriodo_20191112_122613", m)
}

// Run the migrations
func (m *CrearTablaPeriodo_20191112_122613) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_estandar.periodo( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, tipo_periodo_id integer NOT NULL, fecha_inicio_periodo TIMESTAMP, fecha_fin_periodo TIMESTAMP, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_periodo PRIMARY KEY (id), CONSTRAINT fk_tipo_periodo_periodo FOREIGN KEY (tipo_periodo_id) REFERENCES parametros_estandar.tipo_periodo(id) );")
	m.SQL("ALTER TABLE parametros_estandar.periodo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_estandar.periodo IS 'Tabla que parametriza los tiempos de vigencia de un periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.id IS 'Identificador unico de la tabla tipo_periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.tipo_periodo_id IS 'Identificador unico de la tabla tipo_periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_inicio_periodo IS 'Fecha en la que el periodo de tiempo entra en vigencia.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_fin_periodo IS 'Fecha en la que el periodo de tiempo termina.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaPeriodo_20191112_122613) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_estandar.periodo")
}
