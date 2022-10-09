insert into articles (title, contents, username, nice, created_at) values
  ("firstPost", "This is my first blog.", "Me", 0, now());
insert into articles (title, contents, username, nice) values
  ("secondPost", "This is my second blog.", "Me", 0);

insert into comments (article_id, message, created_at) values
  (1, "1 get", now());
insert into comments (article_id, message) values
  (1, "Welcome to the Internet.");
