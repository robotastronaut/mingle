package muddler

type Script struct {
	Item
	EventHandlers []string `json:"eventHandlerList"`
}

type Scripts []Script

func LoadScripts(path string) ([]Script, error) {
	scripts, err := loadItemManifestAt(path)
	if err != nil {
		return nil, err
	}

	// for _, script := range scripts {

	// 	err = script.LoadScript()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return scripts, nil
}
