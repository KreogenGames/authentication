CREATE TABLE IF NOT EXISTS "roles" (
    "id" INTEGER NOT NULL UNIQUE,
    "role_name" varchar(20) NOT NULL UNIQUE,
    "access_level" INTEGER NOT NULL,
    CONSTRAINT "roles_pkey" PRIMARY KEY ("id")
);

alter table roles owner to postgres;

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "hashed_pass" varchar(30) UNIQUE,
    "lastName" varchar(50),
    "firstName" varchar(50),
    "middleName" varchar(50),
    "phoneNumber" TEXT,
	"role" INTEGER NOT NULL REFERENCES roles (id) DEFAULT 0,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

alter table users owner to postgres;

-- create unique index users_id_uindex on users (id);

CREATE TABLE IF NOT EXISTS "grades" (
    "id" SERIAL NOT NULL,
    "teacher_id" INTEGER NOT NULL REFERENCES users (id),
    "discipline" TEXT NOT NULL, -- В будущем создать таблицу под дисциплины
    "student_id" INTEGER NOT NULL REFERENCES users (id),
    "grade" INTEGER CHECK(grade >= 0 AND grade <= 5),
    CONSTRAINT "grades_pkey" PRIMARY KEY ("id")
);

alter table grades owner to postgres;

-- create unique index grades_id_uindex on grades (id)

insert into public.roles (id, role_name, access_level) values (0, 'user', 0);
insert into public.roles (id, role_name, access_level) values (1, 'admin', 10);
insert into public.roles (id, role_name, access_level) values (2, 'student', 1);

insert into public.users (email, role) values ('sibgatulov@gmail.com', 1);
insert into public.users (email, role) values ('ertek.h.i@edu.mirea.ru', 2);
insert into public.users (email) values ('sibgatulov@mirea.ru');
insert into public.users (email) values ('monakov.a.v@edu.mirea.ru');

insert into public.grades (teacher_id, discipline, student_id, grade) values (3, 'ОСТ', 2, 3);