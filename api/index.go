package handler
 
import ( 
    "encoding/json"
    "net/http"
)

type Person struct { 
    Name string `json:"name"` 
    Age int `json:"age"`
}

func Handler(w http.ResponseWriter, r *http.Request) { 
	person := Person{  Name: "John",  Age: 30, } 

    // Encoding - One step
    jsonStr, err := json.Marshal(person) 

    if err != nil {  
        http.Error(w, err.Error(), http.StatusInternalServerError)  
        return 
    } 

    w.Write(jsonStr)
}

func Main() { 
   http.HandleFunc("/", Handler) 
   http.ListenAndServe(":8080", nil)
}