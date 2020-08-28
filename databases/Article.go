package databases

type Article struct {
	*DBModel
	//ID uint 		`gorm:"primary_key"`
	Uid uint 		`form:"uid"`
	LableId uint 	`form:"lable_id"`
	Title string 	`form:"title"`
	Content string 	`form:"content"`
	Pic string 		`form:"pic"`

}

type articleDAO struct {}
var ArticleDAO articleDAO
func(*articleDAO) CreateArticle(article Article) error{
	DBInstance.AutoMigrate(&Article{})
	return DBInstance.Create(&article).Error
}

// First get the first record of user
func (*articleDAO) FirstArtilce() (Article, error) {
	var article Article
	err := DBInstance.First(&article).Error
	return article, err
}

// Update update user record
func (*articleDAO) UpdateArticle(article Article) error {
	return DBInstance.Model(&Article{}).Updates(&article).Error
}

// Delete set all to delete state
func (*articleDAO) DeleteArticle() error {
	return DBInstance.Delete(&Article{}).Error
}

func (*articleDAO) FindArticleLimit(page,lableId int, keyword string) ([]*Article, error) {
	var articles []*Article
	pageSize := 5

	if keyword == "" && lableId > 0 {
		err := DBInstance.Where("lable_id = ?", lableId).Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&articles).Error
		return articles, err
	} else if keyword != "" && lableId > 0 {
		err := DBInstance.Where("able_id = ? AND title LIKE ?", lableId, "%"+keyword+"%").Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&articles).Error
		return articles, err
	} else if keyword != "" && lableId <= 0 {
		err := DBInstance.Where("title LIKE ?", "%"+keyword+"%").Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&articles).Error
		return articles, err
	} else {
		err := DBInstance.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&articles).Error
		return articles, err
	}
}

func (*articleDAO) FinArticleCount(keyword string) int {
	var count int
	if keyword == "" {
		DBInstance.Model(&Article{}).Count(&count)
	} else {
		DBInstance.Model(&Article{}).Where("title LIKE ?", "%"+keyword+"%").Count(&count)
	}
	return count
}