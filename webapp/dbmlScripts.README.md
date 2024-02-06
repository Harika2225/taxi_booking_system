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
-----------------