    Go Meta Linter is an utilize  which can parallel execute wholly, all kinds of Linter for the Go language

#### Installing
There are two options for installing gometalinter.

1. Install a stable version, eg.  `go get -u gopkg.in/alecthomas/gometalinter.v2. `
 The downside is that the binary will be called gometalinter.v2.
1. Install from HEAD with:` go get -u github.com/alecthomas/gometalinter`
       This has the downside that changes to gometalinter may break.
      
#### Supported linters

- [go vet](https://golang.org/cmd/vet/) - Reports potential errors that otherwise compile.
- [go tool vet --shadow](https://golang.org/cmd/vet/#hdr-Shadowed_variables) - Reports variables that may have been unintentionally shadowed.
- [gotype](https://golang.org/x/tools/cmd/gotype) - Syntactic and semantic analysis similar to the Go compiler.
- [gotype -x](https://golang.org/x/tools/cmd/gotype) - Syntactic and semantic analysis in external test packages (similar to the Go compiler).
- [deadcode](https://github.com/tsenart/deadcode) - Finds unused code.
- [gocyclo](https://github.com/alecthomas/gocyclo) - Computes the cyclomatic complexity of functions.
- [golint](https://github.com/golang/lint) - Google's (mostly stylistic) linter.
- [varcheck](https://github.com/opennota/check) - Find unused global variables and constants.
- [structcheck](https://github.com/opennota/check) - Find unused struct fields.
- [maligned](https://github.com/mdempsky/maligned) -  Detect structs that would take less memory if their fields were
 sorted.
- [errcheck](https://github.com/kisielk/errcheck) - Check that error return values are used.
- [megacheck](https://github.com/dominikh/go-tools/tree/master/cmd/megacheck) - Run staticcheck, gosimple and unused,
 sharing work.
- [dupl](https://github.com/mibk/dupl) - Reports potentially duplicated code.
- [ineffassign](https://github.com/gordonklaus/ineffassign/blob/master/list) - Detect when assignments to *existing* 
variables are not used.
- [interfacer](https://github.com/mvdan/interfacer) - Suggest narrower interfaces that can be used.
- [unconvert](https://github.com/mdempsky/unconvert) - Detect redundant type conversions.
- [goconst](https://github.com/jgautheron/goconst) - Finds repeated strings that could be replaced by a constant.
- [gas](https://github.com/GoASTScanner/gas) - Inspects source code for security problems by scanning the Go AST.

Disabled by default (enable with `--enable=<linter>`):

- [testify](https://github.com/stretchr/testify) - Show location of failed testify assertions.
- [test](http://golang.org/pkg/testing/) - Show location of test failures from the stdlib testing module.
- [gofmt -s](https://golang.org/cmd/gofmt/) - Checks if the code is properly formatted and could not be further simplified.
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Checks missing or unreferenced package imports.
- [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - Report simplifications in code.
- [lll](https://github.com/walle/lll) - Report long lines (see `--line-length=N`).
- [misspell](https://github.com/client9/misspell) - Finds commonly misspelled English words.
- [nakedret](https://github.com/alexkohler/nakedret) - Finds naked returns.
- [unparam](https://github.com/mvdan/unparam) - Find unused function parameters.
- [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - Find unused variables.
- [safesql](https://github.com/stripe/safesql) - Finds potential SQL injection vulnerabilities.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - Statically detect bugs, both obvious
 and subtle ones.
 
 ## Quickstart
 
 Install gometalinter (see above).
 
 Install all known linters:
 
 ```
 $ gometalinter --install
 Installing:
   structcheck
   maligned
   nakedret
   deadcode
   gocyclo
   ineffassign
   dupl
   golint
   gotype
   goimports
   errcheck
   varcheck
   interfacer
   goconst
   gosimple
   staticcheck
   unparam
   unused
   misspell
   lll
   gas
   safesql
 ```
 
 Run it:
 
 ```
 $ cd example
 $ gometalinter ./...
 stutter.go:13::warning: unused struct field MyStruct.Unused (structcheck)
 stutter.go:9::warning: unused global variable unusedGlobal (varcheck)
 stutter.go:12:6:warning: exported type MyStruct should have comment or be unexported (golint)
 ```
 
 
 Gometalinter also supports the commonly seen `<path>/...` recursive path
 format. Note that this can be *very* slow, and you may need to increase the linter `--deadline` to allow linters to complete.
