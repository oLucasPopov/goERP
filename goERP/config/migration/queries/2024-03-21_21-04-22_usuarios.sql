CREATE TABLE usuarios (
	id serial4 NOT NULL,
	usuario varchar(64) NULL,
	senha varchar(512) NULL,
	CONSTRAINT usuarios_id_key UNIQUE (id)
);