package pasetobd

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// func TestUpdateGetData(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "healhero_db")
// 	datagedung := GetAllBangunanLineString(mconn, "healhero_db")
// 	fmt.Println(datagedung)
// }


func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PASETO1")
	var userdata User
	userdata.Username = "serbaevent"
	userdata.Password = "begitulah"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PASETO1")
	var userdata User
	userdata.Username = "serbaevent"
	userdata.Password = "begitulah"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}
