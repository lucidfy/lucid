package config

import "time"

const HOST = "0.0.0.0"
const PORT = "8080"
const WRITE_TIMEOUT = time.Second * 15
const READ_TIMEOUT = time.Second * 15
const IDLE_TIMEOUT = time.Second * 60
