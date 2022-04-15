INSERT INTO publishings (name)
VALUES ('Издательство Гиппократ'),('Изд. 2'), ('ТриМ');

INSERT INTO books (title, description, authors, year, edition, publishing_id)
VALUES
       ('Мастер и Маргарита', '123','Булгаков', 999, 'первое', 1),
       ('Белая Гвардия', 'комментарий', 'Булгаков', 1234, 'какое то', 2),
       ('Идиот','345', 'Достоевсий', 0001, 'третье', 3),
       ('Игрок','456', 'Достоевсий', 0055, 'третье', 1),
       ('Черный человек','ererer', 'Есенин', 1905, 'первое', 2);

INSERT INTO libraries (book_id, placement)
VALUES
       (2,'первая полка слева'),
       (1,'вторая полка слева'),
       (2,'первая полка справа'),
       (3,'первая полка сверху'),
       (5,'четвертая полка снизу'),
       (4,'первая полка снизу'),
       (2,'вторая полка сверху'),
       (4,'первая полка по середине'),
       (5,'стеллаж 1'),
       (1,'стеллаж 2');

INSERT INTO users (name)
VALUES ('Василий'), ('Иван'), ('Петр'), ('Константин');
INSERT INTO users (name)
VALUES ('Анна');

INSERT INTO usages (lib_book_id, user_id, date_issue, date_return)
VALUES
    (2,2,'05-05-2021','01-01-2022'),
    (1,3,'01-03-2021','01-04-2021'),
    (2,5,'02-01-2022',NULL),
    (4,3,'01-04-2021',NULL);


SELECT * from usages;

select * from libraries
;

SELECT  lib_book_id,book_id,title,
       case WHEN(date_return IS NULL AND date_issue IS NOT NULL ) THEN users.name
            WHEN(date_return IS NULL AND date_issue IS NULL ) THEN placement
            WHEN (date_return IS NOT NULL) THEN placement
    END AS placement
FROM libraries
    LEFT JOIN books USING(book_id)
    FULL JOIN usages USING(lib_book_id)
    LEFT JOIN users USING(user_id)
    LEFT JOIN publishings USING(publishing_id)
WHERE lib_book_id=2
ORDER BY lib_book_id;

SELECT  libraries.lib_book_id, books.title,
			CASE
			WHEN (date_issue IS NULL) THEN 'Книга не выдавалась'
			WHEN (date_issue IS NOT NULL) THEN users.name
			END as name,
			date_issue, date_return
			FROM libraries
			LEFT JOIN usages USING(lib_book_id)
			LEFT JOIN users USING(user_id)
            LEFT JOIN books USING(book_id)
			WHERE libraries.lib_book_id = 11;