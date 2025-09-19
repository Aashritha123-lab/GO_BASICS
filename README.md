# GO_BASICS
Creating Rest API's using GO

1.We are creating an API for car inventory which will have to create a details of car , get the details of car, Delete the car details
2.For storing details of car like Name,Model,Brand,Year of Manufacturer,Price we may use struct data type

type Car struct {
	ID      int
	Name    string
	Model   string
	Company string
	Year    int
	Price   float64
}

3. Making use of map for in memory databse
   var Cars = make(map[int]Car)
4. Creating mutex to avoid unintented behaviour of concurrency
   var mu sync.Mutex
5. Create a handler function which will Parameters http.Request and http.ResponseWriter 
