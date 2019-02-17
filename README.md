API that manages the inventory of a library

# Setup
1. `docker pull redis` to grab redis image
1. `docker create --name lib-redis -p 6379:6379 redis` to create redis container with exposed port 6379

# Run
1. `docker start lib-redis` to launch container
1. `go install` to compile
1. `library-inventory` to start API at localhost:1323

# Tests
- `mockgen -package=test -destination="test/[file name].go" github.com/hspearman/library-inventory/internal [class name]` to regenerate class mock
- `go test ./test -v` to run tests
