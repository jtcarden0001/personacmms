-- Assumption is this is run on a clean database

INSERT INTO equipment_category (title) VALUES ('Test Category 1');
INSERT INTO equipment_category (title) VALUES ('Test Category 2');

INSERT INTO equipment (title, year, make, model_number, description, category_id) VALUES ('Test Equipment 1', 2020, 'Test Make 1', 'Test Model Number 1', 'Test Description 1', 1);
INSERT INTO equipment (title, year, make, model_number, description, category_id) VALUES ('Test Equipment 2', 2020, 'Test Make 2', 'Test Model Number 2', 'Test Description 2', 2);

INSERT INTO usage_periodicity_unit (title) VALUES ('Miles');
INSERT INTO usage_periodicity_unit (title) VALUES ('Hours');

INSERT INTO time_periodicity_unit (title) VALUES ('Days');
INSERT INTO time_periodicity_unit (title) VALUES ('Weeks');

INSERT INTO task (title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id) VALUES ('Test Task 1', 'Test Instructions 1', 1, 1, 1, 1, 1);
INSERT INTO task (title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, equipment_id) VALUES ('Test Task 2', 'Test Instructions 2', 2, 2, 2, 2, 2);

INSERT INTO work_order_status (title) VALUES ('New');
INSERT INTO work_order_status (title) VALUES ('In Progress');
INSERT INTO work_order_status (title) VALUES ('Complete');

INSERT INTO work_order (task_id, status_id, create_date, complete_date) VALUES (1, 1, '2020-01-01', '2020-01-02');
INSERT INTO work_order (task_id, status_id, create_date) VALUES (2, 2, '2020-02-01');

INSERT INTO tool (title, size) VALUES ('Test Tool 1', 'Test Size 1');

INSERT INTO task_tool (task_id, tool_id) VALUES (1, 1);

INSERT INTO consumable (title) VALUES ('Test Consumable 1');

INSERT INTO task_consumable (task_id, consumable_id, quantity_note) VALUES (1, 1, '1');