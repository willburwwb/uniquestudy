package post

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"test/database"
	"test/model"
	"time"

	"github.com/gin-gonic/gin"
)

func getUserID(name string) uint {
	var user model.User
	db := database.GetDB()
	db.Where("name=?", name).First(&user)
	return user.ID
}

var letters = []rune("0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}
func SetNickname() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func CreatPost(context *gin.Context) {
	var post model.Post
	if err := context.ShouldBind(&post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err,
		})
		return
	}
	log.Println("insert ", post.Title, post.Content)
	author := context.PostForm("author")
	post.AuthorID = uint64(getUserID(author))
	if post.AuthorID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "用户不存在",
		})
		return
	}
	post.Nickname = SetNickname()
	post.Vote = 0
	db := database.GetDB()
	if err := db.Create(post).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "数据创建失败",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  400,
		"message": "发布帖子成功",
	})
}
func DeletePost(context *gin.Context) {
	postid, err := strconv.Atoi(context.Param("id"))
	if err != nil || postid == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "参数错误",
		})
		return
	}
	db := database.GetDB()
	var post model.Post
	db.Where("id=?", postid).First(&post)

	if err := db.Delete(post).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "数据删除失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "删除帖子成功",
	})
	err = db.Where("post_id=?", postid).Delete(&model.Comment{}).Error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err,
		})
	}
	err = db.Where("post_id=?", postid).Delete(&model.Nick{}).Error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err,
		})
	}
}

func GetPostsByTitle(context *gin.Context) {
	title := context.Query("title")
	size, err := strconv.Atoi(context.Query("size"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "参数错误",
		})
	}
	posts := make([]model.Post, size)
	db := database.GetDB()
	db.Limit(size).Select([]string{"id", "title", "nickname", "content", "vote"}).Where("title = ?", title).Find(&posts)
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": posts,
	})
}

var timeLayoutStr = "2006-01-02"

func GetPostsByTime(context *gin.Context) {
	timestring := context.Query("time")
	time, err := time.Parse(timeLayoutStr, timestring)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":   200,
			"messsage": err,
		})
		return
	}
	size, err := strconv.Atoi(context.Query("size"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "参数错误",
		})
	}
	posts := make([]model.Post, size)
	db := database.GetDB()
	db.Limit(size).Select([]string{"id", "title", "nickname", "content", "vote"}).Where("created_at > ?", time).Find(&posts)
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": posts,
	})
}

func UpdatePost(context *gin.Context) {
	id, err := strconv.Atoi(context.PostForm("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "参数错误",
		})
	}
	title := context.PostForm("title")
	content := context.PostForm("content")
	db := database.GetDB()
	var post model.Post
	db.Where("id=?", id).First(&post)
	if post.ID != uint(id) {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "该帖子不存在",
		})
		return
	}
	if title != "" {
		post.Title = title
		db.Save(&post)
	}
	if content != "" {
		post.Content = content
		db.Save(&post)
	}
}
func GetPostByVote(context *gin.Context) {
	size, err := strconv.Atoi(context.Query("size"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "参数错误",
		})
	}
	posts := make([]model.Post, size)
	db := database.GetDB()
	db.Limit(size).Order("vote").Select([]string{"id", "title", "nickname", "content", "vote"}).Find(&posts)
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": posts,
	})
}
