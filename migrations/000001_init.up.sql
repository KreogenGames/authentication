CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT,
    "lastName" TEXT,
    "firstName" TEXT,
    "middleName" TEXT,
    "phoneNumber" TEXT,
	"role" INTEGER NOT NULL DEFAULT 0;

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);