package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPeriodo_20191127_120059 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPeriodo_20191127_120059{}
	m.Created = "20191127_120059"

	migration.Register("CrearTablaPeriodo_20191127_120059", m)
}

// Run the migrations
func (m *CrearTablaPeriodo_20191127_120059) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_estandar.periodo( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), year integer, ciclo character varying(4), codigo_abreviacion character varying(20), activo boolean NOT NULL, aplicacion_id integer NOT NULL, unidad_id integer NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_periodo PRIMARY KEY (id), CONSTRAINT fk_unidad_periodo FOREIGN KEY (unidad_id) REFERENCES parametros_estandar.unidad(id) );")
	m.SQL("ALTER TABLE parametros_estandar.periodo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_estandar.periodo IS 'Tabla que parametriza los diferentes tipos de periodo asociados a un negocio/aplicacion y al tipo de unidad de tiempo que duran Ej: Bimestre, Semestre entre otros.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.id IS 'Identificador unico de la tabla unidad.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.year IS 'Campo que indica el año para el cual es vigente el periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.ciclo IS 'Campo que permite llevar un control de los diferentes periodos que puede tener un año para determiando negocio EJ: Periodos academicos 2020-1, 2020-2, 2020-3.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.aplicacion_id IS 'Identificador unico de la tabla aplicacion del esquema configuracion, define el contexto y el negocio en el que se usa el periodo de vigencia.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.unidad_id IS 'Identificador unico de la tabla unidad, define el periodo de tiempo que se usara Ejemplo: Bimestre, Trimestre, Semestre.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.periodo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaPeriodo_20191127_120059) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_estandar.periodo")
}
