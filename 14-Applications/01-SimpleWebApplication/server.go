package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

var ToDoList []ToDo

type ToDo struct {
	ID          string
	Label       string
	IsCompleted bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("todolist.gohtml"))
}

var funcMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/addToDo", Index)
	router.GET("/mark/:Id", MarkDone)
	router.GET("/delete/:Id", Delete)
	router.POST("/addToDo", AddToDo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Delete from todo list
func Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	toDoID := ps.ByName("Id")
	for i := 0; i < len(ToDoList); i++ {
		if ToDoList[i].ID == toDoID {
			ToDoList = removeFromSlice(ToDoList, i)
			break
		}
	}
	http.Redirect(w, req, "/addToDo", http.StatusSeeOther)
}

// MarkDone Mark todo as read
func MarkDone(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	toDoID := ps.ByName("Id")
	for i := 0; i < len(ToDoList); i++ {
		if ToDoList[i].ID == toDoID {
			if ToDoList[i].IsCompleted {
				ToDoList[i].IsCompleted = false
			} else {
				ToDoList[i].IsCompleted = true
			}
			break
		}
	}
	http.Redirect(w, req, "/addToDo", http.StatusSeeOther)
}

// Index entery page
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	if req.URL.Path == "/" {
		ToDoList = []ToDo{}
	}
	tpl.ExecuteTemplate(w, "todolist.gohtml", ToDoList)
}

// AddToDo Add new to do
func AddToDo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fText := req.Form["fText"][0]
	uid, _ := uuid.NewV4()
	if fText != "" {
		todo := ToDo{uid.String(), fText, false}
		ToDoList = append(ToDoList, todo)
	}
	http.Redirect(w, req, "/addToDo", http.StatusSeeOther)
}

func removeFromSlice(slice []ToDo, s int) []ToDo {
	return append(slice[:s], slice[s+1:]...)
}
