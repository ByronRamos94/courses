package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) add(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *taskList) remove(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)

}

func (t *taskList) print() {
	for _, task := range t.tasks {
		fmt.Println("Name: ", task.name)
		fmt.Println("Description: ", task.description)
	}
}
func (t *taskList) printCompleted() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Name: ", task.name)
			fmt.Println("Description: ", task.description)
		}
	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true

}
func (t *task) updateDescription(description string) {
	t.description = description
}
func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t := &task{
		name:        "complete my course",
		description: "view 22 caps of the golang course",
		completed:   false,
	}
	t1 := &task{
		name:        "complete my course of python",
		description: "view 22 caps of the python course",
		completed:   false,
	}
	t2 := &task{
		name:        "complete my course of javascript",
		description: "view 22 caps of the javascript course",
		completed:   false,
	}
	t3 := &task{
		name:        "complete my course of mongo",
		description: "view 22 caps of the mongo course",
		completed:   false,
	}
	t4 := &task{
		name:        "complete my course of java",
		description: "view 22 caps of the java course",
		completed:   false,
	}
	t5 := &task{
		name:        "complete my course of c#",
		description: "view 22 caps of the c# course",
		completed:   false,
	}
	/* fmt.Printf("%+v\n", t)
	t.markAsCompleted()
	t.updateName("Finish my go course")
	t.updateDescription("Finish my course as soon as possible")
	fmt.Printf("%+v\n", t) */
	list := &taskList{
		tasks: []*task{
			t1, t2, t,
		},
	}
	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}
	list.add(t3)
	list.remove(1)
	list.tasks[0].markAsCompleted()
	// list.print()
	list.printCompleted()

	taskMap := make(map[string]*taskList)
	taskMap["BYRON"] = list
	taskMap["ESTUARDO"] = list2
	// first way to iterate the slice
	/*for i := 0; i < len(list.tasks); i++ {
		fmt.Println(list.tasks[i].name)
	}*/

	// second way to iterate the slice
	/* for index, task := range list.tasks {
		fmt.Println("Index", index, "name", task.name)
	} */
}
