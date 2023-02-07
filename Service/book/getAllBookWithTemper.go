package book

import (
	"bookshop/DataHandler/book"
	"bookshop/Model"
	"bookshop/Service/user"
	"database/sql"
	"fmt"
	"math"
	"time"
)

func GetAllBookWithTemper(db *sql.DB, originToken []string) (int, map[string]interface{}) {
	books, err := book.GetAllBook(db)
	if err != nil {
		return 202, map[string]interface{}{
			"status": 1,
			"info":   "get books fail",
			"data": map[string]interface{}{
				"books": []map[string]interface{}{}},
		}
	}
	//从这里开始，执行一个函数，获取到排序后的books
	books = sortWithTemper(db, books)
	for _, i := range books {
		fmt.Println(i.Temperature)
	}
	//当没有传入token时
	if len(originToken) == 0 {
		//is_star默认赋值为false
		return 200, map[string]interface{}{
			"status": 0,
			"info":   "success",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	data, err1 := user.ParseToken(originToken[0])
	if err1 != nil {
		return 200, map[string]interface{}{
			"status": 2,
			"info":   "parse token fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	username := data.Issuer
	collections, err2 := book.GetCollectionsWithUserName(db, username)
	if err2 != nil {
		return 200, map[string]interface{}{
			"status": 3,
			"info":   "get collections fail, return default books(without is_star)",
			"data": map[string]interface{}{
				"books": books},
		}
	}
	//collection表的记录是有一条就意味着收藏了
	//range是复制，不能修改到值
	for i := 0; i < len(collections); i++ {
		for j := 0; j < len(books); j++ {
			if collections[i].BookId == books[j].Id {
				books[j].IsStar = true
			}
		}
	}
	return 200, map[string]interface{}{
		"status": 0,
		"info":   "success",
		"data": map[string]interface{}{
			"books": books},
	}

}

func sortWithTemper(db *sql.DB, books []Model.Book) []Model.Book {
	//遍历每本书，获取收藏数，评论总数，最后赋予热度值
	for i := 0; i < len(books); i++ {
		var publishTimeStr = books[i].PublishTime
		allTime := computeFromTimeNow(publishTimeStr)
		collectionCount := book.GetCollectionsCount(db, books[i].Id)
		commentCount := book.GetCommentCount(db, books[i].Id)
		scoreMother := math.Pow(float64(allTime/100000+1), 1.1)
		fmt.Println("scoreMother", scoreMother)
		//赋以初值100
		scoreSon := collectionCount + commentCount + 100
		fmt.Println("scoreSon", scoreSon)
		heat := float64(scoreSon) / scoreMother
		books[i].Temperature = heat
		fmt.Println("heat", heat)
	}
	for i := 0; i < len(books); i++ {
		for j := i + 1; j < len(books); j++ {
			if books[i].Temperature < books[j].Temperature {
				temp := books[i]
				books[i] = books[j]
				books[j] = temp
			}
		}
	}
	return books
}

func computeFromTimeNow(publishTimeStr string) int {
	loc, _ := time.LoadLocation("Local")
	formatTime, _ := time.ParseInLocation("2006-01-02 15:04:05", publishTimeStr, loc)
	timeStr := time.Now().Format("2006-01-02 15:04:05")                           //转化所需模板
	formatTimeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) //使用模板在对应时区转化为time.time类型
	return int(formatTimeNow.Sub(formatTime))
}
