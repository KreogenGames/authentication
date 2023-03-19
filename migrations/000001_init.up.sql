CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL,
    "hashed_pass" varchar(30),
    "lastName" varchar(50),
    "firstName" varchar(50),
    "middleName" varchar(50),
    "phoneNumber" TEXT,
	"role" INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

alter table users owner to postgres;

insert into public.users (email) values ('ertek.h.i@edu.mirea.ru');
insert into public.users (email) values ('sibgatulov@mirea.ru');