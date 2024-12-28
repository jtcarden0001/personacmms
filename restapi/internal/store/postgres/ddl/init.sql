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
  manufacturer varchar,
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
  instructions varchar,
  asset_id uuid NOT NULL REFERENCES asset(id),
  PRIMARY KEY (id)
);

CREATE TABLE workorderstatus (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE workorder (
  id uuid,
  title varchar NOT NULL,
  created_date timestamptz NOT NULL,
  completed_date timestamptz,
  instructions varchar,
  notes varchar,
  cumulative_miles int, 
  cumulative_hours int,
  asset_id uuid NOT NULL REFERENCES asset(id),
  status_id uuid NOT NULL REFERENCES workorderstatus(id),
  PRIMARY KEY (id)
);

CREATE TABLE task_workorder (
  workorder_id uuid REFERENCES workorder(id),
  task_id uuid REFERENCES task(id),
  PRIMARY KEY (workorder_id)
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

CREATE TABLE timeunit (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE timetrigger (
  id uuid,
  quantity int NOT NULL,
  timeunit_id uuid NOT NULL REFERENCES timeunit(id), -- managed by application
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);

CREATE TABLE usageunit (
  id uuid,
  title varchar NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE usagetrigger (
  id uuid,
  quantity int NOT NULL,
  usageunit_id uuid NOT NULL REFERENCES usageunit(id), -- managed by application
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);

CREATE TABLE datetrigger (
  id uuid,
  scheduled_date timestamptz NOT NULL,
  task_id uuid NOT NULL REFERENCES task(id),
  PRIMARY KEY (id)
);

------ seed data

-- timeunit
INSERT INTO timeunit (id, title) VALUES ('AA199071-058D-4942-A8CD-77103F88332B', 'day');
INSERT INTO timeunit (id, title) VALUES ('BA199071-058D-4942-A8CD-77103F88332B', 'week');
INSERT INTO timeunit (id, title) VALUES ('CA199071-058D-4942-A8CD-77103F88332B', 'month');
INSERT INTO timeunit (id, title) VALUES ('DA199071-058D-4942-A8CD-77103F88332B', 'year');

-- usageunit
INSERT INTO usageunit (id, title) VALUES ('AACD1682-1FA4-4C75-A870-CF953B8859B9', 'hour');
INSERT INTO usageunit (id, title) VALUES ('BACD1682-1FA4-4C75-A870-CF953B8859B9', 'day');
INSERT INTO usageunit (id, title) VALUES ('CACD1682-1FA4-4C75-A870-CF953B8859B9', 'mile');

-- workorderstatus
INSERT INTO workorderstatus (id, title) VALUES ('AB046BAE-A286-4A4E-ABB5-216C756BF7F9', 'new');
INSERT INTO workorderstatus (id, title) VALUES ('BB046BAE-A286-4A4E-ABB5-216C756BF7F9', 'in progress');
INSERT INTO workorderstatus (id, title) VALUES ('CB046BAE-A286-4A4E-ABB5-216C756BF7F9', 'complete');
INSERT INTO workorderstatus (id, title) VALUES ('DB046BAE-A286-4A4E-ABB5-216C756BF7F9', 'cancelled');


