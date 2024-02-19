CREATE TABLE public.job_title (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	created_at_utc0 int8 NULL,
	created_by varchar NULL,
	updated_at_utc0 int8 NULL,
	updated_by varchar NULL,
	CONSTRAINT job_title_pk PRIMARY KEY (id)
);
