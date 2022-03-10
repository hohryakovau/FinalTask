package main

import (
	"FinalTask/src"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks/Циклическая ротация", firstAPP)            // Устанавливаем роутер
	http.HandleFunc("/tasks/Чудные вхождения в массив", secondAPP)     // Устанавливаем роутер
	http.HandleFunc("/tasks/Проверка последовательности", thirdApp)    // Устанавливаем роутер
	http.HandleFunc("/tasks/Поиск отсутствующего элемента", fourthApp) // Устанавливаем роутер
	http.HandleFunc("/tasks/tasks/", tasksApp)                         // Устанавливаем роутер
	err := http.ListenAndServe(":908", nil)                            // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Для первой задачи
type Res struct {
	UserName string `json:"user_name"`
	Task     string `json:"task"`
	Result   Result `json:"results"`
}

type Result struct {
	Payload []json.RawMessage `json:"payload"`
	Results [10][]int         `json:"results"` // результаты полученные при решении задачи
}

// Для второй, третьей,четвертой задачи
type Res2 struct {
	UserName string  `json:"user_name"`
	Task     string  `json:"task"`
	Result   Result2 `json:"results"`
}

type Result2 struct {
	Payload []json.RawMessage `json:"payload"`
	Results [10]int           `json:"results"` // результаты полученные при решении задачи
}

func firstAPP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://116.203.203.76:3000/tasks/Циклическая ротация")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var A []int
	var K int
	var dat []json.RawMessage
	var arg []json.RawMessage

	json.Unmarshal([]byte(body), &dat)

	res := &Res{UserName: "alexey_khokhryakov", Task: "Циклическая ротация"}

	var ArrayResult [10][]int

	for i, v := range dat {
		json.Unmarshal(v, &arg)
		json.Unmarshal(arg[0], &A)
		json.Unmarshal(arg[1], &K)
		result := src.First(A, K)

		ArrayResult[i] = result
	}
	res.Result.Payload = dat
	res.Result.Results = ArrayResult

	userVar2, _ := json.Marshal(res)

	req, err := http.Post("http://116.203.203.76:3000/tasks/solution", "application/json", bytes.NewBuffer(userVar2))
	if err != nil {
		fmt.Errorf("post request failed: %v", err)
	}
	defer req.Body.Close()

	reqq, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, string(reqq))
	fmt.Println(string(reqq))
}

func secondAPP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://116.203.203.76:3000/tasks/Чудные вхождения в массив")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var A []int
	var dat []json.RawMessage
	var arg []json.RawMessage

	json.Unmarshal([]byte(body), &dat)

	res := &Res2{UserName: "alexey_khokhryakov", Task: "Чудные вхождения в массив"}

	var ArrayResult [10]int

	for i, v := range dat {
		json.Unmarshal(v, &arg)
		json.Unmarshal(arg[0], &A)
		result := src.Second(A)
		ArrayResult[i] = result
	}
	res.Result.Payload = dat
	res.Result.Results = ArrayResult
	userVar2, _ := json.Marshal(res)

	req, err := http.Post("http://116.203.203.76:3000/tasks/solution", "application/json", bytes.NewBuffer(userVar2))
	if err != nil {
		fmt.Errorf("post request failed: %v", err)
	}
	defer req.Body.Close()
	reqq, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, string(reqq))
	fmt.Println(string(reqq))
}

func thirdApp(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://116.203.203.76:3000/tasks/Проверка последовательности")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var A []int

	var dat []json.RawMessage
	var arg []json.RawMessage

	json.Unmarshal([]byte(body), &dat)

	res := &Res2{UserName: "alexey_khokhryakov", Task: "Проверка последовательности"}
	//
	var ArrayResult [10]int
	//
	for i, v := range dat {
		json.Unmarshal(v, &arg)
		json.Unmarshal(arg[0], &A)
		result := src.Third(A)
		ArrayResult[i] = result
	}
	res.Result.Payload = dat
	res.Result.Results = ArrayResult
	//
	//fmt.Println(res)
	userVar2, _ := json.Marshal(res)

	req, err := http.Post("http://116.203.203.76:3000/tasks/solution", "application/json", bytes.NewBuffer(userVar2))
	if err != nil {
		fmt.Errorf("post request failed: %v", err)
	}

	defer req.Body.Close()
	reqq, _ := ioutil.ReadAll(req.Body)

	fmt.Fprintf(w, string(reqq))
	fmt.Fprintf(w, string("/n"))
	fmt.Println(string(reqq))
}

func fourthApp(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://116.203.203.76:3000/tasks/Поиск отсутствующего элемента")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Fprintf(w, string(body))
	var A []int
	//
	var dat []json.RawMessage
	var arg []json.RawMessage
	//
	json.Unmarshal([]byte(body), &dat)
	//
	res := &Res2{UserName: "alexey_khokhryakov", Task: "Поиск отсутствующего элемента"}
	//
	var ArrayResult [10]int
	//
	for i, v := range dat {
		json.Unmarshal(v, &arg)
		json.Unmarshal(arg[0], &A)
		result := src.Fourth(A)
		ArrayResult[i] = result
	}
	res.Result.Payload = dat
	res.Result.Results = ArrayResult

	userVar2, _ := json.Marshal(res)

	req, err := http.Post("http://116.203.203.76:3000/tasks/solution", "application/json", bytes.NewBuffer(userVar2))
	if err != nil {
		fmt.Errorf("post request failed: %v", err)
	}

	defer req.Body.Close()
	reqq, _ := ioutil.ReadAll(req.Body)

	fmt.Fprintf(w, string(reqq))
	fmt.Println(string(reqq))

}

func tasksApp(w http.ResponseWriter, r *http.Request) {
	firstAPP(w, r)
	secondAPP(w, r)
	thirdApp(w, r)
	fourthApp(w, r)
}
