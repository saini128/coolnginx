package nginx

import (
	"coolnginx/db"
	"encoding/json"
	"fmt"

	crossplane "github.com/nginxinc/nginx-go-crossplane"
)

func LoadNginxMainConfigFile() map[string]interface{} {
	defPath := "/etc/nginx/nginx.conf"

	payload, err := crossplane.Parse(defPath, &crossplane.ParseOptions{})
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	var nginxConf map[string]interface{}
	if err := json.Unmarshal(b, &nginxConf); err != nil {
		panic(err)
	}

	return nginxConf
}
func StoreNginxMainConfigFile() {
	conf := LoadNginxMainConfigFile()
	err := db.StoreNginxConfig(conf)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conf Saved Successfully")

	}
}
