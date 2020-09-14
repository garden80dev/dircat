// +build !windows

/*
 * Logcat - a small utility useful for exposing logs to http requests.
 * Copyright (C) 2020  Francesco Giardino
 *
 */

package main

import "os"

var HOME = os.Getenv("HOME")
