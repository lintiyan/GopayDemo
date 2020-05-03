package util

import "github.com/bwmarrin/snowflake"

func GenerateId() string {
	n, _ := snowflake.NewNode(1)
	return n.Generate().String()
}
