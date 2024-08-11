package main
import(

	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/dont9s/mysql-go-server/pkg/routes"
)

func main(){
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", router))
	
}