BEGIN TRANSACTION;

DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS room_members;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS users;


CREATE TABLE rooms (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  created_at DATETIME NOT NULL
);


CREATE TABLE users (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  created_at DATETIME NOT NULL
);


CREATE TABLE room_members (
  id INTEGER NOT NULL PRIMARY KEY,
  room_id TEXT  NOT NULL,
  user_id TEXT NOT NULL, 
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(room_id) REFERENCES rooms(id)
  UNIQUE(room_id, user_id)
);


CREATE TABLE messages (
  id TEXT NOT NULL PRIMARY KEY,
  content TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  creator_id TEXT NOT NULL, 
  room_id TEXT NOT NULL,
  FOREIGN KEY(room_id) REFERENCES rooms(id),
  FOREIGN KEY(creator_id) REFERENCES users(id)
);


-- Insert some users so it's easier to test
INSERT INTO users (id, name, created_at) VALUES ("random-uuid", "carlos", CURRENT_TIMESTAMP);
INSERT INTO users (id, name, created_at) VALUES ("random-uuid-2", "javi", CURRENT_TIMESTAMP);

COMMIT;
