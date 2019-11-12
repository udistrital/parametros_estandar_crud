package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaUnidad_20191112_121554 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaUnidad_20191112_121554{}
	m.Created = "20191112_121554"

	migration.Register("CrearTablaUnidad_20191112_121554", m)
}

// Run the migrations
func (m *CrearTablaUnidad_20191112_121554) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_estandar.unidad( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), tipo character varying(100), valor character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_unidad PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE parametros_estandar.unidad OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_estandar.unidad IS 'Tabla que parametriza los diferentes tipos de unidad de medida entendidos como: cantidad estandarizada de una determinada magnitud física, definida y adoptada por convención o por ley.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.id IS 'Identificador unico de la tabla unidad.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de la unidad.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.tipo IS 'Magnitud fundamental de la unidad acorde al SI, Ejemplo: Longitud, Masa, Tiempo, Intensidad eléctrica, Temperatura entre otras.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.valor IS 'Campo opcional para definir la equivalencia de la unidad de medida en el SI.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.unidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")	
}

// Reverse the migrations
func (m *CrearTablaUnidad_20191112_121554) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_estandar.unidad")
}
