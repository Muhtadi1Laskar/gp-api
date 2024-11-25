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
    person := Person{  Name: "John",  Age: 30 } 

    // Encoding - 2 step . NewEncoder and Encode
    encoder := json.NewEncoder(w) 

    err := encoder.Encode(person) 

    if err != nil {  
        http.Error(w, err.Error(), http.StatusInternalServerError)  

        return 
   }}

func Main() { 
   http.HandleFunc("/", Handler) 
   http.ListenAndServe(":8080", nil)
}