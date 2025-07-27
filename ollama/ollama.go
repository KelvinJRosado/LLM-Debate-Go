package main

import (
	"context"
	"time"

	ollama "github.com/prathyushnallamothu/ollamago"
)

// func main() {
//     // Create a new client with custom timeout
//     client := ollama.NewClient(
//         ollama.WithTimeout(time.Minute*5),
//     )

//     // Generate text
//     resp, err := client.Generate(context.Background(), ollama.GenerateRequest{
//         Model:  "llama3.2",
//         Prompt: "What is the capital of France?",
//     })
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(resp.Response)
// }

func main() {

	systemPrompt := `You are a member of a debate team. Your task is to engage in a debate with another participant, who is also a member of a debate team. You will be given a topic, and you must argue for or against it based on your assigned position. Your responses should be well-reasoned, articulate, and relevant to the topic at hand. Do not deviate from the topic, and ensure that your arguments are clear and concise. Remember, you are debating against another participant, so be prepared to counter their arguments effectively. Stay respectful and professional throughout the debate. Your goal is to persuade the audience of your position, so use evidence and logical reasoning to support your arguments. 
	
	You will be assigned a position (positive or negative) for each debate topic, and you must stick to that position throughout the debate. The debate will consist of multiple rounds, and you will have the opportunity to respond to your opponent's arguments in each round. Be strategic in your responses, and aim to strengthen your position while undermining your opponent's arguments. In each round, you will be given a specific question or prompt related to the topic, and you must respond accordingly. Your responses should be well-structured, with a clear introduction, body, and conclusion. Use examples and evidence to support your arguments, and anticipate potential counterarguments from your opponent. Remember, the goal is to persuade the audience of your position, so make sure your arguments are compelling and well-supported. 
	
	The structure of the debate will be as follows: 1. Opening statements: Each participant will present their initial arguments for their assigned position. 2. Rebuttals: Each participant will respond to the other's opening statements, addressing any weaknesses or counterarguments. 3. Further arguments: Each participant will present additional arguments to strengthen their position. 4. Closing statements: Each participant will summarize their key points and make a final appeal to the audience. Throughout the debate, maintain a professional demeanor, and focus on the topic at hand. Avoid personal attacks or irrelevant arguments, and stay focused on persuading the audience of your position.
	
	In each round, the positive participant will argue in favor of the topic, while the negative participant will argue against it. The positive participant will present arguments that support the topic, while the negative participant will present counterarguments that challenge the topic. Each participant will have the opportunity to respond to the other's arguments, and the debate will continue until a conclusion is reached. The positive participant will be given the first opportunity to present their arguments, followed by the negative participant. After both participants have presented their arguments, they will engage in a back-and-forth exchange, responding to each other's points and counterarguments. The debate will conclude with closing statements from both participants, summarizing their key arguments and making a final appeal to the audience. Remember, the goal is to persuade the audience of your position, so make sure your arguments are compelling and well-supported.`

	debateTopic := "The United States should adopt the metric system as its primary system of measurement."

	// Create two clients for positive and negative participants
	clientPositive := ollama.NewClient(
		ollama.WithTimeout(time.Minute * 5),
	)

	clientNegative := ollama.NewClient(
		ollama.WithTimeout(time.Minute * 5),
	)

	// Define messages for positive and negative participants
	// Each participant will have their own set of messages, starting with the system prompt
	// The messages will be used to generate responses for each participant in the debate
	messagesPositive := []ollama.Message{

		{
			Role:    "system",
			Content: systemPrompt + "\n\nYou have been assigned the positive position in the debate. Your task is to argue in favor of the topic: " + debateTopic,
		},
		{
			Role:    "user",
			Content: "Please provide your opening statement",
		},
	}

	respPositive, errPositive := clientPositive.Chat(context.Background(), ollama.ChatRequest{
		Model:    "llama3.2",
		Messages: messagesPositive,
	})

	if errPositive != nil {
		panic(errPositive)
	}

		println("Positive Participant Response:", respPositive.Message.Content)


	messagesNegative := []ollama.Message{

		{
			Role:    "system",
			Content: systemPrompt + "\n\nYou have been assigned the negative position in the debate. Your task is to argue against the topic: " + debateTopic,
		},
		{
			Role:    "user",
			Content: respPositive.Message.Content,
		},
	}

	respNegative, errNegative := clientNegative.Chat(context.Background(), ollama.ChatRequest{
		Model:    "llama3.2",
		Messages: messagesNegative,
	})

	if errNegative != nil {
		panic(errNegative)
	}

	// Print the responses from both participants
	println("\n\nNegative Participant Response:", respNegative.Message.Content)

}
