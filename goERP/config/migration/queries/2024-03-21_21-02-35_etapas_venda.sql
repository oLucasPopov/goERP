CREATE TABLE etapas_venda (
	id serial4 NOT NULL,
	indice_etapa int2 NULL,
	descricao varchar(64) NULL,
	etapa_ativa bool NULL,
	CONSTRAINT etapas_venda_pkey PRIMARY KEY (id)
);