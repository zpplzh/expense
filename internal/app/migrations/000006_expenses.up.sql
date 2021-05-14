ALTER TABLE expenses
ADD CONSTRAINT fk_category FOREIGN KEY (categoryid) REFERENCES category(categoryid) ON DELETE SET NULL;