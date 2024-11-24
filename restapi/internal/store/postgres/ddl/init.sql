CREATE TABLE category (
  title varchar,
  id uuid NOT NULL UNIQUE,
  description varchar,
  PRIMARY KEY (title)
);

CREATE TABLE assetgroup (
  title varchar,
  id uuid NOT NULL UNIQUE,
  PRIMARY KEY (title)
);

CREATE TABLE asset (
  group_title varchar REFERENCES assetgroup(title),
  title varchar,
  id uuid NOT NULL UNIQUE,
  year int,
  make varchar,
  model_number varchar,
  serial_number varchar,
  description varchar,
  category_title varchar REFERENCES category(title),
  PRIMARY KEY (group_title, title)
);

CREATE TABLE task (
  title varchar,
  id uuid NOT NULL UNIQUE,
  description varchar,
  type int,
  PRIMARY KEY (title)
);

CREATE TABLE asset_task (
  id uuid,
  title varchar,
  unique_instructions varchar,
  asset_id uuid NOT NULL REFERENCES asset(id),
  task_id uuid, 
  PRIMARY KEY (id)
);

CREATE TABLE tool (
  title varchar,
  id uuid NOT NULL UNIQUE,
  size varchar,
  PRIMARY KEY (title)
);

CREATE TABLE assettask_tool (
  assettask_id uuid NOT NULL REFERENCES asset_task(id),
  tool_id uuid NOT NULL REFERENCES tool(id),
  PRIMARY KEY (assettask_id, tool_id)
);

CREATE TABLE consumable (
  title varchar,
  id uuid NOT NULL UNIQUE,
  PRIMARY KEY (title)
);

CREATE TABLE assettask_consumable (
  assettask_id uuid NOT NULL REFERENCES asset_task(id),
  consumable_id uuid NOT NULL REFERENCES consumable(id),
  quantity_note varchar NOT NULL,
  PRIMARY KEY (assettask_id, consumable_id)
);

CREATE TABLE timeunit (
  title varchar,
  id uuid NOT NULL UNIQUE,
  PRIMARY KEY (title)
);

CREATE TABLE timetrigger (
  id uuid,
  quanitity int NOT NULL,
  timeunit_title varchar NOT NULL REFERENCES timeunit(title),
  PRIMARY KEY (id)
);

CREATE TABLE assettask_timetrigger (
  assettask_id uuid NOT NULL REFERENCES asset_task(id),
  timetrigger_id uuid NOT NULL REFERENCES timetrigger(id)
);

CREATE TABLE usageunit (
  title varchar,
  id uuid NOT NULL UNIQUE,
  PRIMARY KEY (title)
);

CREATE TABLE usagetrigger (
  id uuid,
  quanitity int NOT NULL,
  usageunit_title varchar NOT NULL REFERENCES usageunit(title),
  PRIMARY KEY (id)
);

CREATE TABLE assettask_usagetrigger (
  asset_task_id uuid NOT NULL REFERENCES asset_task(id),
  usagetrigger_id uuid NOT NULL REFERENCES usagetrigger(id)
);

CREATE TABLE workorderstatus (
  title varchar,
  id uuid NOT NULL UNIQUE,
  PRIMARY KEY (title)
);

CREATE TABLE workorder (
  id uuid,
  created_date timestamptz NOT NULL,
  completed_date timestamptz,
  notes varchar,
  cumulative_mileage int, 
  cumulative_hours int,
  assettask_id uuid NOT NULL REFERENCES asset_task(id),
  status_title varchar NOT NULL REFERENCES workorderstatus(title),
  PRIMARY KEY (id)
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



