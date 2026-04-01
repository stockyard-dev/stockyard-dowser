package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-dowser/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.Expectation{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var e store.Expectation;json.NewDecoder(r.Body).Decode(&e);if e.Name==""||e.Dataset==""||e.Rule==""{writeError(w,400,"name, dataset, rule required");return};s.db.Create(&e);writeJSON(w,201,e)}
func(s *Server)handleCheck(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var c store.CheckResult;json.NewDecoder(r.Body).Decode(&c);c.ExpectationID=id;s.db.RecordCheck(&c);writeJSON(w,200,c)}
func(s *Server)handleResults(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListResults(id);if list==nil{list=[]store.CheckResult{}};writeJSON(w,200,list)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
