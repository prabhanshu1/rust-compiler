all:
	@ echo "Install GOlang simply using sudo apt-get install golang in debian systems"
	@ echo "Required packages - nex at https://github.com/blynn/nex and GOlang"
	nex -s src/rust.nex
	go build -o bin/lexer src/rust.nn.go
clean:
	rm bin/lexer
	rm src/rust.nn.go

