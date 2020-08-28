package models

import "goModWork/databases"


func CreateLable(lable databases.Lable) error {
	return databases.LableDAO.CreateLable(databases.Lable{
		Name: lable.Name,
		Pic: lable.Pic,
		Sort: lable.Sort,
		Status: 0,
	})
}

func ListLable() ([]*databases.Lable, error) {
	return databases.LableDAO.ListLabe()
}
