package main

type configArgs struct {
	hostinf hostInf
	dsn     dsn
}
type hostInf struct {
	HostIP   string `json:"host_ip"`
	HostPort string `json:"host_port"`
}
type dsn struct {
	User     string `json:"db_user"`
	Pwd      string `json:"db_pwd"`
	Hostname string `json:"db_hostname"`
	Port     string `json:"db_port"`
	Schema   string `json:"db_schema"`
}

func main() {

}
