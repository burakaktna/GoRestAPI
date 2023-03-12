CREATE TABLE articles (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  published_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE tags (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE article_tags (
  id INT NOT NULL AUTO_INCREMENT,
  article_id INT UNSIGNED NOT NULL,
  tag_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (article_id) REFERENCES articles(id),
  FOREIGN KEY (tag_id) REFERENCES tags(id)
);