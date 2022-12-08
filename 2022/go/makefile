# create a new library and append to workspace

mklib:
	bash ./scripts/mk-lib.sh $(day)

rm: 
	bash ./scripts/rm-lib.sh $(day)

main:
	go run .

test: 
	go test ./day-0$(day)/ .