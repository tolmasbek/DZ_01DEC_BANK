package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const settingsFileName = "users.json"

func FileRecordingJson(user User) {
	bytes, err2 := json.MarshalIndent(user,""," ")
	if err2 != nil {
		log.Fatal("Err Marshall",err2)
	}

	err := ioutil.WriteFile(settingsFileName, bytes, 0666)
	if err != nil {
		log.Fatal("Cannot write file:", err)
	}
}
