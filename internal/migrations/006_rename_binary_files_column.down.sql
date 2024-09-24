ALTER TABLE binary_files ADD COLUMN path_name VARCHAR(10);
ALTER TABLE binary_files DROP COLUMN uuid;