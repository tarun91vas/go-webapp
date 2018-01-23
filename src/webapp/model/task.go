package model

//Task defines task components
type Task struct {
	ID          int
	Name        string
	Description string
	Done        bool
}

var tasks []Task

func init() {
	tasks = []Task{
		{
			ID:          1,
			Name:        "Communication",
			Description: "Emails follow up",
			Done:        false,
		},
		{
			ID:          2,
			Name:        "Misc",
			Description: "Book tickets",
			Done:        false,
		},
		{
			ID:          3,
			Name:        "Misc",
			Description: "Bill Payments",
			Done:        false,
		},
		{
			ID:          4,
			Name:        "Meeting",
			Description: "ES as a service",
			Done:        true,
		},
	}
}

//GetTasks fetch current list of tasks
func GetTasks() []Task {
	return tasks
}

//AddTask will add a task
func AddTask(t Task) {
	tasks = append(tasks, t)
}
