package module

const createUsersTable string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  email TEXT NOT NULL,
  password TEXT NOT NULL
  );`
