package Entity

type User struct {
	Id string `json:"user_id"`
	Name string `json:"name"`
}

const TableUser = `
	CREATE TABLE IF NOT EXISTS user (
	    user_id INT AUTO_INCREMENT PRIMARY KEY,
	    name VARCHAR(255) NOT NULL
	)  ENGINE=INNODB;
`

