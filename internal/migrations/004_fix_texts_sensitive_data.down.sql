DROP TABLE texts;

CREATE TABLE text_files(
    sensitive_record_id INT PRIMARY KEY,
    path_name VARCHAR(10) NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);
