require("env")
require("proto")
local yab = require("yab")

local bin_name = "server_bin"
if yab.os_type() == "windows" then
	bin_name = bin_name .. ".exe"
end

yab.task(yab.find("./server/", "**.go"), "server/"..bin_name, function()
	yab.cd("./server/", function()
		os.execute('go build -ldflags="-s -w" -o ' .. bin_name .. " ./cmd/server/")
	end)
end)

return "server/" .. bin_name
