package cmd

import "time"

type Reminder struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

var description string
var dueDate string

var addCmd = &cobra.Command {
	Use: "add",
	Short: "Add a new reminder",
	Run: func(cmd *cobra.Command, args []string) {
		addReminder(description, dueDate)
	},
}

func addReminder(description, dueDate string) {
	date, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		fmt.Println("Invalid date format. use YYYY-MM-DD.")
		return
	}

	reminders := readReminders()
	id := len(reminders) + 1

	reminder := Reminder {
		ID : id,
		Description: description,
		DueDate: date,
	}
}
