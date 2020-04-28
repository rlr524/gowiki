package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

// this is a method named save that takes as its receiver "p" which is a pointer to the "Page" struct (an object literal)
// that we created above. It takes no parameters and returns a value of type error; inside it we create a variable of "filename"
// using the type inference operator ":=" where we concatenate the "Title" of the "Page" object with ".txt" to create a
// file name of page-title.txt
// the method returns type error because that is the return type of WriteFile() which is a standard library function,
// imported in io/ioutil above, that writes a byte slice to a file (a slice is analogous to arrays and are built on top of the array type)
// (a slice in go, unlike an array, does not require a literal count upon declaration)
// It returns the error value in order to allow the
// application to handle it should something go wrong while writing the file. If there are no errors, this Page.save()
// method will return nil which is the zero-value for pointers, interfaces and some other types.
// the octal integer literal 0600 passed as the third parameter to WriteFile indicates that the file should be
// created with read/write permission for the current user only
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// the blank identifier (the _) below in the body variable is used to throw away the error return value, in essence
// assigning the value to nothing; in order to handle actual errors (e.g. the file doesn't exist) we use the if statement
// to determine if an err value is not nil and to return the value in that case

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
}
