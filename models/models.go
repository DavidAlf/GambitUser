package models

type SecretRDSJson struct {
	UserName            string `json:"username"` //ALT Izq + 96 para solo json en minuscula
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Puerto              int    `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
}

type UserSingUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
