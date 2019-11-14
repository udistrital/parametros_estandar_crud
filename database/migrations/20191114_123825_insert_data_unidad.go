package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertDataUnidad_20191114_123825 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertDataUnidad_20191114_123825{}
	m.Created = "20191114_123825"

	migration.Register("InsertDataUnidad_20191114_123825", m)
}

// Run the migrations
func (m *InsertDataUnidad_20191114_123825) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('NO APLICA', 'NO APLICA', FALSE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('UNIDAD', 'UNIDAD', TRUE, 1, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MILÍMETRO', 'MILÍMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('CENTIMETRO', 'CENTIMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÍMETRO', 'DECÍMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('METRO', 'METRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÁMETRO', 'DECÁMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HECTÓMETRO', 'HECTÓMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('KILOMETRO', 'KILOMETRO', TRUE, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('CENTIGRAMO', 'CENTIGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MILIGRAMO', 'MILIGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECIGRAMO', 'DECIGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('GRAMO', 'GRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECAGRAMO', 'DECAGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HECTOGRAMO', 'HECTOGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('KILOGRAMO', 'KILOGRAMO', TRUE, 3, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('SEGUNDOS', 'SEGUNDOS', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MINUTOS', 'MINUTOS', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HORAS', 'HORAS', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DIAS', 'DIAS', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MESES', 'MESES', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('AÑOS', 'AÑOS', TRUE, 4, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('CENTILITRO', 'CENTILITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MILILITRO', 'MILILITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECILITRO', 'DECILITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('LITRO', 'LITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECALITRO', 'DECALITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HECTOLITRO', 'HECTOLITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('KILOLITRO', 'KILOLITRO', TRUE, 5, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MILÍMETRO CUADRADO', 'MILÍMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('CENTÍMETRO CUADRADO', 'CENTÍMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÍMETRO CUADRADO', 'DECÍMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('METRO CUADRADO', 'METRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÁMETRO CUADRADO', 'DECÁMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HECTÓMETRO CUADRADO', 'HECTÓMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('KILÓMETRO CUADRADO', 'KILÓMETRO CUADRADO', TRUE, 6, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('MILÍMETRO CÚBICO', 'MILÍMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('CENTÍMETRO CÚBICO', 'CENTÍMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÍMETRO CÚBICO', 'DECÍMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('METRO CÚBICO', 'METRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('DECÁMETRO CÚBICO', 'DECÁMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('HECTÓMETRO CÚBICO', 'HECTÓMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO parametros_estandar.unidad (nombre, codigo_abreviacion, activo, tipo_unidad_id, fecha_creacion, fecha_modificacion) VALUES ('KILÓMETRO CÚBICO', 'KILÓMETRO CÚBICO', TRUE, 7, LOCALTIMESTAMP, LOCALTIMESTAMP);")	
}

// Reverse the migrations
func (m *InsertDataUnidad_20191114_123825) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM parametros_estandar.unidad WHERE tipo_unidad_id BETWEEN 1 AND 7;")
}
