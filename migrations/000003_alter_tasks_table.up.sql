ALTER TABLE task
ADD COLUMN fk INT,
ADD CONSTRAINT fk_task_list FOREIGN KEY (fk) REFERENCES list(id);
