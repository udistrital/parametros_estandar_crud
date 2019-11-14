package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertDataTipoUnidad_20191114_123512 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertDataTipoUnidad_20191114_123512{}
	m.Created = "20191114_123512"

	migration.Register("InsertDataTipoUnidad_20191114_123512", m)
}

// Run the migrations
func (m *InsertDataTipoUnidad_20191114_123512) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'SIN UNIDAD', 'SIN UNIDAD', 'Sin Unidad de medida', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'LONGITUD', 'LONGITUD', 'Unidad de medida de Longitud', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (3, 'MASA', 'MASA', 'Unidad de medida de Masa', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (4, 'TIEMPO', 'TIEMPO', 'Unidad de medida de Tiempo', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (5, 'CAPACIDAD', 'CAPACIDAD', 'Unidad de medida de Capacidad', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (6, 'SUPERFICIE', 'SUPERFICIE', 'Unidad de medida de Superficie', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.tipo_unidad (id, nombre, codigo_abreviacion, descripcion, activo, fecha_creacion, fecha_modificacion) VALUES (7, 'VOLUMEN', 'VOLUMEN', 'Unidad de medida de Volumen', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertDataTipoUnidad_20191114_123512) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM parametros_estandar.tipo_unidad WHERE id BETWEEN 1 AND 7;")
}
