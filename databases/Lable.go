package databases

type Lable struct {
	*DBModel
	//ID uint `gorm:"primary_key"`
	Name string `form:"name"`
	Pic string `form:"pic"`
	Sort uint `form:"sort"`
	Status uint `form:"status"`
}

type lableDAO struct {}
var LableDAO lableDAO

func(*lableDAO) CreateLable(lable Lable) error{
	DBInstance.AutoMigrate(&Lable{})
	return DBInstance.Create(&lable).Error
}

//获取全部标签
func (*lableDAO) ListLabe() ([]*Lable, error) {
	var lables []*Lable
	err := DBInstance.Where("status > ?", -1).Order("id desc").Find(&lables).Error
	return lables, err
}