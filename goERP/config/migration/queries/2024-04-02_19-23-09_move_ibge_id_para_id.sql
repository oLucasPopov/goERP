ALTER TABLE public.empresas drop CONSTRAINT empresas_id_cidade_fkey;
ALTER TABLE public.empresas drop CONSTRAINT empresas_id_estado_fkey;
ALTER TABLE public.cidades drop CONSTRAINT cidades_id_estado_fkey;

update clientes
set id_cidade = (select id_ibge from cidades where id = id_cidade)
  ,id_estado = (select id_ibge from estados where id = id_estado)
where id_cidade is not null
   or id_estado is not null;

update cidades
set id_estado = (select id_ibge from estados where id = id_estado)
where id_estado is not null;

alter table estados add column temp_id serial4;
update estados set temp_id = id_ibge;
ALTER TABLE estados DROP COLUMN id;
alter table estados rename column temp_id to id;
alter table estados drop column id_ibge;

alter table cidades add column temp_id serial4;
update cidades set temp_id = id_ibge ;
ALTER TABLE cidades DROP COLUMN id;
alter table cidades rename column temp_id to id;
alter table  cidades drop column id_ibge;

ALTER TABLE public.estados ADD CONSTRAINT estados_id_key UNIQUE (id);
ALTER TABLE public.cidades ADD CONSTRAINT cidades_id_key UNIQUE (id);


ALTER TABLE public.clientes ADD CONSTRAINT fk_estado_id FOREIGN KEY (id_estado) REFERENCES estados(id);

ALTER TABLE public.clientes ADD CONSTRAINT fk_cidade_id FOREIGN KEY (id_cidade) REFERENCES cidades(id);

ALTER TABLE public.empresas ADD CONSTRAINT empresas_id_cidade_fkey FOREIGN KEY (id_cidade) REFERENCES cidades(id);

ALTER TABLE public.empresas ADD CONSTRAINT empresas_id_estado_fkey FOREIGN KEY (id_estado) REFERENCES estados(id);

ALTER TABLE public.cidades ADD CONSTRAINT cidades_id_estado_fkey FOREIGN KEY (id_estado) REFERENCES estados(id);