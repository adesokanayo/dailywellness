package repository

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/adesokanayo/dailywellness/entity"
	"google.golang.org/api/iterator"
	//citerator "google.golang.org/api"
)

const (
	projectID      string = "flash-chat-b4fd3"
	collectionNAME string = "posts"
)

//var FireStoreRepo PostRepositoryInterface
type repo struct{}

//NewFireStoreRepository should be
func NewFireStoreRepository() PostRepositoryInterface {
	return &repo{}
}

func (*repo) Save(post *entity.Tip) (*entity.Tip, error) {

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
		"Number":post.Number,
	})

	if err != nil {
		log.Fatalf("failed to add a post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Tip, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Tip
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
		post := entity.Tip{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
			Number: doc.Data()["Number"].(int64),
		}
		posts = append(posts, post)

	}
	return posts, nil
}

func (*repo) FindOne(num int64) (*entity.Tip, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var post entity.Tip
	done := iterator.Done
	myiterator := client.Collection(collectionNAME).Where("Number", "==", num).Documents(ctx)

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
		post = entity.Tip{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
			Number: doc.Data()["Number"].(int64),
		}

	}

	return &post, nil
}

func (*repo) FindToday() (*entity.Tip, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	t:= time.Now().YearDay()
	today := int64(t)
	defer client.Close()

	var post entity.Tip
	done := iterator.Done
	myiterator := client.Collection(collectionNAME).Where("Number", "==", today).Documents(ctx)

	for {
		doc, err := myiterator.Next()
		if err != nil {
			if err == done {
				break
			}
			log.Fatalf("Failed to read documents: %v", err)
			return nil, err
		}
		post = entity.Tip{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
			Number: doc.Data()["Number"].(int64),
		}

	}

	return &post, nil
}