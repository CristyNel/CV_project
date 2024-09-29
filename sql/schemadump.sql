DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS template;

CREATE TABLE users (
  id                INT AUTO_INCREMENT PRIMARY KEY,
  jobtitle          VARCHAR(128) NOT NULL,
  firstname         VARCHAR(128) NOT NULL,
  lastname          VARCHAR(128) NOT NULL,
  email             VARCHAR(128) NOT NULL,
  phone             VARCHAR(128) NOT NULL,
  address           VARCHAR(128) NOT NULL,
  city              VARCHAR(128) NOT NULL,
  country           VARCHAR(128) NOT NULL,
  postalcode        VARCHAR(128) NOT NULL,
  dateofbirth       DATE NOT NULL,
  nationality       VARCHAR(128) NOT NULL,
  summary           VARCHAR(256) DEFAULT '',
  workexperience    VARCHAR(128) DEFAULT '',
  education         VARCHAR(128) DEFAULT '',
  skills            VARCHAR(128) DEFAULT '',
  languages         VARCHAR(128) DEFAULT ''
);

CREATE TABLE template (
  id            INT AUTO_INCREMENT PRIMARY KEY,
  path          VARCHAR(128) NOT NULL
);

INSERT INTO users
  (jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality)
VALUES
  ('Inginer Software', 'Damaris', 'Nicolae', 'damarisnicolae@google.com', '07076965173', 'Str. Trandafirilor nr.10', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Tabita', 'Petruneac', 'tabitapetruneac@google.com', '07076965173', 'Str. Crizantemelor nr.5', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Cristy', 'Buliga', 'cristybuliga@google.com', '07076965173', 'Str. Lalelelor nr.17', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Cezar', 'Sandra', 'cezarsandra@google.com', '07076965173', 'Str. Violetei nr.3', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Filip', 'Petruneac', 'filippetruneac@google.com', '07076965173', 'Str. Bujorilor nr.8', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Delu', 'Giurgiu', 'deludelu@google.com', '07076965173', 'Str. Magnoliei nr.12', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Robert', 'Oros', 'robertoros@google.com', '07076965173', 'Str. Zambilelor nr.4', 'Stei', 'Romania', '010001', '1990-11-11', 'română'),
  ('Inginer Software', 'Lois', 'Nicoras', 'loisnicoras@google.com', '07076965173', 'Str. Garoafelor nr.9', 'Stei', 'Romania', '010001', '1990-11-11', 'română');

INSERT INTO template
  (path)
VALUES
  ('./template1.html'),
  ('./template2.html'),
  ('./template3.html');
