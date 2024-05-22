package database

type DBModel struct {
	ServerMode   string `config:"server_mode"`
	Driver       string `config:"db_driver"`
	Host         string `config:"db_host"`
	Port         string `config:"db_port"`
	Name         string `config:"db_name"`
	Username     string `config:"db_username"`
	Password     string `config:"db_password"`
	MaxIdleConn  int    `config:"conn_idle"`
	MaxOpenConn  int    `config:"conn_max"`
	ConnLifeTime int    `config:"conn_lifetime"`
}

var (
	POSGRES_CONFIG = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	MYSQL_CONFIG   = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)
