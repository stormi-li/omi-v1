package omiclient

import "time"

const namespace_separator = ":"
const config_expire_time = 2 * time.Second

const Prefix_Config = "stormi:config:"
const Prefix_Server = "stormi:server:"
const Prefix_Web = "stormi:web:"

var Server = "Server"
var Config = "Config"
var Web = "Web"
