-- Create Booking table
CREATE TABLE Booking (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  pickupaddress VARCHAR(255) NOT NULL,
  destination VARCHAR(255) NOT NULL,
  date VARCHAR(50) NOT NULL,
  status VARCHAR(50) DEFAULT 'Pending' ,
  createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- Insert sample bookings
INSERT INTO Booking (customer_id, pickupaddress, destination, date) VALUES
(1, 'First Pickup Address', 'First Destination', '2024-02-09 10:00:00'),
(2, 'Second Pickup Address', 'Second Destination', '2024-02-10 12:30:00');