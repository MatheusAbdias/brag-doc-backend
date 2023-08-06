CREATE TABLE
    "events"(
        "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
        "name" VARCHAR(120) NOT NULL,
        "description" VARCHAR(255) NOT NULL,
        "date" DATE NOT NULL,
        CONSTRAINT "events_pkey" PRIMARY KEY ("id")
    );