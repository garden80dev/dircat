// +build windows

/*
 * dircat - a small utility useful for exposing directories to http requests.
 * Copyright (C) 2020  Francesco Giardino
 *
 */

package main

import "os"

var HOME = os.Getenv("UserProfile")
