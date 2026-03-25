package gamedata

import (
	"encoding/json"
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/log"
	"os"
)

var Tables *cfg.Tables

func loader(file string) ([]map[string]interface{}, error) {
	if bytes, err := os.ReadFile("./game-data-gen/excel-json/" + file + ".json"); err != nil {
		return nil, err
	} else {
		jsonData := make([]map[string]interface{}, 0)
		if err = json.Unmarshal(bytes, &jsonData); err != nil {
			return nil, err
		}
		return jsonData, nil
	}
}

func LoadGameData() {
	if tables, err := cfg.NewTables(loader); err != nil {
		panic(err.Error())
	} else {
		Tables = tables
		log.Infof("=======  game data loaded finished")
	}
}
