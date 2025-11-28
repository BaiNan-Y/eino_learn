package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中女生"),
		schema.UserMessage("你好"),
	}
	/*
		response, err := model.Generate(ctx, input)
		if err != nil {
			panic(err)
		}
		print(response.Content)
	*/
	reader, err := model.Stream(ctx, input)
	if err != nil {
		panic(err)
	}
	defer reader.Recv()

	// 这种无限循环的循环结束方式是，所有信息打印完毕之后的报错导致的循环推出，这是不好的写法，这里由于demo先这样写了
	for {
		chunk, err := reader.Recv()
		if err != nil {
			panic(err)
		}
		print(chunk.Content)
	}
}
