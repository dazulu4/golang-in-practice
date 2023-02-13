package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

type taskList struct {
	tasks []*task
}

func (tl *taskList) addToTaskList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeToTaskList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) printTaskList() {
	for _, task := range tl.tasks {
		fmt.Println("-------------------------------")
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
}

func (tl *taskList) printTaskListCompleted() {
	for _, task := range tl.tasks {
		if task.completed {
			fmt.Println("-------------------------------")
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}
	}
}

func main() {
	// fmt.Printf("%+v\n", *t)
	// t.markCompleted()
	// t.updateName("Finalizar mi curso de go")
	// t.updateDescription("Completar mi curso cuanto antes")
	// fmt.Printf("%+v\n", *t)

	t1 := &task{
		name:        "Completar mi curso de go",
		description: "Completar mi curso de go el domingo",
	}
	t2 := &task{
		name:        "Completar mi curso de python",
		description: "Completar mi curso de python el domingo",
	}
	tList := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	// fmt.Println(tList.tasks[0])

	t3 := &task{
		name:        "Completar mi curso de node",
		description: "Completar mi curso de node el domingo",
	}
	tList.addToTaskList(t3)
	// fmt.Println(len(tList.tasks))
	// tList.removeToTaskList(1)
	// fmt.Println(len(tList.tasks))

	// for i := 0; i < len(tList.tasks); i++ {
	// 	fmt.Println("Index:", i, "Name:", tList.tasks[i].name)
	// }

	// for index, task := range tList.tasks {
	// 	fmt.Println("Index:", index, "Name:", task.name)
	// }

	t1.markCompleted()

	// fmt.Print("\n\n=== Tareas ===")
	// tList.printTaskList()
	// fmt.Println("\n\n=== Tareas completadas ===")
	// tList.printTaskListCompleted()

	t4 := &task{
		name:        "Completar mi curso de Java",
		description: "Completar mi curso de Java el domingo",
	}
	t5 := &task{
		name:        "Completar mi curso de C#",
		description: "Completar mi curso de C# el domingo",
	}

	tList2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapTask := make(map[string]*taskList)
	mapTask["Nestor"] = tList
	mapTask["Ricardo"] = tList2

	fmt.Print("\n\n=== Tareas Nestor ===")
	mapTask["Nestor"].printTaskList()
	fmt.Println("\n\n=== Tareas Ricardo ===")
	mapTask["Ricardo"].printTaskList()

}
