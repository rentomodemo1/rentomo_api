package config

type Settings struct{ Address string }

func Load() Settings { return Settings{Address: ":8081"} }
