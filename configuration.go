package main

type Configuration struct {
	Host struct {
		IP   string
		Port string
	}
	Loaded bool
}

const Config string = "config.json"

// var configuration := Configuration{}

// func Configuration() {
// 	if Configuration.loaded {
// 		return Configuration
// 	}
//
// 	file, _ := os.Open(Config)
// 	decoder := json.NewDecoder(file)
//
// 	err := decoder.Decode(&Configuration)
// 	if err != nil {
// 		log.Fatal("Error loading config", err)
// 	} else {
// 		Configuration.loaded = true
// 		return Configuration
// 	}
// }
