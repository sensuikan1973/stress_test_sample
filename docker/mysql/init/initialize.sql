CREATE DATABASE IF NOT EXISTS sample_for_qiita;
CREATE TABLE IF NOT EXISTS sample_for_qiita.greetings (id int auto_increment, text varchar(255), index(id));
INSERT INTO greetings (text) VALUES ("hello"), ("goodbye"), ("good morning"), ("good evening");
