CREATE TABLE IF NOT EXISTS users (
  id INT NOT NULL UNIQUE AUTO_INCREMENT,
  username VARCHAR(127) NOT NULL UNIQUE,
  password VARCHAR(127) NOT NULL,
  PRIMARY KEY (id)
)