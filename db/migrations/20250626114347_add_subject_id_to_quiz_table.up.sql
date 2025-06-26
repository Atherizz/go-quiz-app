ALTER TABLE quiz ADD subject_id INT,
ADD CONSTRAINT fk_subject FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE CASCADE;