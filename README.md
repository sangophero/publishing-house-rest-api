Possible ISBN Types **10/13**

METHOD | URL PATH               |Action     | Response Type|Example
-------|------------------------|---------- |--------------|-------
GET    |/books                  |list books |json          |/books
GET    |/book/{isbn}/{isbn_type}|get book   |json          |/book/0-201-53377-4/10
POST   |/book/{isbn_type}       |add book   |text          |/book/10
PUT    |/book/{isbn_type}       |update book|text          |/book/10
DELETE |/book/{isbn}/{isbn_type}|delete book|text          |/book/0-201-53377-4/10
