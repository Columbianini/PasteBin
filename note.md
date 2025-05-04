1. Go module: (`go.mod` file under project directory)
- Decide an identifier of your project (must be unique, by convention, based on your own URL, e.g. snippet.haha.net)
- Turn project directory into a module by `go mod init {identifier}`
2. Route logic
- with / at the end: subtree (will match all the file under the folder, unless {$}); without / at the end: exact path
- exact path will override the subtree. However, without exact path, exact path will be redirected to subtree. /foo -> /foo/
- wildcard route patterns: `/products/{category}/item/{itemID}`, you could fetch the value in string by `r.PathValue("category")`, wildcard segment `{path}` must not contain /, otherwise use `{path...}`
3. Go interface
- use duck typing (implicit implementation): e.g. `http.ResponseWriter` is an interface which contains a function `Write`. `io.Write` is an interface which only has a function `Write`. In this case, `http.ResponseWriter` is of type `io.Write`. So for any function that has `io.Write` parameter, you could pass in `http.ResponseWriter` value. For function `io.WriteString`, it will finally call `w.Write`, so we could replace `w.Write([]byte("Hello world"))` with `io.WriteString(w, "Hello World")`
4. project structure
```
.
├── cmd
│   └── web % application-specific code
│       ├── handlers.go
│       └── main.go
├── go.mod
├── internal % normal code shared by all applications under the parent directory of internal
├── note.md
└── ui
    ├── html
    └── static
```
5. template
- use `{{define "{TemplateName}"}}...{some html contents or string}...{{end}}` to set a named template value
- use `{{template {TemplateName} .}}` or `{{block {TemplateName} .}}...{{end}}`(can provide default value inside if not then not render if fail) to declare a template, base template does not need to declare but need to mention in `(t *template.Template) ExecuteTemplate(wr io.Writer, name string, data any) error`
- in the code, you should first parse all related templates by `func template.ParseFiles(filenames ...string) (*template.Template, error)`
6. golang fileserver
- initialize with `http.FileServer(http.Dir(path relative to project directory))`
- start with `mux.Handle(http.StripPrefix)`
7. golang command line variables
- use flag package
- declare a command line args by `addr:=flag.String("addr", ":4000", "HTTP network address")`
- then parse in the main function by `flag.Parse()`

