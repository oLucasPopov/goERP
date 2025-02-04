CREATE TABLE empresas (
	id serial4 NOT NULL,
	nome_fantasia varchar(256) NULL,
	razao_social varchar(256) NULL,
	cpf_cnpj varchar(14) NULL,
	telefone varchar(11) NULL,
	celular varchar(11) NULL,
	email varchar(256) NULL,
	id_estado int4 NULL,
	id_cidade int4 NULL,
	cep varchar(8) NULL,
	rua varchar(256) NULL,
	numero varchar(5) NULL,
	bairro varchar(128) NULL,
	referencia varchar(128) NULL,
	link_facebook varchar(256) NULL,
	usuario_instagram varchar(64) NULL,
	CONSTRAINT empresas_pkey PRIMARY KEY (id),
	CONSTRAINT empresas_id_cidade_fkey FOREIGN KEY (id_cidade) REFERENCES cidades(id),
	CONSTRAINT empresas_id_estado_fkey FOREIGN KEY (id_estado) REFERENCES estados(id)
);