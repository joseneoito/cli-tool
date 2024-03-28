package main
import(
    "fmt"
    "net/http"
    "sync"
    "encoding/json"
)

type Todo struct {
    UserId int `json: "userId"`
    Id int `json: "id"`
    Title string `json: "title"`
    Completed bool `json: "completed"`
}

func main(){
    var wg sync.WaitGroup
    todoChannel := make(chan Todo,20);
    for i :=2; i<=40; i+=2{
        wg.Add(1)
        go fetchTodo(i, &wg, todoChannel)
    }
    go func(){
        wg.Wait()
        close(todoChannel)
    }()
    for todo := range todoChannel{
        fmt.Printf("\n Title is : %s and completed status is: %t", todo.Title, todo.Completed)
    }
}

func fetchTodo(id int, wg *sync.WaitGroup, ch chan <-Todo){
    defer wg.Done()
    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id)
    resp ,err := http.Get(url)
    if err != nil{
        fmt.Printf("faield to fetch data %d , \n error is %s", id, err)
    }
    defer resp.Body.Close()
    var todo Todo
    if err := json.NewDecoder(resp.Body).Decode(&todo); err !=nil{
        fmt.Printf("error decoding resp data %d: \n error is %s: ", id, err)
    } 
    ch <-todo
}
