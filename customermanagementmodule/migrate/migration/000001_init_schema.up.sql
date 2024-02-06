-- Create Customer table
CREATE TABLE Customer (
  id SERIAL PRIMARY KEY,
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  phone VARCHAR(255),
  address VARCHAR(255),
  createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO Customer (firstName, lastName, email, phone, address) VALUES
  ('John', 'Doe', 'john.doe@example.com', '1234567890', '123 Main St'),
  ('Jane', 'Smith', 'jane.smith@example.com', '9876543210', '456 Oak St');
