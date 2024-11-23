-- Assumption is this is run on a clean database

INSERT INTO category (title) VALUES ('seed-get-cat');
INSERT INTO category (title) VALUES ('seed-delete-cat');

/*
INSERT INTO asset_category (title) VALUES ('Test Category 1');
INSERT INTO asset_category (title) VALUES ('Test Category 2');

INSERT INTO asset (title, year, make, model_number, description, category_id) VALUES ('Test asset 1', 2020, 'Test Make 1', 'Test Model Number 1', 'Test Description 1', 1);
INSERT INTO asset (title, year, make, model_number, description, category_id) VALUES ('Test asset 2', 2020, 'Test Make 2', 'Test Model Number 2', 'Test Description 2', 2);

INSERT INTO usage_periodicity_unit (title) VALUES ('Miles');
INSERT INTO usage_periodicity_unit (title) VALUES ('Hours');

/* TODO: make this some kind of enum that is sourced from one place */
INSERT INTO time_periodicity_unit (title) VALUES ('Day');
INSERT INTO time_periodicity_unit (title) VALUES ('Week');
INSERT INTO time_periodicity_unit (title) VALUES ('Month');
INSERT INTO time_periodicity_unit (title) VALUES ('Year');

INSERT INTO preventativeTask (title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, asset_id) VALUES ('Test PreventativeTask 1', 'Test Instructions 1', 1, 1, 1, 1, 1);
INSERT INTO preventativeTask (title, instructions, time_periodicity_quantity, time_periodicity_unit_id, usage_periodicity_quantity, usage_periodicity_unit_id, asset_id) VALUES ('Test PreventativeTask 2', 'Test Instructions 2', 2, 2, 2, 2, 2);

/* TODO: make this some kind of enum that is sourced from one place */
INSERT INTO work_order_status (title) VALUES ('New');
INSERT INTO work_order_status (title) VALUES ('In Progress');
INSERT INTO work_order_status (title) VALUES ('Complete');

INSERT INTO work_order (preventativeTask_id, status_id, create_date, complete_date) VALUES (1, 1, '2020-01-01', '2020-01-02');
INSERT INTO work_order (preventativeTask_id, status_id, create_date) VALUES (2, 2, '2020-02-01');

INSERT INTO tool (title, size) VALUES ('Test Tool 1', 'Test Size 1');

INSERT INTO preventativeTask_tool (preventativeTask_id, tool_id) VALUES (1, 1);

INSERT INTO consumable (title) VALUES ('Test Consumable 1');

INSERT INTO preventativeTask_consumable (preventativeTask_id, consumable_id, quantity_note) VALUES (1, 1, '1');
*/