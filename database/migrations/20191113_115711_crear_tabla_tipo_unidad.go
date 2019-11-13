package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoUnidad_20191113_115711 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoUnidad_20191113_115711{}
	m.Created = "20191113_115711"

	migration.Register("CrearTablaTipoUnidad_20191113_115711", m)
}

// Run the migrations
func (m *CrearTablaTipoUnidad_20191113_115711) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_estandar.tipo_unidad( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_unidad PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE parametros_estandar.tipo_unidad OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_estandar.tipo_unidad IS 'Tabla que parametriza diferentes tipos de unidades tanto de medida como de tiempo, Ejemplo: Bimestre, Trimestre, Cuatrimestre, Semetre, Longitud, Masa, Tiempo, Intensidad eléctrica, Temperatura entre otras.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.id IS 'Identificador unico de la tabla tipo_unidad.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre de la unidad de medida.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo de unidad.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_estandar.tipo_unidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoUnidad_20191113_115711) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_estandar.tipo_unidad")
}
