package database

type Dsn struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func NewDsn(h string, port string, u string, pass string, n string) *Dsn {
	return &Dsn{Host: h, Port: port, User: u, Pass: pass, Name: n}
}

func (d Dsn) ToString() string {
	return 	d.User + ":" + d.Pass + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.Name + "?charset=utf8&parseTime=True&loc=Local"
}