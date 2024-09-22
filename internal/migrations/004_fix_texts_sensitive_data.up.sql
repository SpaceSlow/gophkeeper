DROP TABLE text_files;

CREATE TABLE texts(
    sensitive_record_id INT PRIMARY KEY,
    text TEXT NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);
