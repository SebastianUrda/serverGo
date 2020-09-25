package main

import (
	"awesomeProject/entity"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	_ "google.golang.org/api/iterator"
	"log"
)

type QuestionRepository interface {
	Save(question *entity.Question) (*entity.Question, error)
	FindAll() ([]entity.Question, error)
}

type repository struct{}

func NewQuestionRepository() QuestionRepository {
	return &repository{}
}

const (
	projectId      string = "masterapp-5c255"
	collectionName string = "questions"
)

func (*repository) Save(question *entity.Question) (*entity.Question, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var answers []map[string]interface{}
	for _, element := range question.PossibleAnswers {

		elementMap, _ := json.Marshal(element)
		var inInterface map[string]interface{}
		json.Unmarshal(elementMap, &inInterface)
		answers = append(answers, inInterface)
	}

	var toAdd = map[string]interface{}{
		"id":              question.Id,
		"possibleAnswers": answers,
		"text":            question.Text,
		"type":            question.Type,
	}
	fmt.Println(toAdd)
	_, _, err = client.Collection(collectionName).Add(ctx, toAdd)

	if err != nil {
		log.Fatal("Failed to add to collection: %v", err)
		return nil, err
	}

	return question, nil
}

func (r *repository) FindAll() ([]entity.Question, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Println("Failed to create Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var questions []entity.Question
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Println("Failed to iterate list ", err)
			return nil, err
		}
		var answers []entity.PossibleAnswer

		if doc.Data()["possibleAnswers"] != nil {
			for _, element := range doc.Data()["possibleAnswers"].([]interface{}) {
				elementMap := element.(map[string]interface{})
				answer := entity.PossibleAnswer{
					Ans:   elementMap["ans"].(string),
					Value: elementMap["value"].(float64),
				}
				answers = append(answers, answer)
			}
		}

		question := entity.Question{
			Id:              doc.Data()["id"].(int64),
			Text:            doc.Data()["text"].(string),
			Type:            doc.Data()["type"].(string),
			PossibleAnswers: answers,
		}
		fmt.Println(question)
		questions = append(questions, question)
	}
	fmt.Println(questions)

	return questions, nil
}
