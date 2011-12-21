package main

import(
    "fmt"
    "io/ioutil"
    "http"
    "os"
    "regexp"
)

const lenPath = len("/view/")
var titlValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() os.Error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, os.Error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    
    if err != nil {
        return nil, err
    }
    
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p, err := loadPage(title)
    if err == nil {
        fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
    } else {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    
    fmt.Fprintf(w, "<h1>Editing %s</h1>" +
        "<form action=\"/save/%s\" method=\"POST\">" +
        "<textarea name=\"body\">%s</textarea><br />" +
        "<input type=\"submit\" value=\"Save\" />" +
        "</form>",
        p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    
    http.ListenAndServe(":8080", nil)
}
