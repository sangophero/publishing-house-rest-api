Possible ISBN Types **10/13**

METHOD | URL PATH               |Action     | Response Type|Example
-------|------------------------|---------- |--------------|-------
GET    |/books                  |list books |json          |/books
GET    |/book/{isbn}            |get book   |json          |/book/0-201-53377-4
POST   |/book/                  |add book   |text          |/book
PUT    |/book/                  |update book|text          |/book
DELETE |/book/{isbn}            |delete book|text          |/book/0-201-53377-4
