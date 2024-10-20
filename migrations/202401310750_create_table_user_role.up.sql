CREATE TABLE public.user_role (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	role_id uuid NOT NULL,
	created_at_utc0 int8 NULL,
	created_by varchar NULL,
	updated_at_utc0 int8 NULL,
	updated_by varchar NULL,
	CONSTRAINT newtable_pk PRIMARY KEY (id),
	CONSTRAINT user_role_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id),
	CONSTRAINT user_role_role_id_fk FOREIGN KEY (role_id) REFERENCES public."role"(id)
);
