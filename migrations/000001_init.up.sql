CREATE TABLE IF NOT EXISTS "roles" (
    "id" SERIAL,
    "role_name" varchar(20) NOT NULL UNIQUE,
    "access_level" INTEGER NOT NULL,
    CONSTRAINT "roles_pkey" PRIMARY KEY ("id")
);

alter table roles owner to postgres;

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL,
    "email" TEXT NOT NULL UNIQUE,
    "hashed_pass" varchar(255) UNIQUE,
    "last_name" varchar(50),
    "first_name" varchar(50),
    "middle_name" varchar(50),
    "phone_number" TEXT,
	"role" INTEGER NOT NULL REFERENCES roles (id),
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

alter table users owner to postgres;

-- create unique index users_id_uindex on users (id);

CREATE TABLE IF NOT EXISTS "grades" (
    "id" SERIAL,
    "teacher_id" INTEGER NOT NULL REFERENCES users (id),
    "discipline" TEXT NOT NULL, -- В будущем создать таблицу под дисциплины
    "student_id" INTEGER NOT NULL REFERENCES users (id),
    "grade" INTEGER CHECK(grade >= 0 AND grade <= 5),
    CONSTRAINT "grades_pkey" PRIMARY KEY ("id")
);

alter table grades owner to postgres;

-- create unique index grades_id_uindex on grades (id)

insert into public.roles (role_name, access_level) values ('user', 0);
insert into public.roles (role_name, access_level) values ('admin', 10);
insert into public.roles (role_name, access_level) values ('student', 1);

insert into public.users (email, role) values ('sibgatulov@gmail.com', 2);
insert into public.users (email, role) values ('ertek.h.i@edu.mirea.ru', 3);

-- insert into public.grades (teacher_id, discipline, student_id, grade) values (3, 'ОСТ', 2, 3);