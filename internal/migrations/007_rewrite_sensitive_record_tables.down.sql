CREATE TABLE sensitive_record_types(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(20) NOT NULL
);

INSERT INTO sensitive_record_types (name) VALUES
    ('text'),
    ('binary-file'),
    ('credential'),
    ('payment-card');

CREATE TABLE sensitive_records(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    preview VARCHAR(50) NOT NULL,
    metadata VARCHAR(100),
    user_id INT NOT NULL,
    sensitive_record_type_id INT NOT NULL,

    FOREIGN KEY (sensitive_record_type_id) REFERENCES sensitive_record_types (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE payment_cards(
    sensitive_record_id INT PRIMARY KEY,
    number VARCHAR(19) NOT NULL,
    expire_date VARCHAR(5) NOT NULL,
    cardholder VARCHAR(26) NOT NULL,
    code VARCHAR(3) NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);

CREATE TABLE texts(
    sensitive_record_id INT PRIMARY KEY,
    text TEXT NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);

CREATE TABLE binary_files(
    sensitive_record_id INT PRIMARY KEY,
    uuid UUID NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);

CREATE TABLE credentials(
    sensitive_record_id INT PRIMARY KEY,
    username VARCHAR(25) NOT NULL,
    password VARCHAR(25) NOT NULL,

    FOREIGN KEY (sensitive_record_id) REFERENCES sensitive_records (id)
);

ALTER TABLE sensitive_records ADD FOREIGN KEY (sensitive_record_type_id) REFERENCES sensitive_record_types (id);
ALTER TABLE sensitive_records DROP COLUMN data;
ALTER TABLE sensitive_records DROP COLUMN type;

DROP TYPE SENSITIVE_RECORD_TYPES;
