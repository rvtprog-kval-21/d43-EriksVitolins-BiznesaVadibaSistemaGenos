package projects

import (
	"api/database"
	"fmt"
)

type Tag struct {
	ID        int    `json:"id" gorm:"primaryKey;not null"`
	Name      string `json:"name"  gorm:"index;not null"`
	ProjectID int    `json:"project_id" gorm:"index"`
}

func GetTag(name string, projectID string) (Tag, interface{}) {
	var tag Tag
	response := database.DBConn.Where("project_id = ?", projectID).Where("name = ?", name).Find(&tag)
	return tag, response.Error
}

func AddTag(tag *Tag) interface{} {
	results := database.DBConn.Create(&tag)
	return results.Error
}

func RemoveOrphans(names []string, projectID string) {
	var tags []Tag
	var ids []int

	database.DBConn.Where("project_id = ?", projectID).Where("name NOT IN ?", names).Find(&tags)
	for _, iter := range tags {
		ids = append(ids, iter.ID)
	}
	database.DBConn.Where("id IN ?", ids).Delete(&Tag{})
	database.DBConn.Where("tag_id IN ?", ids).Update("tag_id", nil)
	fmt.Println(tags)
}
