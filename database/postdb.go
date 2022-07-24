package database

import "test/model"

func CreatPost(post *model.Post) {
	db := GetDB()
	db.Create(post)
}
func DeletePost(post *model.Post) {
	db := GetDB()
	db.Delete(post)
}
