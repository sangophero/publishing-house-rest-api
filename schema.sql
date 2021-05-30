CREATE TABLE `books` (
  `isbn` TEXT NOT NULL UNIQUE,
  `title` TEXT NOT NULL,
  `author` TEXT NOT NULL,
  `price` FLOAT NOT NULL,
  `description` TEXT NOT NULL,
  `url_cover` TEXT NOT NULL,
  `status` BOOLEAN NOT NULL
);
