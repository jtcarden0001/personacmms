CREATE TABLE "category" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  "description" varchar,
  PRIMARY KEY ("title")
  );

CREATE TABLE "assetgroup" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  PRIMARY KEY ("title")
  );

CREATE TABLE "asset" (
  "group_title" varchar,
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  "year" int,
  "make" varchar,
  "model_number" varchar,
  "description" varchar,
  "category_title" varchar,
  PRIMARY KEY ("group_title", "title"),
  CONSTRAINT "FK_asset.group_title"
    FOREIGN KEY ("group_title")
      REFERENCES "assetgroup"("title"),
  CONSTRAINT "FK_asset.category_title"
    FOREIGN KEY ("category_title")
      REFERENCES "category"("title")
);

CREATE TABLE "preventativetask" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  "description" varchar,
  PRIMARY KEY ("title")
);

CREATE TABLE "correctivetask" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  "description" varchar,
  PRIMARY KEY ("title")
);

CREATE TABLE "asset_task" (
  "id" uuid,
  "unique_instructions" varchar,
  "asset_id" uuid NOT NULL,
  "preventativetask_id" uuid, 
  "correctivetask_id" uuid,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_asset_task.asset_id"
    FOREIGN KEY ("asset_id")
      REFERENCES "asset"("id"),
  CONSTRAINT "FK_asset_task.preventativetask_id"
    FOREIGN KEY ("preventativetask_id")
      REFERENCES "preventativetask"("id"),
  CONSTRAINT "FK_asset_task.correctivetask_id"
    FOREIGN KEY ("correctivetask_id")
      REFERENCES "correctivetask"("id")
);

CREATE TABLE "tool" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  "size" varchar,
  PRIMARY KEY ("title")
);

CREATE TABLE "assettask_tool" (
  "assettask_id" uuid NOT NULL,
  "tool_id" uuid NOT NULL,
  PRIMARY KEY ("assettask_id", "tool_id"),
  CONSTRAINT "FK_assettask_tool.asset_task_id"
    FOREIGN KEY ("assettask_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_tool.tool_id"
    FOREIGN KEY ("tool_id")
      REFERENCES "tool"("id")
);

CREATE TABLE "consumable" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  PRIMARY KEY ("title")
);

CREATE TABLE "assettask_consumable" (
  "assettask_id" uuid NOT NULL,
  "consumable_id" uuid NOT NULL,
  "quantity_note" varchar NOT NULL,
  PRIMARY KEY ("assettask_id", "consumable_id"),
  CONSTRAINT "FK_assettask_consumable.asset_task_id"
    FOREIGN KEY ("assettask_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_consumable.consumable_id"
    FOREIGN KEY ("consumable_id")
      REFERENCES "consumable"("id")
);

CREATE TABLE "timeunit" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  PRIMARY KEY ("title")
);

CREATE TABLE "timetrigger" (
  "id" uuid,
  "quanitity" int NOT NULL,
  "timeunit_title" varchar NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_timetrigger.timeunit_title"
    FOREIGN KEY ("timeunit_title")
      REFERENCES "timeunit"("title")
);

CREATE TABLE "assettask_timetrigger" (
  "assettask_id" uuid NOT NULL,
  "timetrigger_id" uuid NOT NULL,
  CONSTRAINT "FK_assettask_timetrigger.asset_task_id"
    FOREIGN KEY ("assettask_id")
      REFERENCES "asset_task"("id"),
  CONSTRAINT "FK_assettask_timetrigger.timetrigger_id"
    FOREIGN KEY ("timetrigger_id")
      REFERENCES "timetrigger"("id")
);

CREATE TABLE "usageunit" (
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  PRIMARY KEY ("title")
);

CREATE TABLE "usagetrigger" (
  "id" uuid,
  "quanitity" int NOT NULL,
  "usageunit_title" varchar NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_usagetrigger.usageunit_title"
    FOREIGN KEY ("usageunit_title")
      REFERENCES "usageunit"("title")
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
  "title" varchar,
  "id" uuid NOT NULL UNIQUE,
  PRIMARY KEY ("title")
);

CREATE TABLE "workorder" (
  "id" uuid,
  "created_date" timestamptz NOT NULL,
  "completed_date" timestamptz,
  "notes" varchar,
  "cumulative_mileage" int, 
  "cumulative_hours" int,
  "assettask_id" uuid NOT NULL,
  "status_title" varchar NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_workorder.status_title"
    FOREIGN KEY ("status_title")
      REFERENCES "workorderstatus"("title"),
  CONSTRAINT "FK_workorder.assettask_id"
    FOREIGN KEY ("assettask_id")
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



