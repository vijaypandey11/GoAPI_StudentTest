/*
 * Student Swagger
 *
 * Student Test Rest API using Swagger
 *
 * API version: 1.0.0
 * Contact: vijay.pandey1@globallogic.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/getStudentDetails/",
		Index,
	},

	Route{
		"GetAverageMarks",
		strings.ToUpper("Get"),
		"/getStudentDetails/class/getAvgMarks",
		GetAverageMarks,
	},

	Route{
		"GetMaximumMarks",
		strings.ToUpper("Get"),
		"/getStudentDetails/class/getMaxMarks",
		GetMaximumMarks,
	},

	Route{
		"GetMinimumMarks",
		strings.ToUpper("Get"),
		"/getStudentDetails/class/getMinMarks",
		GetMinimumMarks,
	},

	Route{
		"GetStudentsCount",
		strings.ToUpper("Get"),
		"/getStudentDetails/class/getStudentCount",
		GetStudentsCount,
	},

	Route{
		"AddStudent",
		strings.ToUpper("Post"),
		"/getStudentDetails/student",
		AddStudent,
	},

	Route{
		"DeleteStudent",
		strings.ToUpper("Delete"),
		"/getStudentDetails/student/{id}",
		DeleteStudent,
	},

	Route{
		"FindStudentsByAge",
		strings.ToUpper("Get"),
		"/getStudentDetails/student/findByAge",
		FindStudentsByAge,
	},

	Route{
		"FindStudentsByName",
		strings.ToUpper("Get"),
		"/getStudentDetails/student/findByName",
		FindStudentsByName,
	},

	Route{
		"FindStudentsByStatus",
		strings.ToUpper("Get"),
		"/getStudentDetails/student/findByStatus",
		FindStudentsByStatus,
	},

	Route{
		"GetStudentById",
		strings.ToUpper("Get"),
		"/getStudentDetails/student/{id}",
		GetStudentById,
	},

	Route{
		"UpdateStudent",
		strings.ToUpper("Put"),
		"/getStudentDetails/student",
		UpdateStudent,
	},

	Route{
		"UpdateStudentById",
		strings.ToUpper("Post"),
		"/getStudentDetails/student/{id}",
		UpdateStudentById,
	},
}