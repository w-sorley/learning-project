CREATE TABLE IF NOT EXISTS user (
  id            BIGINT AUTO_INCREMENT,
  username      VARCHAR(100),
  password_salt VARCHAR(100),
  CONSTRAINT pk_user PRIMARY KEY (id)
)CHARSET =utf8, ENGINE =InnoDB;
CREATE UNIQUE INDEX idx_user_username ON users(username);

CREATE TABLE IF NOT EXISTS user_roles(
  id BIGINT AUTO_INCREMENT,
  username VARCHAR(100),
  role_name VARCHAR(100),
  CONSTRAINT pk_user_roles PRIMARY KEY (id)
)CHARSET =utf8,ENGINE =InnoDB;
CREATE UNIQUE INDEX idx_user_roles ON user_roles(username, role_name) ;

CREATE TABLE IF NOT EXISTS roles_permission(
  id BIGINT AUTO_INCREMENT,
  role_name VARCHAR(100),
  permission VARCHAR(100),
  CONSTRAINT pk_roles_permission PRIMARY KEY (id)
)CHARSET =utf8,ENGINE =InnoDB;
CREATE UNIQUE INDEX idx_roles_permission ON roles_permission(role_name, permission) ;

INSERT INTO users(username, password) VALUES ('wang','123')
