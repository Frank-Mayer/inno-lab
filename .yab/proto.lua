local yab = require("yab")
local proto_files = yab.find("schema", "**.proto")
require("env")

-- compile proto files
for _, proto_file in ipairs(proto_files) do
	-- compile proto file to go
	local go_file = "server/internal/" .. proto_file:sub(1, -6) .. "pb.go"
	yab.task(proto_file, go_file, function()
		os.execute("protoc --go_out=./server/internal/ --go_opt=paths=source_relative " .. proto_file)
	end)

	-- compile proto file to ts
	local ts_file = "extension/src/" .. proto_file:sub(1, -6) .. "ts"
	yab.task(proto_file, ts_file, function()
		os.execute(
			"protoc --plugin=./extension/node_modules/ts-proto/protoc-gen-ts_proto --ts_proto_out=./extension/src/ "
				.. proto_file
		)
	end)
end
