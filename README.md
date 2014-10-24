expiring-memory-store [![Circle CI](https://circleci.com/gh/skidder/expiring-memory-store/tree/master.png?style=badge)](https://circleci.com/gh/skidder/expiring-memory-store/tree/master)
=====================

Go implementation of an expiring key-value map

Installation
------------

Install expiring-memory-store using the "go get" command:

    go get github.com/skidder/expiring-memory-store/ems

The Go distribution and [concurrent-map](https://github.com/streamrail/concurrent-map) (a thread-safe map for Go) are the only dependencies.


Tests
-----
Run the test suite from the ```ems``` directory:

```shell
$ git clone https://github.com/skidder/expiring-memory-store.git
$ cd expiring-memory-store/ems/
$ go test
PASS
ok  	_/Users/skidder/git/expiring-memory-store/ems	2.010s
```

Samples
-------

### Creating a new store

```go
package main

import "github.com/skidder/expiring-memory-store/ems"

func main() {
	m := NewExpiringMemoryStore()
}
```


### Writing to the store with the default expiration (20 seconds)

```go
package main

import "github.com/skidder/expiring-memory-store/ems"

func main() {
	m := NewExpiringMemoryStore()
	m.Write("key", "value")
}
```

### Writing to the store with a specific expiration

```go
package main

import "github.com/skidder/expiring-memory-store/ems"

func main() {
	m := NewExpiringMemoryStore()

	// expire in 60 seconds
	m.WriteWithExpiration("key", "value", 60)
}
```

### Read an item from the store

```go
package main

import "time"
import "github.com/skidder/expiring-memory-store/ems"

func main() {
	m := NewExpiringMemoryStore()

	m.Write("key", "value")
	
	// val will have "value", and err will be nil
	val, err := m.Read("key")
	
 	// val will have "", and err equals ems.elementNotFoundError
	val, err := m.Read("missing")
  
	time.Sleep(time.Duration(30)*time.Second)

	// val will have "", and err equals ems.expiredElementError
	val, err := m.Read("key")
}
```

License
-------

expiring-memory-store is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
