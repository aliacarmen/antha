---
layout: default
type: api
navgroup: docs
shortname: antha/parser
title: antha/parser
apidocs:
  published: 2014-11-14
  antha_version: 0.0.1
  package: antha/parser
---
# parser
--
    import "."

Package parser implements a parser for Go source files. Input may be provided in
a variety of forms (see the various Parse* functions); the output is an abstract
syntax tree (AST) representing the Go source. The parser is invoked through one
of the Parse* functions.

## Usage

#### func  ParseDir

```go
func ParseDir(fset *token.FileSet, path string, filter func(os.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)
```
ParseDir calls ParseFile for all files with names ending in ".go" in the
directory specified by path and returns a map of package name -> package AST
with all the packages found.

If filter != nil, only the files with os.FileInfo entries passing through the
filter (and ending in ".go") are considered. The mode bits are passed to
ParseFile unchanged. Position information is recorded in fset.

If the directory couldn't be read, a nil map and the respective error are
returned. If a parse error occurred, a non-nil but incomplete map and the first
error encountered are returned.

#### func  ParseExpr

```go
func ParseExpr(x string) (ast.Expr, error)
```
ParseExpr is a convenience function for obtaining the AST of an expression x.
The position information recorded in the AST is undefined. The filename used in
error messages is the empty string.

#### func  ParseFile

```go
func ParseFile(fset *token.FileSet, filename string, src interface{}, mode Mode) (f *ast.File, err error)
```
ParseFile parses the source code of a single Go source file and returns the
corresponding ast.File node. The source code may be provided via the filename of
the source file, or via the src parameter.

If src != nil, ParseFile parses the source from src and the filename is only
used when recording position information. The type of the argument for the src
parameter must be string, []byte, or io.Reader. If src == nil, ParseFile parses
the file specified by filename.

The mode parameter controls the amount of source text parsed and other optional
parser functionality. Position information is recorded in the file set fset.

If the source couldn't be read, the returned AST is nil and the error indicates
the specific failure. If the source was read but syntax errors were found, the
result is a partial AST (with ast.Bad* nodes representing the fragments of
erroneous source code). Multiple errors are returned via a scanner.ErrorList
which is sorted by file position.

#### type Mode

```go
type Mode uint
```

A Mode value is a set of flags (or 0). They control the amount of source code
parsed and other optional parser functionality.

```go
const (
	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                    // stop parsing after import declarations
	ParseComments                                  // parse comments and add them to AST
	Trace                                          // print a trace of parsed productions
	DeclarationErrors                              // report declaration errors
	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)
```
