package admin

import (
	"errors"
	invitation_db "invitation-app/datasources/mysql/invitation_db"
	helper "invitation-app/helper"
	"log"
)

const (
	queryFindByUsernameAndPassword = "SELECT id, username FROM admin WHERE username=? AND password=?"
)

func (admin *Admin) FindByUsernameAndPassword() error {
	stmt, err := invitation_db.Client.Prepare(queryFindByUsernameAndPassword)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	password := helper.MD5hash(admin.Password)

	result := stmt.QueryRow(admin.Username, password)
	if getErr := result.Scan(&admin.Id, &admin.Username); getErr != nil {
		log.Println(getErr)
		return errors.New("error when tying to find admin")
	}

	return nil
}
