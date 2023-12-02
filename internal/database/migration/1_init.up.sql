CREATE TYPE PLAYBACK_ORDER_TYPE AS ENUM (
    'IN_ORDER',
    'REPLAY',
    'RANDOM'
);

CREATE TABLE "queue_items"
(
    "id"           SERIAL PRIMARY KEY,
    "song_id"      INTEGER NOT NULL,
    "prev_item_id" INTEGER UNIQUE,
    "next_item_id" INTEGER UNIQUE,
    FOREIGN KEY ("prev_item_id") REFERENCES "queue_items" ("id"),
    FOREIGN KEY ("next_item_id") REFERENCES "queue_items" ("id")
);

CREATE TABLE "rooms"
(
    "id"                    SERIAL PRIMARY KEY,
    "owner_id"              INTEGER             NOT NULL,
    "current_queue_item_id" INTEGER UNIQUE,
    "name"                  TEXT                NOT NULL,
    "playback_order_type"   PLAYBACK_ORDER_TYPE NOT NULL,
    FOREIGN KEY ("current_queue_item_id") REFERENCES "queue_items" ("id")
);

CREATE TABLE "share_codes"
(
    "id" SERIAL PRIMARY KEY,
    "room_id" INTEGER UNIQUE NOT NULL,
    "code" CHAR(32) UNIQUE NOT NULL,
    FOREIGN KEY ("room_id") REFERENCES "rooms" ("id")
);

CREATE TABLE "roommates"
(
    "id"      SERIAL PRIMARY KEY,
    "account_id" INTEGER NOT NULL,
    "room_id" INTEGER NOT NULL,
    FOREIGN KEY ("room_id") REFERENCES "rooms" ("id")
);

CREATE TABLE "sessions"
(
    "id"        SERIAL PRIMARY KEY,
    "device_id" INTEGER UNIQUE NOT NULL,
    "room_id"   INTEGER        NOT NULL,
    FOREIGN KEY ("room_id") REFERENCES "rooms" ("id")
);
