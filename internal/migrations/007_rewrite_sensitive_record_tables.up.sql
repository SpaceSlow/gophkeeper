DROP TABLE binary_files;
DROP TABLE credentials;
DROP TABLE payment_cards;
DROP TABLE texts;

ALTER TABLE sensitive_records DROP COLUMN sensitive_record_type_id;
DROP TABLE sensitive_record_types;

CREATE TYPE SENSITIVE_RECORD_TYPES AS ENUM ('binary', 'credential', 'payment-card', 'text');
ALTER TABLE sensitive_records ADD COLUMN data BYTEA;
ALTER TABLE sensitive_records ADD COLUMN type SENSITIVE_RECORD_TYPES;
