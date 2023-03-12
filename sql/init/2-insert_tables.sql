INSERT INTO articles (title, content, published_at) VALUES
  ('Article 1 Title', 'Article 1 Content', NOW()),
  ('Article 2 Title', 'Article 2 Content', NOW()),
  ('Article 3 Title', 'Article 3 Content', NOW()),
  ('Article 4 Title', 'Article 4 Content', NOW()),
  ('Article 5 Title', 'Article 5 Content', NOW()),
  ('Article 6 Title', 'Article 6 Content', NOW()),
  ('Article 7 Title', 'Article 7 Content', NOW()),
  ('Article 8 Title', 'Article 8 Content', NOW()),
  ('Article 9 Title', 'Article 9 Content', NOW()),
  ('Article 10 Title', 'Article 10 Content', NOW());

INSERT INTO tags (name) VALUES
  ('tag1'),
  ('tag2'),
  ('tag3'),
  ('tag4'),
  ('tag5');

INSERT INTO article_tags (article_id, tag_id) VALUES
  (1, 1),
  (2, 2),
  (3, 3),
  (4, 4),
  (5, 5),
  (6, 1),
  (7, 2),
  (8, 3),
  (9, 4),
  (10, 5);