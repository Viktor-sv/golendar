
CREATE DATABASE golendar;
CREATE TABLE user (name VARCHAR(200), pass VARCHAR(200), token VARCHAR(200));
CREATE TABLE event (id VARCHAR(200),
                    title VARCHAR(200),
                    description VARCHAR(200),
                    name VARCHAR(200),
                    time VARCHAR(200),
                    timezone VARCHAR(200),
                    duration uint32,
                    Notes TEXT );
