// antha/cmd/anthadoc/handlers.go: Part of the Antha language
// Copyright (C) 2014 The Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@antha-lang.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 1 Royal College St, London NW1 0NH UK

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The /doc/codewalk/ tree is synthesized from codewalk descriptions,
// files named $GOROOT/doc/codewalk/*.xml.
// For an example and a description of the format, see
// http://golang.org/doc/codewalk/codewalk or run anthadoc -http=:6060
// and see http://localhost:6060/doc/codewalk/codewalk .
// That page is itself a codewalk; the source code for it is
// $GOROOT/doc/codewalk/codewalk.xml.

package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/antha-lang/antha-tools/anthadoc"
	"github.com/antha-lang/antha-tools/anthadoc/redirect"
	"github.com/antha-lang/antha-tools/anthadoc/vfs"
)

var (
	pres *anthadoc.Presentation
	fs   = vfs.NameSpace{}
)

func registerHandlers(pres *anthadoc.Presentation) {
	if pres == nil {
		panic("nil Presentation")
	}
	http.HandleFunc("/doc/codewalk/", codewalk)
	http.Handle("/doc/play/", pres.FileServer())
	http.Handle("/robots.txt", pres.FileServer())
	http.Handle("/", pres)
	http.Handle("/pkg/C/", redirect.Handler("/cmd/cgo/"))
	redirect.Register(nil)
}

func readTemplate(name string) *template.Template {
	if pres == nil {
		panic("no global Presentation set yet")
	}
	path := "lib/anthadoc/" + name

	// use underlying file system fs to read the template file
	// (cannot use template ParseFile functions directly)
	data, err := vfs.ReadFile(fs, path)
	if err != nil {
		log.Fatal("readTemplate: ", err)
	}
	// be explicit with errors (for app engine use)
	t, err := template.New(name).Funcs(pres.FuncMap()).Parse(string(data))
	if err != nil {
		log.Fatal("readTemplate: ", err)
	}
	return t
}

func readTemplates(p *anthadoc.Presentation, html bool) {
	p.PackageText = readTemplate("package.txt")
	p.SearchText = readTemplate("search.txt")

	if html || p.HTMLMode {
		codewalkHTML = readTemplate("codewalk.html")
		codewalkdirHTML = readTemplate("codewalkdir.html")
		p.CallGraphHTML = readTemplate("callgraph.html")
		p.DirlistHTML = readTemplate("dirlist.html")
		p.ErrorHTML = readTemplate("error.html")
		p.ExampleHTML = readTemplate("example.html")
		p.GodocHTML = readTemplate("anthadoc.html")
		p.ImplementsHTML = readTemplate("implements.html")
		p.MethodSetHTML = readTemplate("methodset.html")
		p.PackageHTML = readTemplate("package.html")
		p.SearchHTML = readTemplate("search.html")
		p.SearchDocHTML = readTemplate("searchdoc.html")
		p.SearchCodeHTML = readTemplate("searchcode.html")
		p.SearchTxtHTML = readTemplate("searchtxt.html")
		p.SearchDescXML = readTemplate("opensearch.xml")
	}
}
