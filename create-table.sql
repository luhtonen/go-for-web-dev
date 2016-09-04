create table books(
  pk integer primary key autoincrement,
  title text,
  author text,
  id text,
  classification text
);

-- Update: adding user column to books table
alter table books add column user varchar(255);

--delete from books where pk in (select pk from books where title = '');
--select * from books where title = '';
