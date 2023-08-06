CREATE TABLE
    "tags" (
        "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
        "name" VARCHAR(120) NOT NULL,
        CONSTRAINT "tags_pkey" PRIMARY KEY ("id")
    );