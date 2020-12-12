package utils

import (
	"context"
	"log"
	"os"

	"github.com/Bendomey/goutilities/pkg/hashpassword"

	"github.com/Bendomey/task-assignment/repository"
)

// SaveAdminInfo checks on startup if there is a super admin and then creates one when there isn't
func SaveAdminInfo(repo repository.Repository) error {
	//check if there are any users
	rows, fetchRowErr := repo.GetAll(context.Background(), "select COUNT(*) from users")
	if fetchRowErr != nil {
		return fetchRowErr
	}
	var length int
	for rows.Next() {
		errOnScan := rows.Scan(&length)
		if errOnScan != nil {
			return errOnScan
		}
	}

	if length == 0 {
		log.Println("No users in db, creating one")
		//hash password
		hash, hashErr := hashpassword.HashPassword(os.Getenv("ADMIN_PASSWORD"))
		if hashErr != nil {
			return hashErr
		}
		//create admin
		_, err := repo.Insert(context.Background(), "insert into users (fullname,password,email,user_type) values($1,$2,$3,$4)", os.Getenv("ADMIN_NAME"), hash, os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_TYPE"))
		if err != nil {
			return err
		}
	}
	log.Println("Taskerman booted successfully")
	return nil
}
