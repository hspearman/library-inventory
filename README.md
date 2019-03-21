API that manages the inventory of a library

# Run
1. `docker-compose up` to launch containers
1. `go install` to compile
1. `library-inventory` to start API at localhost:1323

# Tests
- `mockgen -package=test -destination="test/[file name].go" github.com/hspearman/library-inventory/internal [class name]` to regenerate class mock
- `go test ./test -v` to run tests
