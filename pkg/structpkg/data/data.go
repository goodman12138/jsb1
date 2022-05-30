package data

type Userregister struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	State     int
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Time      string `json:"time"`
	Emailcode string `json:"emailcode"`
	Headphoto string `json:"headphoto"`
}

type Logoutuser struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	Reason  string `json:"Reason"`
	Emailcode string `json:"emailcode"`
}

type Modifyuseremail struct {
	Oldemail string `json:"oldemail"`
	Newemail string `json:"newemail"`
	Reason   string `json:"reason"`
	Emailcode string `json:"emailcode"`
}

type Dairy struct {
	Dairyid     int    `json:"dairyid"`
	Account     string `json:"account"`
	Time        string `json:"time"`
	Details     string `json:"details"`
	Title       string `json:"title"`
	Classify    string `json:"classify"`
	Performance string `json:"performance"`
}

type Admin struct {
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Emailcode string `json:"emailcode"`
}

type Usertoken struct {
	Account     string `json:"account"`
	Tokenstring string `json:"tokenstring"`
}
