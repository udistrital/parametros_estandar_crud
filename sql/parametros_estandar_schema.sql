--Schema parametros_estandar :)
CREATE SCHEMA parametros_estandar;

--Creacion tabla unidad
CREATE SEQUENCE parametros_estandar.unidad_id_seq
	INCREMENT BY 1 
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_estandar.unidad(
	id integer NOT NULL DEFAULT nextval('parametros_estandar.unidad_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
    tipo character varying(100),
    valor character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_unidad PRIMARY KEY (id)
);
COMMENT ON TABLE parametros_estandar.unidad IS 'Tabla que parametriza los diferentes tipos de unidad de medida entendidos como: cantidad estandarizada de una determinada magnitud física, definida y adoptada por convención o por ley.';
COMMENT ON COLUMN parametros_estandar.unidad.id IS 'Identificador unico de la tabla unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_estandar.unidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de la unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.tipo IS 'Magnitud fundamental de la unidad acorde al SI, Ejemplo: Longitud, Masa, Tiempo, Intensidad eléctrica, Temperatura entre otras.';
COMMENT ON COLUMN parametros_estandar.unidad.valor IS 'Campo opcional para definir la equivalencia de la unidad de medida en el SI.';
COMMENT ON COLUMN parametros_estandar.unidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_estandar.unidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_estandar.unidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_estandar.unidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_estandar.unidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--Tomado de agora.unidad
/**
CREATE TABLE agora.unidad
(
  id integer NOT NULL DEFAULT nextval('agora.param_unidades_id_unidad_seq'::regclass),
  unidad character varying(250) NOT NULL,
  tipo character varying(250) NOT NULL,
  descripcion character varying(250), -- descripcion del parametro
  estado boolean NOT NULL, -- estado dle registro
  CONSTRAINT pk_unidad PRIMARY KEY (id)
)
**/

--Creacion tabla tipo_periodo
CREATE SEQUENCE parametros_estandar.tipo_periodo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_estandar.tipo_periodo(
	id integer NOT NULL DEFAULT nextval('parametros_estandar.tipo_periodo_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_periodo PRIMARY KEY (id)
);
COMMENT ON TABLE parametros_estandar.tipo_periodo IS 'Tabla que parametriza diferentes tipos de periodos de tiempo, Ejemplo: Bimestre, Trimestre, Cuatrimestre, Semetre entre otros.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.id IS 'Identificador unico de la tabla tipo_periodo.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_periodo.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_estandar.tipo_periodo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla periodo
CREATE SEQUENCE parametros_estandar.periodo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_estandar.periodo(
	id integer NOT NULL DEFAULT nextval('parametros_estandar.periodo_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
    tipo_periodo_id integer NOT NULL,
    fecha_inicio_periodo TIMESTAMP,
	fecha_fin_periodo TIMESTAMP,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_periodo PRIMARY KEY (id),
    CONSTRAINT fk_tipo_periodo_periodo FOREIGN KEY (tipo_periodo_id) REFERENCES parametros_estandar.tipo_periodo(id)
);
COMMENT ON TABLE parametros_estandar.periodo IS 'Tabla que parametriza los tiempos de vigencia de un periodo.';
COMMENT ON COLUMN parametros_estandar.periodo.id IS 'Identificador unico de la tabla tipo_periodo.';
COMMENT ON COLUMN parametros_estandar.periodo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_estandar.periodo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_periodo.';
COMMENT ON COLUMN parametros_estandar.periodo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_estandar.periodo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_estandar.periodo.tipo_periodo_id IS 'Identificador unico de la tabla tipo_periodo.';
COMMENT ON COLUMN parametros_estandar.periodo.fecha_inicio_periodo IS 'Fecha en la que el periodo de tiempo entra en vigencia.';
COMMENT ON COLUMN parametros_estandar.periodo.fecha_fin_periodo IS 'Fecha en la que el periodo de tiempo termina.';
COMMENT ON COLUMN parametros_estandar.periodo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_estandar.periodo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

-Opcion 2 tipo_unidad, unidad

--Creacion tabla tipo_unidad
CREATE SEQUENCE parametros_estandar.tipo_unidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_estandar.tipo_unidad(
	id integer NOT NULL DEFAULT nextval('parametros_estandar.tipo_unidad_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_unidad PRIMARY KEY (id)
);
COMMENT ON TABLE parametros_estandar.tipo_unidad IS 'Tabla que parametriza diferentes tipos de unidades tanto de medida como de tiempo, Ejemplo: Bimestre, Trimestre, Cuatrimestre, Semetre, Longitud, Masa, Tiempo, Intensidad eléctrica, Temperatura entre otras.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.id IS 'Identificador unico de la tabla tipo_unidad.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre de la unidad de medida.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo de unidad.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_estandar.tipo_unidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla unidad
CREATE SEQUENCE parametros_estandar.unidad_id_seq
	INCREMENT BY 1 
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_estandar.unidad(
	id integer NOT NULL DEFAULT nextval('parametros_estandar.unidad_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	tipo_unidad_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_unidad PRIMARY KEY (id),
	CONSTRAINT fk_tipo_unidad_unidad FOREIGN KEY (tipo_unidad_id) REFERENCES parametros_estandar.tipo_unidad(id)
);
COMMENT ON TABLE parametros_estandar.unidad IS 'Tabla que parametriza los diferentes tipos de unidad de medida entendidos como: cantidad estandarizada de una determinada magnitud física, definida y adoptada por convención o por leyTabla que parametriza las diferentes unidades de medida entendidos como: cantidad estandarizada de una determinada magnitud referenciada en tipo_unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.id IS 'Identificador unico de la tabla unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_estandar.unidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de la unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_estandar.unidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_estandar.unidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_estandar.unidad.tipo_unidad_id IS 'Identificador unico de la tabla tipo_unidad.';
COMMENT ON COLUMN parametros_estandar.unidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_estandar.unidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
