CREATE TABLE tabela_preco_produtos (
	id serial4 NOT NULL,
	pro_codigo int4 NOT NULL,
	tp_codigo int4 NOT NULL,
	preco_unt_produto numeric(18, 2) NULL,
	CONSTRAINT tabela_preco_produtos_pkey PRIMARY KEY (id),
	CONSTRAINT tabela_preco_produtos_pro_codigo_fkey FOREIGN KEY (pro_codigo) REFERENCES produtos(id),
	CONSTRAINT tabela_preco_produtos_tp_codigo_fkey FOREIGN KEY (tp_codigo) REFERENCES tabelas_preco(id)
);
CREATE UNIQUE INDEX idx_cliente_produtos_tabela_precos ON public.tabela_preco_produtos USING btree (tp_codigo, pro_codigo);
CREATE UNIQUE INDEX idx_tabela_preco ON public.tabela_preco_produtos USING btree (tp_codigo, pro_codigo);