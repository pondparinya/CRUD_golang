test:
	go test -v CRUD

test-coverage:
	go test -covermode=count -coverprofile cover.out -v CRUD/...



