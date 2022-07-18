DROP DATABASE IF EXISTS library;
CREATE DATABASE library;
USE library;

DROP TABLE IF EXISTS Author;
create table Author (
                        AuthorId INT ,
                        FirstName varchar(50),
                        LastName varchar(50),
                        DateOfBirth varchar(50),
                        PenName     varchar(50),
                        PRIMARY KEY (AuthorId)
)

DROP TABLE IF EXISTS Books;
CREATE TABLE Books(
                      BookId INT ,
                      AuthorId INT,
                      Title  VARCHAR(50),
                      Publications VARCHAR(50),
                      PublishedDate VARCHAR(50),
                      PRIMARY KEY (BookId),
                      FOREIGN KEY (AuthorId) REFERENCES Author(AuthorId)
)