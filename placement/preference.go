package placement

import "github.com/FrosTiK-SD/models-go/student"

type Preference struct {
	Slot        Slot
	Student     student.Student
	Application []Application
}
