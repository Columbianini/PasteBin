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
8. golang logger
- you can use by `logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource=true}))`
9. dependencies injection: allow all functions access to the same object
- all in one package: create a struct `application` in `main.go`, Add it as method receiver to existing functions
- in different package: create a package exporting struct `application`, enclose the exisitng function into a closure with parameter `application` (i.e. function that returns a function) 
10. centralized error handling
- use package `runtime/debug` to write debug trace into log
11. `defer` keyword
- In Go, defer schedules a function call to run ​​just before the surrounding function returns​​, regardless of how it returns (normally, via return, or due to a panic). Much like finally in python. However, `defer` won't run if application is terminated by a signal interrupt (i.e. Ctrl+C or by os.Exit(1))
12. DBA: set up a database
- install db server (mssql, mysql, etc)
- change db server config to allow listening to all IPs, `sudo systemctl start mysql` to update it
- create new user and grant privileges by `CREATE USER 'username'@['host'|'%']; ALTER USER 'usename'@['host'|'%'] IDENTIFIED BY 'pass'; GRANT [SELECT,INSERT,UPDATE,DELETE|ALL PRIVILEGES] ON ['database'|*].['table'|*] TO 'username'@'%';`
13. golang to display a struct
- `fmt.FPrintf(w, "%+v", snippet)`
14. go template method
- `{{if .Snippet}} C1 {{else}} C2 {{end}}`: render C1 if we have non-empty Snippet field in data else render C2
- `{{with .Snippet}} C1 {{else}} C2 {{end}}`: same as above, but we will set dot = .Snippet in the with block to save some typings
- `{{range .Snippets}} C1 {{else}} C2 {{end}}`: for each Snippet in Snippets, render C1. If Snippets is empty, then render C2
15. [Embed struct in struct](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)
- Embeding struct will auto inherit the property and method of the embedded struct
16. Check and Grant permissions in mysql
```sql
SHOW GRANTS FOR 'web'@'192.168.10.1';
GRANT SELECT, INSERT ON `mydb`.* TO 'web'@'192.168.10.1' 
```
17. Generate TLS certificate for HTTPS server
- create a tls folder under workspace
- `go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2049 --host=localhost`
18. Named type in golang
- we can create a new type named contextKey with underlying type string
- the pattern is always used to avoid key collision
- you can think of new type as a completely different types without inheriting the underlying type's method 
19. Embedded filesystem
- bundle files directly into executable program
- how?
    - put a go file under the directory e.g. `ui/efs.go`
    - add comment directive in the file immediately above the variable in which you want to store the embedded files
    - The directive has the format `go:embed "<path>"`. The path is relative to the .go file containing the directive, so — in our case — `go:embed "static"` embeds the directory
    `ui/static` from our project
    - You can only use the go:embed directive on global variables at package level, not within functions or methods.
    -  The embedded file system is always rooted in the directory which contains the go:embed
    directive. So, in the example above, our Files variable contains an embed.FS embedded
    filesystem and the root of that filesystem is our ui directory.
20. Unit-Test
- standard practice: write your tests in `*_test.go` files which live directly alongside the code you are testing
- Your unit tests are contained in a normal Go function with the signature `func(*testing.T)`
- To be a valid unit test, the name of the function must begin with the word `Test`.
- You can use the `t.Errorf()` function to make a test as failed and log a descriptive message about the failure. It's important to note that calling `t.Errorf()` doesn't stop execution of your tests
- use `go test ./cmd/web` command to run all the tests in our `cmd/web` package like
- Table-driven tests: a way to run multiple test cases
