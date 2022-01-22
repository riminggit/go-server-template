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

func RandSql(tableName string, queryCount string) string {
	queryMin := "(SELECT MIN(id) FROM " + tableName + ")"
	queryMax := "(SELECT MAX(id) FROM " + tableName + ")"
	querySelect := "SELECT * FROM " + tableName + " AS t1" + " "
	queryWhere := "WHERE t1.id >= t2.id" + " "
	queryOrder := "ORDER BY t1.id LIMIT " + queryCount
	return querySelect + "JOIN (SELECT ROUND(RAND() * (" + queryMax + " - " + queryMin + ") + " + queryMin + ") AS id) AS t2 " + queryWhere + queryOrder
}
