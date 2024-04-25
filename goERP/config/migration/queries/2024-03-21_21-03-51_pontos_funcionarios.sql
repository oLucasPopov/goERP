CREATE TABLE pontos_funcionarios (
	id serial4 NOT NULL,
	id_funcionario int4 NOT NULL,
	data_hora_entrada timestamp NOT NULL,
	data_hora_saida timestamp NULL,
	salario_funcionario numeric(18, 2) NULL,
	pago bool NOT NULL DEFAULT false,
	CONSTRAINT pontos_funcionarios_pkey PRIMARY KEY (id)
);