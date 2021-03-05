package main

// TEST WILL FAIL
// Tried to learn how to test endpoints but the following code does not work.
// The request is not made to the endpoint. The overall idea was to get the input
// response and put in an file and compare it to the output response (that I get
// from Postman) for easier comparison.
// I also wanted to try to test adding a document to MongoDB but it required me
// to build a mock DB which is quite out of my scope.

// func TestGetAllJobsEndpoint(t *testing.T) {
// 	req, err := http.NewRequest("GET", "http://localhost:5000/api/jobs", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetAllJobsEndpoint)
// 	// fmt.Println("REQ", req)
// 	// fmt.Println("RR", rr)
// 	// fmt.Println("HANDLER", handler)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	file, _ := json.MarshalIndent(rr.Body.String(), "", " ")
// 	_ = ioutil.WriteFile("jobs_input.json", file, 0644)

// 	input, err := ioutil.ReadFile("jobs_input.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	output, err2 := ioutil.ReadFile("jobs_output.json")
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	same := bytes.Equal(input, output)
// 	// fmt.Println(same)

// 	if same != true {
// 		t.Errorf("Handler returned unexpected body")
// 	}

// 	// Check the response body is what we expect.
// 	expected := ``
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
