/** add table users which store user's uuid, email, password and email verification **/

CREATE TABLE users (
  uuid VARCHAR(255) PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  email_verified BOOLEAN DEFAULT FALSE
);
