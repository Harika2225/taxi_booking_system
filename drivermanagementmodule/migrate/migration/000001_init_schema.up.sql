-- Create Driver table
CREATE TABLE Driver (
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(255) NOT NULL,
  lastname VARCHAR(255) NOT NULL,
  phone VARCHAR(15) UNIQUE NOT NULL,
  license VARCHAR(20) UNIQUE NOT NULL,
  createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Insert a sample record
INSERT INTO Driver (firstname, lastname, phone, license) VALUES
    ('John', 'Doe', '+1234567890', 'ABC123'),
    ('Jane', 'Smith', '+9876543210', 'XYZ789');
