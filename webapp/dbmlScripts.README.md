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