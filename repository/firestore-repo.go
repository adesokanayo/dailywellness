package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/adesokanayo/innovation/entity"
	"google.golang.org/api/iterator"
	//citerator "google.golang.org/api"
)

const (
	projectID      string = "flash-chat-b4fd3"
	collectionNAME string = "posts"
)

var FireStoreRepo PostRepositoryInterface
type repo struct{}

//NewFireStoreRepository should be
func NewFireStoreRepository() PostRepositoryInterface {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v ", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionNAME).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("failed to add a post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	done := iterator.Done
	myiterator := client.Collection(collectionNAME).Documents(ctx)

	for {
		doc, err := myiterator.Next()
		//log.Println("Item: %v", doc)
		if err != nil {
			if err == done {
				break
			}
			log.Fatalf("Failed to read documents: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)

	}
	return posts, nil
}
