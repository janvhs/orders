package config

func Default() Config {
	return Config{
		IsDevelopment: false,
		DB: DB{
			DSN: "orders.db",
		},
		Server: Server{
			Host: "127.0.0.1",
			Port: 8080,
		},
		LDAP: LDAP{
			Host:         "",
			Port:         389,
			BindDN:       "",
			BaseDN:       "",
			BindPassword: "",
		},
	}
}
