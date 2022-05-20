#!/bin/sh
psql -U $POSTGRES_USER -d $POSTGRES_DB <<EOF
  CREATE TABLE users (
      user_id         BIGINT       NOT NULL PRIMARY KEY,
      username        VARCHAR(255) NOT NULL,
      correct_answers INTEGER      NOT NULL DEFAULT 0,
      total_answers   INTEGER      NOT NULL DEFAULT 0,
      is_passing      BOOLEAN      NOT NULL DEFAULT FALSE,
      registered_at   TIMESTAMP    NOT NULL DEFAULT now()
  );
EOF