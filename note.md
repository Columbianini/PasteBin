1. Go module: (`go.mod` file under project directory)
- Decide an identifier of your project (must be unique, by convention, based on your own URL, e.g. snippet.haha.net)
- Turn project directory into a module by `go mod init {identifier}`
2. Route logic
- with / at the end: subtree (will match all the file under the folder, unless {$}); without / at the end: exact path
- exact path will override the subtree. However, without exact path, exact path will be redirected to subtree. /foo -> /foo/
- wildcard route patterns: `/products/{category}/item/{itemID}`, you could fetch the value in string by `r.PathValue("category")`, wildcard segment `{path}` must not contain /, otherwise use `{path...}`


