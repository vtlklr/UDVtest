package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upTestdata, downTestdata)
}

func upTestdata(tx *sql.Tx) error {
	_, err := tx.Exec(`
INSERT INTO publishings (name)
VALUES ('Издательство Гиппократ'),('Изд. 2'), ('ТриМ');
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

INSERT INTO usages (lib_book_id, user_id, date_issue, date_return)
VALUES
    (2,2,'05-05-2021','01-01-2022'),
    (1,3,'01-03-2021','01-04-2021'),
    (2,5,'02-01-2022',NULL),
    (4,3,'01-04-2021',NULL);
INSERT INTO books (title, description, authors, year, edition, publishing_id)
VALUES
       ('Мастер и Маргарита', '123','Булгаков', 999, 'первое', 1),
       ('Белая Гвардия', 'комментарий', 'Булгаков', 1234, 'какое то', 2),
       ('Идиот','345', 'Достоевсий', 0001, 'третье', 3),
       ('Игрок','456', 'Достоевсий', 0055, 'третье', 1),
       ('Черный человек','ererer', 'Есенин', 1905, 'первое', 2);
`)
	return err
}

func downTestdata(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
