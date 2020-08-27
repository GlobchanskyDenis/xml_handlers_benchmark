go test -bench . benchmark_test.go main.go standart.go stream.go
go test -bench . -benchmem benchmark_test.go main.go standart.go stream.go
go test -bench . -benchmem -cpuprofile=cpu.out benchmark_test.go main.go standart.go stream.go libxml2.go
go test -bench . -benchmem -memprofile=mem.out benchmark_test.go main.go standart.go stream.go libxml2.go
go tool pprof main.test mem.out
	web
	alloc_objects
	list <Часть названия или полное название функции>