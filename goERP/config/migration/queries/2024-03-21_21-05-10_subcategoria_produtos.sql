CREATE TABLE subcategoria_produtos (
	id serial4 NOT NULL,
	id_categoria int4 NOT NULL,
	subcategoria varchar(128) NULL,
	CONSTRAINT subcategoria_produtos_pkey PRIMARY KEY (id),
	CONSTRAINT subcategoria_produtos_id_categoria_fkey FOREIGN KEY (id_categoria) REFERENCES categoria_produtos(id)
);