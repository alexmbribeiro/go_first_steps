package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "text/tabwriter"
    "time"
    "github.com/mergestat/timediff"
    "github.com/spf13/cobra"
)

type Todo struct {
    Task     string
    Created  time.Time
	Done	 bool
}

const fileName = "todos.csv"

func loadTodos() ([]Todo, error) {
    var todos []Todo

    file, err := os.Open(fileName)
    if os.IsNotExist(err) {
        return todos, nil
    }
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    for _, record := range records {
        ts, _ := strconv.ParseInt(record[1], 10, 64)
		done, _ := strconv.ParseBool(record[2])
        todos = append(todos, Todo{
            Task:    record[0],
            Created: time.Unix(ts, 0),
			Done: done,
        })
    }
    return todos, nil
}

func saveTodos(todos []Todo) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, t := range todos {
        record := []string{
            t.Task,
            strconv.FormatInt(t.Created.Unix(), 10),
			strconv.FormatBool(t.Done),
        }
        writer.Write(record)
    }
    return nil
}

func addTask(task string) {
    todos, _ := loadTodos()
    todos = append(todos, Todo{Task: task, Created: time.Now(), Done: false})
    saveTodos(todos)
    fmt.Println("Tarefa adicionada:", task)
}

func listTasks(showAll bool) {
    todos, _ := loadTodos()

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    fmt.Fprintln(w, "ID\tTask\tCreated at\tDone")
    fmt.Fprintln(w, "----\t------\t---------------------------\t---------------")

    for i, t := range todos {
		if showAll {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, t.Task, timediff.TimeDiff(t.Created), strconv.FormatBool(t.Done))
		} else if !t.Done {
        	fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, t.Task, timediff.TimeDiff(t.Created), strconv.FormatBool(t.Done))
		}
    }

    w.Flush()
}

func completeTask(index int) {
    todos, _ := loadTodos()
    if index < 1 || index > len(todos) {
        fmt.Println("Índice inválido.")
        return
    }
    todos[index-1].Done = true
    saveTodos(todos)
    fmt.Println("Tarefa concluída:", todos[index-1].Task)
}

func deleteTask(index int) {
    todos, _ := loadTodos()

    if index < 1 || index > len(todos) {
        fmt.Println("Índice inválido.")
        return
    }

    removed := todos[index-1].Task
    todos = append(todos[:index-1], todos[index:]...)
    saveTodos(todos)

    fmt.Println("Tarefa removida:", removed)
}

func main() {
	var showAll bool

    var rootCmd = &cobra.Command{Use: "todo"}

    var addCmd = &cobra.Command{
        Use:   "add [tarefa]",
        Short: "Adiciona uma nova tarefa",
        Args:  cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            task := args[0]
            addTask(task)
        },
    }

    var listCmd = &cobra.Command{
        Use:   "list",
        Short: "Lista todas as tarefas",
        Run: func(cmd *cobra.Command, args []string) {
            listTasks(showAll)
        },
    }

	var completeCmd = &cobra.Command{
        Use:   "complete",
        Short: "Completar uma tarefa",
        Run: func(cmd *cobra.Command, args []string) {
        
		task_id, err := strconv.Atoi(args[0])
	    if err != nil {
            fmt.Println("Número inválido.")
            return
        }
        completeTask(task_id)
        },
    }

    var deleteCmd = &cobra.Command{
        Use:   "delete",
        Short: "Elimina uma tarefa",
        Run: func(cmd *cobra.Command, args []string) {
		task_id, err := strconv.Atoi(args[0])
	    if err != nil {
            fmt.Println("Número inválido.")
            return
        }
        deleteTask(task_id)
        },
    }

	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")

    rootCmd.AddCommand(addCmd, listCmd, completeCmd, deleteCmd)
    rootCmd.Execute()
}
