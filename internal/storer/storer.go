package storer

import (
	"context"
	"fmt"
	"sticky-notes-go-backend/internal/model"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MyStorerDb struct {
	db *mongo.Client
}


func NewDbStorer(db *mongo.Client) *MyStorerDb {
	return &MyStorerDb{
		db: db,
	}
}

func (ds *MyStorerDb) CreateNotes(ctx context.Context, m model.StickyNote) (*model.StickyNote, error) {

	if m.ID == ""{
		m.ID = uuid.NewString()
	}

	_, err := ds.db.Database("go-backend").Collection("notes").InsertOne(ctx, m)

	if err != nil {
		return nil, fmt.Errorf("Error in creating notes %w", err)
	}

	return &m, nil

}

func (ds *MyStorerDb) GetNoteById(ctx context.Context,id string) (*model.StickyNote,error) {

	var note model.StickyNote

	filter := bson.M{"_id" : id}

	err := ds.db.Database("go-backend").Collection("notes").FindOne(ctx,filter).Decode(&note)

	if err != nil {
		return nil,fmt.Errorf("Error getting teh particular note %w",err)
	}

	return &note,nil
}

func (ds *MyStorerDb) GettALLNotes(ctx context.Context) ([]*model.StickyNote,error){

	var notes []*model.StickyNote

	filter := bson.M{"status": bson.M{"$ne": "trash"}}

	get,err := ds.db.Database("go-backend").Collection("notes").Find(ctx,filter)

	if err != nil {
		return nil ,fmt.Errorf("ERror getting all the notes %w",err)
	}

	err = get.All(ctx,&notes)

	if err != nil {
		return nil ,fmt.Errorf("Error in notes %w",err)
	}

	return notes,nil

}

func (ds *MyStorerDb) EditNote(ctx context.Context, id string, data model.StickyNote) (model.StickyNote,error) {


	update := bson.M{
	"$set": bson.M{
		"title":     data.Title,
		"content":   data.Content,
		"color":     data.Color,
		"pinned":    data.Pinned,
		"updatedAt": time.Now(),
	},
}
     
	var updatedNote model.StickyNote

    filter := bson.M{"_id" : id}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := ds.db.Database("go-backend").Collection("notes").FindOneAndUpdate(ctx,filter,&update, opts).Decode(&updatedNote)

	if err != nil {
		return model.StickyNote{},fmt.Errorf("Error editing note %w",err)
	}

	return updatedNote,nil

}

func (ds *MyStorerDb) GetTrashNotes(ctx context.Context) ([]*model.StickyNote,error) {

	var n []*model.StickyNote

	filter := bson.M{"status":"trash"}

	get,err := ds.db.Database("go-backend").Collection("notes").Find(ctx,filter)

	if err != nil {
		return nil,err
	}

	err = get.All(ctx,&n)

	return n,nil
}


func (ds *MyStorerDb) RestoreNote(ctx context.Context, id string) (error) {


	filter := bson.M{"_id" : id}

	update := bson.M{
	"$set": bson.M{
		"status" : "active" }}


	_,err := ds.db.Database("go-backend").Collection("notes").UpdateOne(ctx,filter,update)

	if err != nil {
		return  fmt.Errorf("Error Restoring note %w",err)
	}

	return nil

}


func (ds *MyStorerDb) DeleteNote(ctx context.Context, id string) error {

	filter := bson.M{"_id" : id}

	update := bson.M{"$set":
		bson.M{"status": "trash",
	}}

	_,err := ds.db.Database("go-backend").Collection("notes").UpdateOne(ctx,filter,update)

	if err != nil {
		return fmt.Errorf("Error Delteing the note %w",err)
	}

	return nil

}

func (ds *MyStorerDb) PermanentDelete(ctx context.Context , id string) error {

	filter := bson.M{"_id" : id}

	_,err := ds.db.Database("go-backend").Collection("notes").DeleteOne(ctx,filter)

	if err != nil {
		return fmt.Errorf("Error Delteing the note %w",err)
	}

	return nil

}

func (ds *MyStorerDb) TogglePin(ctx context.Context,id string) (error) {


	var note model.StickyNote
	
	filter := bson.M{"_id" : id}

	err := ds.db.Database("go-backend").Collection("notes").FindOne(ctx,filter).Decode(&note)

	if err != nil {
		return fmt.Errorf("error finding note to toggle pin: %w", err)
	}

	newPinnedValue := !note.Pinned

	
	update := bson.M{
		"$set" :
		bson.M{ 
			"pinned" : newPinnedValue,
		},
	}

   _,err = ds.db.Database("go-backend").Collection("notes").UpdateOne(ctx,filter,&update)

   if err != nil {
	return err
   }

   return nil

}
