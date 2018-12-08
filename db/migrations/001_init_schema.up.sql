create table if not exists images(
  id INTEGER NOT NULL PRIMARY KEY,
  imageurl VARCHAR NOT NULL UNIQUE,
  baseurl VARCHAR,
  rating INTEGER DEFAULT 0,
  created DATETIME DEFAULT CURRENT_TIMESTAMP
);

create table if not exists tags(
  id INTEGER NOT NULL PRIMARY KEY,
  name VARCHAR NOT NULL UNIQUE,
  created DATETIME DEFAULT CURRENT_TIMESTAMP
);

create table if not exists image_tags(
  id INTEGER NOT NULL PRIMARY KEY,
  image_id INTEGER,
  tag_id INTEGER,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(image_id) REFERENCES images(id),
  FOREIGN KEY(tag_id) REFERENCES tags(id)
);
