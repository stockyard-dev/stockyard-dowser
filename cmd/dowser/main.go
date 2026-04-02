package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-dowser/internal/server";"github.com/stockyard-dev/stockyard-dowser/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./dowser-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("dowser: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Dowser — Self-hosted data quality monitor\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("dowser: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
