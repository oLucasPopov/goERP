CREATE TABLE CIDADES (
	ID SERIAL4 NOT NULL,
	ID_IBGE INT4 NULL,
	ID_ESTADO INT4 NULL,
	CIDADE VARCHAR(64) NULL,
	CONSTRAINT CIDADES_ID_KEY UNIQUE (ID),
  CONSTRAINT CIDADES_ID_ESTADO_FKEY FOREIGN KEY (ID_ESTADO) REFERENCES ESTADOS(ID)
);