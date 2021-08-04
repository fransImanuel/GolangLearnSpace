package repository

import (
	BelajarGolangDatabase "BelajarGolangDatabase"
	"BelajarGolangDatabase/entity"
	"context"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentRepository(BelajarGolangDatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}

	result, err :=  commentRepository.Insert(ctx, comment)
	if err!=nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T){
	commentRepository := NewCommentRepository(BelajarGolangDatabase.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 35)
	if err!=nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T){
	commentRepository := NewCommentRepository(BelajarGolangDatabase.GetConnection())

	comments , err := commentRepository.FindAll(context.Background())
	if err!=nil {
		panic(err)
	}
	
	for _,comment := range comments {
		fmt.Println(comment)
	}
}