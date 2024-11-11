---- in case you want to wipe the db schema
DROP TABLE workorder;
DROP TABLE workorderstatus;
DROP TABLE assettask_usagetrigger;
DROP TABLE usagetrigger;
DROP TABLE usageunit;
DROP TABLE assettask_timetrigger;
DROP TABLE timetrigger;
DROP TABLE timeunit;
DROP TABLE assettask_consumable;
DROP TABLE consumable;
DROP TABLE assettask_tool;
DROP TABLE tool;
DROP TABLE asset_task;
DROP TABLE task;
DROP TABLE asset;
DROP TABLE assetgroup;
DROP TABLE category;

CREATE TABLE "category" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "assetgroup" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "asset" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  "year" int,
  "make" varchar,
  "model_number" varchar,
  "description" varchar,
  "category_id" uuid,
  "group_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_asset.group_id"
    FOREIGN KEY ("group_id")
      REFERENCES "assetgroup"("id"),
  CONSTRAINT "FK_asset.category_id"
    FOREIGN KEY ("category_id")
      REFERENCES "category"("id")
);

CREATE TABLE "task" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  PRIMARY KEY ("id")
);

CREATE TABLE "asset_task" (
  "id" uuid NOT NULL,
  "unique_instructions" varchar,
  "asset_id" uuid NOT NULL,
  "task_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_asset_task.asset_id"
    FOREIGN KEY ("asset_id")
      REFERENCES "asset"("id"),
  CONSTRAINT "FK_asset_task.task_id"
    FOREIGN KEY ("task_id")
      REFERENCES "task"("id")
);

CREATE TABLE "tool" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  "size" varchar,
  PRIMARY KEY ("id")
);

CREATE TABLE "assettask_tool" (
  "asset_task_id" uuid NOT NULL,
  "tool_id" uuid NOT NULL,
  CONSTRAINT "FK_assettask_tool.asset_task_id"
    FOREIGN KEY ("asset_task_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_tool.tool_id"
    FOREIGN KEY ("tool_id")
      REFERENCES "tool"("id")
);

CREATE TABLE "consumable" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "assettask_consumable" (
  "quantity_note" varchar NOT NULL,
  "asset_task_id" uuid NOT NULL,
  "consumable_id" uuid NOT NULL,
  CONSTRAINT "FK_assettask_consumable.asset_task_id"
    FOREIGN KEY ("asset_task_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_consumable.consumable_id"
    FOREIGN KEY ("consumable_id")
      REFERENCES "consumable"("id")
);

CREATE TABLE "timeunit" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "timetrigger" (
  "id" uuid NOT NULL,
  "quanitity" int NOT NULL,
  "timeunit_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_timetrigger.timeunit_id"
    FOREIGN KEY ("timeunit_id")
      REFERENCES "timeunit"("id")
);

CREATE TABLE "assettask_timetrigger" (
  "asset_task_id" uuid NOT NULL,
  "timetrigger_id" uuid NOT NULL,
  CONSTRAINT "FK_assettask_timetrigger.asset_task_id"
    FOREIGN KEY ("asset_task_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_timetrigger.timetrigger_id"
    FOREIGN KEY ("timetrigger_id")
      REFERENCES "timetrigger"("id")
);

CREATE TABLE "usageunit" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "usagetrigger" (
  "id" uuid NOT NULL,
  "quanitity" int NOT NULL,
  "usageunit_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_usagetrigger.usageunit_id"
    FOREIGN KEY ("usageunit_id")
      REFERENCES "usageunit"("id")
);

CREATE TABLE "assettask_usagetrigger" (
  "asset_task_id" uuid NOT NULL,
  "usagetrigger_id" uuid NOT NULL,
  CONSTRAINT "FK_assettask_usagetrigger.asset_task_id"
    FOREIGN KEY ("asset_task_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_usagetrigger.usagetrigger_id"
    FOREIGN KEY ("usagetrigger_id")
      REFERENCES "usagetrigger"("id")
);

CREATE TABLE "workorderstatus" (
  "id" uuid NOT NULL,
  "title" varchar NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "workorder" (
  "id" uuid NOT NULL,
  "created_date" timestamptz NOT NULL,
  "completed_date" timestamptz,
  "task_id" uuid NOT NULL,
  "status_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_workorder.status_id"
    FOREIGN KEY ("status_id")
      REFERENCES "workorderstatus"("id"),
  CONSTRAINT "FK_workorder.task_id"
    FOREIGN KEY ("task_id")
      REFERENCES "asset_task"("id")
);

/* static data not modified by app */

/* TODO: make this some kind of enum that is sourced from one place */
INSERT INTO timeunit (id, title) VALUES ('4137d18f-d548-4d75-b84b-5a92d36acbc7', 'Day');
INSERT INTO timeunit (id, title) VALUES ('8deb9996-d0a2-4c3d-b551-de07f9d4b91b','Week');
INSERT INTO timeunit (id, title) VALUES ('e85451ba-f3e2-41a2-85e3-97aa782e76b8','Month');
INSERT INTO timeunit (id, title) VALUES ('f14475db-e969-4fb6-af55-341110e29df6','Year');

/* TODO: make this some kind of enum that is sourced from one place */
INSERT INTO workorderstatus (id, title) VALUES ('a0f7f76f-9b1b-4758-83d9-bc42da07fbe7','New');
INSERT INTO workorderstatus (id, title) VALUES ('ff0ff8df-05fc-48b1-a5cf-d7af8b41cde4','In Progress');
INSERT INTO workorderstatus (id, title) VALUES ('72d51235-c1c5-4387-a94f-efebd40d9eed','Complete');



