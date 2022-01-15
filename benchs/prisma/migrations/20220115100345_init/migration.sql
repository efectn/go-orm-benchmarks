-- CreateTable
CREATE TABLE "ModelsPrisma" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "fax" TEXT NOT NULL,
    "web" TEXT NOT NULL,
    "age" INTEGER NOT NULL,
    "right" BOOLEAN NOT NULL,
    "counter" BIGINT NOT NULL,

    PRIMARY KEY ("id")
);
