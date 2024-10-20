CREATE TABLE public."role" (
	id uuid NULL,
	"name" varchar NULL,
	"scope" varchar NULL,
	created_at_ut0 int8 NULL,
	created_by varchar NULL,
	updated_at_utc0 int8 NULL,
	updated_by varchar NULL,
	CONSTRAINT role_pk PRIMARY KEY (id)
);
