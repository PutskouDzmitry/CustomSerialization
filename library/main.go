package main

import (
	"fmt"
	"github.com/kataras/go-serializer"
	"github.com/kataras/go-serializer/markdown"
)

func main() {
	markdownSerializer := markdown.New()
	serializer.For("markdown", markdownSerializer)
	result, _ := serializer.SerializeToString("markdown", "## Hello")
	// Optional third parameter is a map[string]interface{}, which is any runtime options for the serializers registered to 'markdown' key
	fmt.Println(result)
}
