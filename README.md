API that manages the inventory of a library

# Setup
1. `docker pull redis` to grab redis image
1. `docker run -d --name lib-redis -p 6379:6379 redis` to start container with redis
1. `go install` to compile and place executable in `$GOPATH/bin`
1. `library-inventory` to start API at localhost:1323

# Tests
Regenerate class mock via `mockgen -package=test -destination="test/[file name].go" github.com/hspearman/library-inventory/internal [class name]`
