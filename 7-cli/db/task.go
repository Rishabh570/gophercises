package db

import (
	"fmt"
	"os"
	"task/utils"

	bolt "go.etcd.io/bbolt"
)

type Task struct {
	Id   int
	Text string
}

// Adds a task to bolt DB
func AddTask(text string) int {
	var id int

	// Store the task id and text in bolt DB; uses read-write transaction
	Database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id64, _ := b.NextSequence()
		id = int(id64)
		idSlice := utils.ConvertIntToByteArray(id)
		err := b.Put(idSlice, []byte(text))
		return err
	})

	return id
}

// Lists all the open tasks
func ListTasks() ([]Task, error) {
	var tasksList []Task

	// uses read-only transaction
	err := Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// Convert id from []byte to int
			id := utils.ConvertByteArrayToInt(k)

			// build a task object
			task := Task{
				Id:   id,
				Text: string(v),
			}

			// add fetched task to the list
			tasksList = append(tasksList, task)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasksList, nil
}

// Mark a task as completed, the task is deleted from DB upon completion
func DoTask(id int) error {
	_, err := GetTask(id)
	if err != nil {
		os.Exit(1)
	}

	return Database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(utils.ConvertIntToByteArray(id))
	})
}

// Fetches the task using task id
func GetTask(taskId int) (string, error) {
	var text string

	err := Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id := utils.ConvertIntToByteArray(int(taskId))
		text = string(b.Get(id))
		return nil
	})

	if err != nil || len(text) == 0 {
		fmt.Println("No task found with provided task ID: ", taskId)
		os.Exit(1)
	}

	return text, nil
}
