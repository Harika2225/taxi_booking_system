-- Create payment table
CREATE TABLE Payment (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10, 2),
    payment_date VARCHAR(50),
    customer_id INT,
    driver_id INT,
    payment_status VARCHAR(20) DEFAULT 'Pending',
    payment_method_id INT
);

-- Insert sample data into payment table
INSERT INTO Payment (amount, payment_date, payment_method_id) VALUES
(50.00, '2024-02-12 12:00:00', 1),
(30.50, '2024-02-12 13:30:00', 2),
(25.75, '2024-02-12 15:45:00', 3);


-- Create payment_method table (assuming it doesn't already exist for the foreign key reference)
CREATE TABLE Payment_method (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    description TEXT
);
-- Insert sample data into payment_method table
INSERT INTO Payment_method (name, description) VALUES
('Credit Card', 'Payment using credit card'),
('Debit Card', 'Payment using debit card'),
('Mobile Payment', 'Payment using mobile app');
