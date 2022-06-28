package schema

import (
	invitation_db "invitation-app/datasources/mysql/invitation_db"
	"log"
)

func CreateSchema() {
	sql := make(map[int]string)

	sql[0] = `CREATE TABLE IF NOT EXISTS  admin (
			id INT NOT NULL AUTO_INCREMENT,
			username VARCHAR(250),
			password VARCHAR(250),
			PRIMARY KEY ( id )
		);
	`

	sql[1] = `CREATE TABLE IF NOT EXISTS  client_tokens (
		id INT NOT NULL AUTO_INCREMENT,
		token VARCHAR(250),
		status VARCHAR(250),
		created_at datetime,
		updated_at datetime,
		PRIMARY KEY ( id )
	);`

	sql[2] = `INSERT INTO admin(username, password)
		SELECT * FROM (
			SELECT 'userPulse', md5('qwerty1234')	
		) as t0
		WHERE NOT EXISTS(SELECT '' FROM admin WHERE username = 'userPulse')
	`

	log.Println("Creating database tables")

	for _, v := range sql {
		stmt, err := invitation_db.Client.Prepare(v)
		if err != nil {
			log.Println("error when trying to prepare save user statement", err)
		}

		defer stmt.Close()
		_, saveErr := stmt.Exec()
		if saveErr != nil {
			log.Println("error when trying to save user", saveErr)
		}
	}

}
