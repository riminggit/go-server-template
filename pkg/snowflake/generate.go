package snowflake

import logging "go-server-template/pkg/log"

func GenerateID(n int) int {
	node, err := NewNode(int64(n))
	if err != nil {
		logging.Error(err)
	}
	return int(node.Generate().Int64())
}