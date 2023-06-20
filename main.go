 
 package main

 import (
    "encoding/json"
    "fmt"
    "net/http"
)

 type Item struct {
	ShortDescription	string	`json:"shortDescription"`
	Price				string	`json:"price"`
 }

 type Items []Item

 type Receipt struct {
	ID				string  `json:"ID"`
	Retailer		string	`json:"retailer"`
	PurchaseDate	string  `json:"purchaseDate"`
	PurchaseTime	string	`json:"purchaseTime"`
	Total			string	`json:"total"`
	Items			Items	`json:"items"`
 }

 type receiptHandlers struct {
	store map[string]Receipt
 }

 func (h *receiptHandlers) receipts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w,r)
		return
	case "POST":
		h.post(w,r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
 }


 func (h *receiptHandlers) post(w http.ResponseWriter, r *http.Request) {

 }

 func (h *receiptHandlers) get(w http.ResponseWriter, r *http.Request) {
	testID:="7fb1377b-b223-49d9-a31a-5a02701dd310"
	//badID:="asdfj;laksdjf;"

	receiptObj:=h.store[testID]
	jsonBytes, err := json.Marshal(receiptObj)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
	w.Header().Add("content-type","application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
 }

 func newReceiptHandlers() *receiptHandlers {
	return &receiptHandlers{
		store: map[string]Receipt{
			"7fb1377b-b223-49d9-a31a-5a02701dd310": Receipt{
				ID: "7fb1377b-b223-49d9-a31a-5a02701dd310",
				Retailer: "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Total: "2.65",
				Items: []Item{
					{ShortDescription:  "Pepsi - 12-oz", Price: "1.25"},
					{ShortDescription: "Dasani", Price: "1.40"},
				},
			},
		},
	}
 }



 func main() {
	fmt.Println("Hello World")
	receiptHandlers := newReceiptHandlers()
	http.HandleFunc("/receipts/process", receiptHandlers.receipts)

	http.ListenAndServe(":8080",nil)
 }