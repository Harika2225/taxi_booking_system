---------------CUSTOMER MANAGEMENT MODULE----------------

Table Customer {
  id            Int       [primary key, increment]
  firstName     String    [not null]
  lastName      String    [not null]
  email         String    [unique, not null]
  phone         String
  address       String
  createdAt     DateTime  [default: `now()`]
  updatedAt     DateTime  [default: `now()`]
}

======================================================================
-----------------DRIVER MANAGMENT MODULE------------------------------


Table Driver {
  id          Int       [pk, increment]
  firstname   String    [not null]
  lastname    String    [not null]
  phone String    [unique, not null]
  license String [unique, not null]
  createdAt   DateTime  [default: `now()`]
  updatedAt   DateTime  [default: `now()`]
}


========================================================================
--------------------BOOKING MANAGAMENT MODULE---------------------------

Table Booking {
  id            Int       [primary key, increment]
  customer_id   Int       [not null]
  pickupaddress String    [not null]
  destination   String    [not null]
  date          DateTime  [not null]
  status        String    [default: "Pending"]
  createdAt     DateTime  [default: `now()`]
  updatedAt     DateTime  [default: `now()`]
}