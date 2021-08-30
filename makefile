program_target=reverse-proxy
mainfile=main.go

# build and run
build:
	env go build -o ${program_target} ${mainfile}

run: build
	./${program_target}

# automatic install 
install_linux: build
	sudo mv ${program_target} /bin/

install_termux: build
	mv ${program_target} /usr/bin/

