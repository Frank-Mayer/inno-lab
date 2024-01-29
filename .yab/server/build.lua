require("env")
local yab = require("yab")

local bin_name = "server_bin"
if yab.os_type() == "windows" then
	bin_name = bin_name .. ".exe"
end

os.execute("mkdir -p out")

yab.task(yab.find("./server/", "**.go"), "out/"..bin_name, function()
	yab.cd("./server/", function()
		os.execute('go build -ldflags="-s -w" -o ../out/' .. bin_name .. " ./cmd/server/")
	end)
end)

return "out/" .. bin_name
