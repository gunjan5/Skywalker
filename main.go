package main

/* Copyright (C) 2015  Gunjan Patel
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation,
 * either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details. You should have received a copy of
 * the GNU Affero General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import (
	"flag"
	_ "fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	_ "html/template"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	yell "github.com/gunjan5/Skywalker/services/yellingService"
)

//var templates = template.Must(template.ParseFiles("index.html"))
var staticPath = flag.String("staticPath", "static/", "/Volumes/Other/Dropbox/Skywalker/static/")

func main() {
	router := mux.NewRouter()

	ctx := context.Background()
	svc := yell.YellingService{}

	yellHandler := httptransport.Server{
		Context:            ctx,
		Endpoint:           yell.MakeYellEndpoint(svc),
		DecodeRequestFunc:  yell.DecodeYellRequest,
		EncodeResponseFunc: yell.EncodeResponse,
	}

	router.HandleFunc("/", RootHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(*staticPath))))
	http.Handle("/yell", yellHandler)
	go http.ListenAndServe(":8080", nil)
	go log.Fatal(http.ListenAndServe(":8081", router))
}
