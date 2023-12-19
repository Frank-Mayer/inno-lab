local proto_files = Yab.find("schema", "**.proto")

-- compile proto files
for _, proto_file in ipairs(proto_files) do
	local proto_info = Yab.fileinfo(proto_file)

    -- compile proto file to go
	local go_file = "server/internal/" .. proto_file:sub(1, -6) .. "pb.go"
	local go_info_success, go_info = pcall(Yab.fileinfo, go_file)
	if not go_info_success or go_info.modtime < proto_info.modtime then
		print("compile proto file: " .. proto_file .. " to go")
		os.execute("protoc --go_out=./server/internal/ --go_opt=paths=source_relative " .. proto_file)
	end

    -- compile proto file to ts
	local ts_file = "extension/src/" .. proto_file:sub(1, -6) .. "ts"
	local ts_info_success, ts_info = pcall(Yab.fileinfo, ts_file)
	if not ts_info_success or ts_info.modtime < proto_info.modtime then
		print("compile proto file: " .. proto_file .. " to ts")
		os.execute(
			"protoc --plugin=./extension/node_modules/ts-proto/protoc-gen-ts_proto --ts_proto_out=./extension/src/ "
				.. proto_file
		)
	end
end
