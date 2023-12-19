CREATE TABLE equipment_category (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE equipment (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  year INT NOT NULL,
  make VARCHAR NOT NULL,
  model_number VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  category_id INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_equipment__equipment_category
    FOREIGN KEY (category_id)
        REFERENCES equipment_category(id)
);

CREATE TABLE usage_periodicity_unit (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE time_periodicity_unit (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE task (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  instructions VARCHAR NOT NULL,
  time_periodicity_quantity INT,
  time_periodicity_unit_id INT,
  usage_periodicity_quantity INT,
  usage_periodicity_unit_id INT,
  PRIMARY KEY (id),
  CONSTRAINT fk_task__time_periodicity_unit_id
    FOREIGN KEY (time_periodicity_unit_id)
        REFERENCES time_periodicity_unit(id)
        ON DELETE SET NULL,
  CONSTRAINT fk_task__usage_periodicity_unit_id
    FOREIGN KEY (usage_periodicity_unit_id)
        REFERENCES usage_periodicity_unit(id)
        ON DELETE SET NULL
);

CREATE TABLE equipment_task (
  equipment_id INT NOT NULL,
  task_id INT NOT NULL,
  CONSTRAINT fk_equipment_task__equipment_id
    FOREIGN KEY (equipment_id)
      REFERENCES equipment(id)
      ON DELETE CASCADE,
  CONSTRAINT fk_equipment_task__task_id
    FOREIGN KEY (task_id)
      REFERENCES task(id)
      ON DELETE CASCADE
);

CREATE TABLE work_order_status (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE work_order (
  id INT GENERATED ALWAYS AS IDENTITY,
  task_id INT NOT NULL,
  status_id INT NOT NULL,
  create_date TIMESTAMPTZ NOT NULL,
  complete_date TIMESTAMPTZ,
  PRIMARY KEY (id),
  CONSTRAINT fk_work_order__task_id
    FOREIGN KEY (task_id)
      REFERENCES task(id)
      ON DELETE CASCADE,
  CONSTRAINT fk_work_order__status_id
    FOREIGN KEY (status_id)
      REFERENCES work_order_status(id)
);

CREATE TABLE tool (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  size VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE task_tool (
  task_id INT NOT NULL,
  tool_id INT NOT NULL,
  CONSTRAINT fk_task_tool__task_id
    FOREIGN KEY (task_id)
      REFERENCES task(id)
      ON DELETE CASCADE,
  CONSTRAINT fk_task_tool__tool_id
    FOREIGN KEY (tool_id)
      REFERENCES tool(id)
);

CREATE TABLE consumable (
  id INT GENERATED ALWAYS AS IDENTITY,
  title VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE task_consumable (
  task_id INT NOT NULL,
  consumable_id INT NOT NULL,
  quantity_note VARCHAR NOT NULL,
  CONSTRAINT fk_task_consumable__task_id
    FOREIGN KEY (task_id)
      REFERENCES task(id)
      ON DELETE CASCADE,
  CONSTRAINT fk_task_consumable__consumable_id
    FOREIGN KEY (consumable_id)
      REFERENCES consumable(id)
);

---- in case you want to wipe the db schema
DROP TABLE equipment_task;
DROP TABLE equipment;
DROP TABLE equipment_category;
DROP TABLE task_tool;
DROP TABLE tool;
DROP TABLE task_consumable;
DROP TABLE consumable;
DROP TABLE work_order;
DROP TABLE work_order_status;
DROP TABLE task;
DROP TABLE usage_periodicity_unit;
DROP TABLE time_periodicity_unit;



