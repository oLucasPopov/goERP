CREATE TABLE configuracao_impressoras (
	id serial4 NOT NULL,
	nome_impressora varchar(64) NULL,
	apelido_impressora varchar(64) NULL,
	CONSTRAINT configuracao_impressoras_pkey PRIMARY KEY (id)
);