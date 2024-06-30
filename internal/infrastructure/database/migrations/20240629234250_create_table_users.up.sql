CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  middle_name VARCHAR(50),
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,

  status INT DEFAULT 0,
  email_verified BOOLEAN DEFAULT FALSE,
  phone_verified BOOLEAN DEFAULT FALSE,

  username VARCHAR(50) UNIQUE,
  phone VARCHAR(15),

  encrypted_password VARCHAR(255) NOT NULL,
  password_updated_at TIMESTAMP,
  reset_password_sent_at TIMESTAMP,
  remember_me_created_at TIMESTAMP,

  confirmation_token VARCHAR(255),
  confirmed_at TIMESTAMP,
  confirmation_sent_at TIMESTAMP,

  provider VARCHAR(50),
  uid VARCHAR(100),

  sign_in_count INT DEFAULT 0,
  current_sign_in_at TIMESTAMP,
  last_sign_in_at TIMESTAMP,
  current_sign_in_ip VARCHAR(45),
  last_sign_in_ip VARCHAR(45),
  sign_up_ip VARCHAR(45),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
) ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;
