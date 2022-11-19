package main

import (
	"os"

	golog "github.com/withmandala/go-log"
)

const INPUT_DATA int = 3
const ORDINAL_FIRST_NUMBER int = 0
const ORDINAL_SECOND_NUMBER int = 2
const OPERATION_SIGN int = 1

var OPERATIONS [4]string = [4]string{"+", "-", "*", "/"}

var logger = golog.New(os.Stderr).WithColor()
