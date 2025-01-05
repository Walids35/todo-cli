package todo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"github.com/mergestat/timediff"
)

type Todo struct{
	ID int
	Description string
	CreatedAt time.Time
	IsComplete bool
}

type TodoList struct{
	Todos []Todo
}

// Add todo Function
func AddTodo(name string, todos *[]Todo){
	maxId := 0
	for _, todo := range *todos {
		maxId = max(maxId, todo.ID)
	}

	todo := Todo{
		ID: maxId + 1,
		Description: name,
		CreatedAt: time.Now(),
		IsComplete: false,
	}

	*todos = append(*todos, todo)
	SaveTodos(*todos, "data.csv")
}

// List todo Function
func ListTodos(todos []Todo){
	w := tabwriter.NewWriter(os.Stdout, 15, 5, 2, ' ', 0)

	fmt.Fprintln(w, "\nID\tDescription\tCreatedAt\tDone")
	
	for _,todo := range todos{
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n", todo.ID, todo.Description, timediff.TimeDiff(todo.CreatedAt), todo.IsComplete)
	}
	w.Flush()
}

// Mark Completed Todo Function
func MarkCompleted(todoId string, todos *[]Todo){
	id,_ := strconv.Atoi(todoId)
	for i := range *todos{
		if (*todos)[i].ID == id{
			(*todos)[i].IsComplete = true
			break
		}
	}
	SaveTodos(*todos, "data.csv")
}

// Delete Todo Function
func DeleteTodo(todoID string, todos *[]Todo){
	id,_ := strconv.Atoi(todoID)
	for i,task := range *todos{
		if task.ID == id{
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
		}
	}
	SaveTodos(*todos, "data.csv")
}

// Load Tasks from CSV File
func LoadFromCSVFile(filename string)([]Todo, error){
	data, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	reader := csv.NewReader(data)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	var todos []Todo
	for i,record := range records{
		if i == 0{
			continue
		}
		
		// Parsing Todo Attributes if not string
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid ID")
		}
		
		done, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, fmt.Errorf("Invalid done value!")
		}
		
		layout := "2006-01-02T15:04:05-07:00"
		createdAt, err := time.Parse(layout, record[2])
		if err != nil{
			return nil, fmt.Errorf("Invalid Time format!")
		}

		todo := Todo{
			ID: id,
			Description: record[1],
			CreatedAt: createdAt,
			IsComplete: done,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Save Todos to the CSV file
func SaveTodos(todos []Todo, filename string) error{
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to open the file")
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID","Description","CreatedAt","Done"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %v", err)
	}

	for _,todo := range todos {
		row := []string{
			strconv.Itoa(todo.ID),
			todo.Description,
			todo.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(todo.IsComplete),
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %v", err)
		}
	}
	
	return nil
}
