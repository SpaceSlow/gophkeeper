ALTER TABLE sensitive_records DROP COLUMN data;

CREATE TABLE sensitive_datas(
    sensitive_record_id INT PRIMARY KEY,
    data BYTEA,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);
