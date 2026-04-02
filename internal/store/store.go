package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Check struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Rule string `json:"rule"`
	Column string `json:"column_name"`
	Severity string `json:"severity"`
	LastResult string `json:"last_result"`
	FailCount int `json:"fail_count"`
	LastRunAt string `json:"last_run_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"dowser.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS checks(id TEXT PRIMARY KEY,name TEXT NOT NULL,source TEXT DEFAULT '',rule TEXT DEFAULT '',column_name TEXT DEFAULT '',severity TEXT DEFAULT 'warning',last_result TEXT DEFAULT 'pending',fail_count INTEGER DEFAULT 0,last_run_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Check)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO checks(id,name,source,rule,column_name,severity,last_result,fail_count,last_run_at,created_at)VALUES(?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Source,e.Rule,e.Column,e.Severity,e.LastResult,e.FailCount,e.LastRunAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Check{var e Check;if d.db.QueryRow(`SELECT id,name,source,rule,column_name,severity,last_result,fail_count,last_run_at,created_at FROM checks WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Source,&e.Rule,&e.Column,&e.Severity,&e.LastResult,&e.FailCount,&e.LastRunAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Check{rows,_:=d.db.Query(`SELECT id,name,source,rule,column_name,severity,last_result,fail_count,last_run_at,created_at FROM checks ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Check;for rows.Next(){var e Check;rows.Scan(&e.ID,&e.Name,&e.Source,&e.Rule,&e.Column,&e.Severity,&e.LastResult,&e.FailCount,&e.LastRunAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM checks WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM checks`).Scan(&n);return n}
