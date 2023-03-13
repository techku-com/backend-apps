package User

func (u user) GetUser() {
	u.repo.User.CheckExistsUser()
}
