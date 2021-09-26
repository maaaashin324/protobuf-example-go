protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/simple/simple.proto
protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/enum_example/enum_example.proto
