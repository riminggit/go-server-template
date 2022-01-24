package util

import (
	"math/rand"
	"time"
)

func RandSliceValue(xs []string) string {
	return xs[rand.Intn(len(xs))]
}

func RandNum() int {
	// 取当前时间戳
	var timeStamp = time.Now().Unix()
	// 构造一个rand，并使用时间戳作为他的随机种子
	r := rand.New(rand.NewSource(timeStamp))
	// 取100以内的随机数
	num := r.Intn(100)
	return num
}

func RandNumC(count int, min int) int {
	// 取当前时间戳
	var timeStamp = time.Now().Unix()
	// 构造一个rand，并使用时间戳作为他的随机种子
	r := rand.New(rand.NewSource(timeStamp))
	// 取100以内的随机数
	num := r.Intn(count)
	return num + min
}

// 有问题，如果有的id不存在，会拿错id
func RandSql(tableName string, queryCount string) string {
	queryMin := "(SELECT MIN(id) FROM " + tableName + ")"
	queryMax := "(SELECT MAX(id) FROM " + tableName + ")"
	querySelect := "SELECT * FROM " + tableName + " AS t1" + " "
	queryWhere := "WHERE t1.id >= t2.id" + " "
	queryOrder := "ORDER BY t1.id LIMIT " + queryCount
	return querySelect + "JOIN (SELECT ROUND(RAND() * (" + queryMax + " - " + queryMin + ") + " + queryMin + ") AS id) AS t2 " + queryWhere + queryOrder
}

// 最慢的随机方法
func RandSqlCommon(tableName string, queryCount string) string {
	return "SELECT * FROM " + tableName + " WHERE  is_use = 1 ORDER BY RAND() LIMIT " + queryCount
}
