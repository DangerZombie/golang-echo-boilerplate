CREATE TABLE public.teacher (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	job_title_id uuid NULL,
	status varchar NULL,
	experience int NULL,
	"degree" varchar NULL,
	created_at_utc0 int8 NULL,
	created_by varchar NULL,
	updated_at_utc0 int8 NULL,
	updated_by varchar NULL,
	CONSTRAINT teacher_pk PRIMARY KEY (id),
	CONSTRAINT teacher_user_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id),
	CONSTRAINT teacher_job_title_job_title_id_fk FOREIGN KEY (job_title_id) REFERENCES public.job_title(id)
);