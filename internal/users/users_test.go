package users_test

import (
	"fmt"
)

const tableCreationSQL = `
CREATE TABLE IF NOT EXISTS users
(
	id SERIAL PRIMARY KEY,
	email VARCHAR (50) UNIQUE,
	firstName VARCHAR (50)
);
`
