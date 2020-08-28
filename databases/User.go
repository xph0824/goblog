package databases

type User struct {
	*DBModel
	Name string 	`form:"name"`
	Age uint 		`form:"age"`
	Address string 	`form:"address"`
	Phone string 	`form:"phone"`
	Email string 	`form:"email"`

}

type userDAO struct {}

// UserDAO user dao
var UserDAO userDAO

func(*userDAO) CreateUser(user User) error{
	DBInstance.AutoMigrate(&User{})
	return DBInstance.Create(&user).Error
}

// First get the first record of user
func (*userDAO) FirstUser() (User, error) {
	var user User
	err := DBInstance.First(&user).Error
	return user, err
}

// Update update user record
func (*userDAO) UpdateUser(user User) error {
	return DBInstance.Model(&User{}).Updates(&user).Error
}

// Delete set all to delete state
func (*userDAO) DeleteUser() error {
	return DBInstance.Delete(&User{}).Error
}
