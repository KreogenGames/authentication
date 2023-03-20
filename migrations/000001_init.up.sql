CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "hashed_pass" varchar(30) UNIQUE,
    "lastName" varchar(50),
    "firstName" varchar(50),
    "middleName" varchar(50),
    "phoneNumber" TEXT,
	"role" INTEGER NOT NULL DEFAULT 0, -- В будущем создать таблицу под роли
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

insert into public.users (email, role) values ('sibgatulov@gmail.com', 5);
insert into public.users (email) values ('ertek.h.i@edu.mirea.ru');
insert into public.users (email) values ('sibgatulov@mirea.ru');
insert into public.users (email) values ('monakov.a.v@edu.mirea.ru');

insert into public.grades (teacher_id, discipline, student_id, grade) values (3, 'ОСТ', 2, 3);