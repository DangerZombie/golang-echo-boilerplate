CREATE TABLE public."user" (
	id uuid NOT NULL,
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	status varchar NOT NULL,
	nickname varchar NOT NULL,
	created_at_utc0 int8 NULL,
	created_by varchar NULL,
	updated_at_utc0 int8 NULL,
	updated_by varchar NULL,
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_unique UNIQUE (username)
);