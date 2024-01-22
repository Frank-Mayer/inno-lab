require("env")
require("proto")
local yab = require("yab")

local bin_name = require("server.build")

if yab.os_type() == "windows" then
	os.execute("start " .. bin_name)
else
	os.execute("./" .. bin_name)
end
