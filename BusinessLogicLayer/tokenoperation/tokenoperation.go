package tokenoperation

import (
	"go/DataAccessLayer/dbfunction"
	"go/pkg/structpkg/data"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func Buildtoken(account string, email string) data.Usertoken {
	if account == "" {
		account = dbfunction.SelectUserAccount(email)
	}

	var t1 data.Usertoken
	t1.Account = account
	t1.Tokenstring = uuid.NewV4().String()
	return t1
}
func Buildadmintoken(email string) data.Usertoken {

	account := dbfunction.SelectUserAccount(email)

	var t1 data.Usertoken
	t1.Account = account
	t1.Tokenstring = uuid.NewV4().String()
	return t1
}

func GetUsertoken(r *http.Request) (string, string) {

	r.ParseForm()

	t := r.URL.Query()

	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	return keys[0], r.FormValue(keys[0])
}
