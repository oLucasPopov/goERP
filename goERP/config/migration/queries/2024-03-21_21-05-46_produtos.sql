CREATE TABLE produtos (
	id serial4 NOT NULL,
	data_inclusao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	data_alteracao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	descricao_produto varchar(256) NULL,
	codigo_barras varchar(13) NULL,
	preco_custo numeric(18, 2) NULL,
	preco_venda numeric(18, 2) NULL,
	tipo_unidade varchar(2) NULL,
	peso_bruto numeric(18, 3) NULL,
	peso_liquido numeric(18, 3) NULL,
	tipo_produto int4 NOT NULL DEFAULT 4,
	id_categoria int4 NULL,
	id_subcategoria int4 NULL,
	CONSTRAINT check_tipo_unidade CHECK (((tipo_unidade)::text = ANY (ARRAY[('KG'::character varying)::text, ('UN'::character varying)::text]))),
	CONSTRAINT produtos_id_key PRIMARY KEY (id),
	CONSTRAINT produtos_tipo_produto_check CHECK ((((tipo_produto >= 0) AND (tipo_produto <= 9)) OR (tipo_produto = 99))),
	CONSTRAINT produtos_id_categoria_fkey FOREIGN KEY (id_categoria) REFERENCES categoria_produtos(id),
	CONSTRAINT produtos_id_subcategoria_fkey FOREIGN KEY (id_subcategoria) REFERENCES subcategoria_produtos(id)
);