CREATE TABLE
    "event_tag" (
        "event_id" UUID NOT NULL,
        "tag_id" UUID NOT NULL,
        CONSTRAINT "fk_event_tag_event" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE,
        CONSTRAINT "fk_event_tag_tag" FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE,
        CONSTRAINT "event_tag_pk" PRIMARY KEY ("event_id", "tag_id")
    );