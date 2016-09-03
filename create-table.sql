create table books(
  pk integer primary key autoincrement,
  title text,
  author text,
  id text,
  classification text
);

--delete from books where pk in (select pk from books where title = '');
--select * from books where title = '';
