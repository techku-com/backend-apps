package User

import (
	"fmt"
)

func (u user) CheckExistsUser() {
	fmt.Println(u.dbCon.PostgreMainCon())
}
