/*
 * dircat - a small utility useful for exposing files to http requests.
 * Copyright (C) 2020  Francesco Giardino
 *
 */

package main

import "github.com/BurntSushi/toml"

type Dir struct {
	Path string `toml:"path"`
}

type Config struct {
	Port int            `toml:"port"`
	Dirs map[string]Dir `toml:"dir"`
}

func readConfig(path string) (Config, error) {
	var c Config

	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
