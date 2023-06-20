package reform

//go:generate reform

// reform_models represents a row in models table.
//
//reform:models
type ReformModels struct {
	ID      int    `reform:"id,pk"`
	Name    string `reform:"name"`
	Title   string `reform:"title"`
	Fax     string `reform:"fax"`
	Web     string `reform:"web"`
	Age     int32  `reform:"age"`
	Right   bool   `reform:"right"`
	Counter int64  `reform:"counter"`
}
