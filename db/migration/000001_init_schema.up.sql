CREATE TABLE "persons" ( 
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL
);

CREATE INDEX ON "persons" ("name");