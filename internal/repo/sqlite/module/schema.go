package module

const createUsersTable string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  email TEXT NOT NULL,
  password TEXT NOT NULL
  );`

const createCoinsTable string = `
  CREATE TABLE IF NOT EXISTS coins (
  id INTEGER NOT NULL PRIMARY KEY,
  coin_id TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
  );`
