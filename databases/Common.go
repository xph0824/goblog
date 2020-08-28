//数据库操作
package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type DBModel struct {
	*gorm.Model
}

//数据库实例
var DBInstance *gorm.DB

//初始化数据库实例
func InitDB() {

	log.Println("Init Databases.....")
	//修改默认的表明规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "tbl_" + defaultTableName;
	}
	//线上
	//db, err := gorm.Open("mysql", "root:admin123@tcp(122.51.207.225:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,不为复数
	//db.Callback().Create().Replace("gorm:update_time_stamp",updateStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DBInstance = db
}

//关闭数据库
func CloserDB() {
	DBInstance.Close()
}

func updateStampForCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		//此处时间如果是int64就修改time.Now().Unix()就行了
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//此处时间如果是int64就修改time.Now().Unix()就行了
		scope.SetColumn("UpdatedAt", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				//此处时间如果是int64就修改time.Now().Unix()就行了
				scope.AddToVars(time.Now().Format("2006-01-02 15:04:05")),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

