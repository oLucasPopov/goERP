CREATE TABLE itens_venda (
	id serial4 NOT NULL,
	ven_codigo int4 NOT NULL,
	pro_codigo int4 NOT NULL,
	quantidade numeric(18, 3) NOT NULL,
	valor_unitario numeric(18, 2) NOT NULL,
	valor_total_bruto numeric(18, 4) NULL,
	valor_total_liquido numeric(18, 4) NULL,
	desconto_unitario numeric(18, 4) NOT NULL DEFAULT 0,
	tipo_unidade varchar(2) NOT NULL,
	diferenca_valor_unt_preco_tabela numeric(18, 2) NULL,
	quantidade_expedida numeric(18, 3) NOT NULL DEFAULT 0,
	desconto_expedicao numeric(18, 4) NOT NULL DEFAULT 0,
	desconto_total numeric(18, 4) NULL GENERATED ALWAYS AS ((quantidade * desconto_unitario + desconto_expedicao)) STORED,
	CONSTRAINT itens_venda_id_key PRIMARY KEY (id),
	CONSTRAINT itens_venda_pro_codigo_fkey FOREIGN KEY (pro_codigo) REFERENCES produtos(id),
	CONSTRAINT itens_venda_ven_codigo_fkey FOREIGN KEY (ven_codigo) REFERENCES vendas(id)
);