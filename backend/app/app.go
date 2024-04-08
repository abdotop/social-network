package octopus

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type HandlerFunc func(*Context)
type ErrorHandlerFunc func(*Context, int)

type Route struct {
	pattern  string
	handlers []HandlerFunc
	methods  map[string]bool
}

type App struct {
	Db               *db
	routes           []*Route
	onErrorCode      ErrorHandlerFunc
	globalMiddleware []HandlerFunc
}

func New() *App {
	return new(App)
}

// func New(migration sqlite.Migrations) *App {
// 	database := sqlite.OpenDB(migration)
// 	app := &App{
// 		Db: &db{
// 			Conn: database,
// 		},
// 		routes:      make([]*Route, 0),
// 		onErrorCode: nil,
// 	}

// 	return app
// }

func (a *App) UseDb(conn *sql.DB) {
	d := &db{Conn: conn}
	a.Db = d
}

func (app *App) handle(pattern string, handlers []HandlerFunc, methods ...string) {
	methodsMap := make(map[string]bool)
	for _, method := range methods {
		methodsMap[method] = true
	}
	handlers = append(app.globalMiddleware, handlers...)
	route := &Route{pattern: pattern, handlers: handlers, methods: methodsMap}
	app.routes = append(app.routes, route)
}

func (a *App) Use(handlers ...HandlerFunc) {
	a.globalMiddleware = append(a.globalMiddleware, handlers...)
}

func (app *App) Static(path string, dir string) {
	fileServer := http.FileServer(http.Dir(dir))
	app.GET(path+"*", func(c *Context) {
		http.StripPrefix(path, fileServer).ServeHTTP(c.ResponseWriter, c.Request)
	})
}

func (app *App) GET(path string, handler ...HandlerFunc) {
	app.handle(path, handler, "GET")
}

func (app *App) PUT(path string, handler ...HandlerFunc) {
	app.handle(path, handler, "PUT")
}

func (app *App) POST(path string, handler ...HandlerFunc) {
	app.handle(path, handler, "POST")
}

func (app *App) DELETE(path string, handler ...HandlerFunc) {
	app.handle(path, handler, "DELETE")
}

func (app *App) OnErrorCode(handler ErrorHandlerFunc) {
	app.onErrorCode = handler
}

func (app *App) NotAllowed(c *Context) {
	http.Error(c.ResponseWriter, "405 Method not allowed", http.StatusMethodNotAllowed)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{ResponseWriter: w, Request: r, Db: app.Db, Values: make(map[any]any)}
	for _, route := range app.routes {
		if strings.HasSuffix(route.pattern, "*") {
			if strings.HasPrefix(r.URL.Path, strings.TrimSuffix(route.pattern, "*")) {
				if route.methods[r.Method] {
					c.handlers = route.handlers
					c.Next()
					return
				} else {
					app.NotAllowed(c)
					return
				}
			}
		} else {
			if r.URL.Path == route.pattern {
				if route.methods[r.Method] || r.Method == "OPTIONS" {
					c.handlers = route.handlers
					c.Next()
					return
				} else {
					app.NotAllowed(c)
					return
				}
			}
		}
	}
	if app.onErrorCode != nil {
		app.onErrorCode(c, http.StatusNotFound)
	} else {
		http.NotFound(w, r)
	}
}

var wg sync.WaitGroup

func checkServer(addr string) {
	resp, err := http.Get("http://" + addr)
	if err != nil {
		log.Println("Server is not running")
	} else {
		defer resp.Body.Close()
		displayLaunchMessage(addr)
	}
}

func (app *App) Run(addr string) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(addr, app); err != nil {
			log.Fatal(err)
		}
	}()

	// Attendre que le serveur démarre
	time.Sleep(time.Second)

	// Vérifier si le serveur est en cours d'exécution
	checkServer(addr)

	wg.Wait()
	return nil
}

func displayLaunchMessage(addr string) {
	fmt.Println("*********************************************")
	fmt.Println("***************** Octopus *******************")
	fmt.Println("*********************************************")
	host, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", host)
	fmt.Printf("Listening on address: %s\n", addr)
	fmt.Println("*********************************************")
}
