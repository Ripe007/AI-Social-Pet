package config

type Config struct {
	Mysql `ini:"mysql"`
}

type Mysql struct {
	AliasName  string `ini:"alias_name"`
	DriverName string `ini:"driver_name"`
	User       string `ini:"user"`
	Password   string `ini:"password"`
	Ip         string `ini:"ip"`
	Db         string `ini:"db"`
	Port       string `ini:"port"`
}
