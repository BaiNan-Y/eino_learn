package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
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
	// input := []*schema.Message{
	// 	schema.SystemMessage("你是一个可爱的高中女生"),
	// 	schema.UserMessage("你好"),
	// }

	// 这里是 ChatTemplate 的使用
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，堂吉诃德先生:{task}",
		},
	)
	params := map[string]any{
		"role": "著名文学人物堂吉诃德",
		"task": "写一首关于周末的押韵的诗句",
	}
	message, err := template.Format(ctx, params)

	response, err := model.Generate(ctx, message)
	if err != nil {
		panic(err)
	}
	print(response.Content)
}


