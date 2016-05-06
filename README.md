This is my attempt at learning the Go programming language. To compile the Go scripts in this repository make sure your GOPATH environment variable is set to /path/to/learngo. If you want to add custom modules, place them in a separate directory under the src folder and make sure the first line in the Go script is package <package_name>. If GOPATH is set correctly you should be able to import "package-name" and use the module. Typically the package name and the folder under the src directory would be the same name. If they are different, then import "<src_dir_name>" and use the imported functions by calling <package_name>.<function>.

* `factor.go` - Will factor any unsigned 64-bit integer.
* `isprime.go` - Will tell you if an unsigned 64-bit integer is prime.
* `sieve.go` - Will tell you how many prime numbers are below a given unsigned 64-bit integer. Can easily be modified to write out the primes. I don't know how large a sieve it can do but I've tested up to 4,300,000,000 which takes about 1.5 minutes on my 4 year old MBP.

