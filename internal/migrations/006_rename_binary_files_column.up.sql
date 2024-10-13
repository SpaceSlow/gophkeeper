ALTER TABLE binary_files ADD COLUMN uuid UUID;
ALTER TABLE binary_files DROP COLUMN path_name;
