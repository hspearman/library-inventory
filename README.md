API that manages the inventory of a library

# Run
1. `docker-compose up --build` to launch containers
1. `GET localhost:1323` to test endpoint

# Tests
- `mockgen -package=test -destination="test/[file name].go" github.com/hspearman/library-inventory/internal [class name]` to regenerate class mock
- `go test ./test -v` to run tests
