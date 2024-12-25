CREATE TABLE category (
  id uuid,
  title varchar NOT NULL UNIQUE,
  description varchar,
  PRIMARY KEY (id)
);

CREATE TABLE agroup (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE asset (
  id uuid,
  title varchar NOT NULL UNIQUE,
  year int,
  make varchar,
  model_number varchar,
  serial_number varchar,
  description varchar,
  PRIMARY KEY (id)
);

CREATE TABLE category_asset (
  category_id uuid REFERENCES category(id),
  asset_id uuid REFERENCES asset(id),
  PRIMARY KEY (category_id, asset_id)
);

CREATE TABLE agroup_asset (
  agroup_id uuid REFERENCES agroup(id),
  asset_id uuid REFERENCES asset(id),
  PRIMARY KEY (agroup_id, asset_id)
);

CREATE TABLE task (
  id uuid,
  title varchar NOT NULL,
  unique_instructions varchar,
  PRIMARY KEY (id)
);

CREATE TABLE workorder (
  id uuid,
  title varchar NOT NULL,
  created_date timestamptz NOT NULL,
  completed_date timestamptz,
  notes varchar,
  cumulative_miles int, 
  cumulative_hours int,
  PRIMARY KEY (id)
);

CREATE TABLE tool (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE consumable (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE task_tool (
  task_id uuid REFERENCES task(id),
  tool_id uuid REFERENCES tool(id),
  size_note varchar,
  PRIMARY KEY (task_id, tool_id, size_note)
);

CREATE TABLE task_consumable (
  task_id uuid REFERENCES task(id),
  consumable_id uuid REFERENCES consumable(id),
  quantity_note varchar NOT NULL,
  PRIMARY KEY (task_id, consumable_id)
);

CREATE TABLE workorder_tool (
  workorder_id uuid REFERENCES workorder(id),
  tool_id uuid REFERENCES tool(id),
  size_note varchar,
  PRIMARY KEY (workorder_id, tool_id, size_note)
);

CREATE TABLE workorder_consumable (
  workorder_id uuid REFERENCES workorder(id),
  consumable_id uuid REFERENCES consumable(id),
  quantity_note varchar NOT NULL,
  PRIMARY KEY (workorder_id, consumable_id)
);

CREATE TABLE timetrigger (
  id uuid,
  quantity int NOT NULL,
  timeunit varchar NOT NULL, -- managed by application
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);

CREATE TABLE usagetrigger (
  id uuid,
  quantity int NOT NULL,
  usageunit varchar NOT NULL, -- managed by application
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);

CREATE TABLE datetrigger (
  id uuid,
  scheduled_date timestamptz NOT NULL,
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);


