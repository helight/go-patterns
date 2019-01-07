export GOPATH=$(shell pwd)

test:
	# creational
	go test -v creational/01_factory/ 
	go test -v creational/02_abstractfactory/
	go test -v creational/03_singleton/
	go test -v creational/04_builder/
	go test -v creational/05_prototype/
	
	# structral
	go test -v structural/06_adapter/
	go test -v structural/07_bridge/
	go test -v structural/08_filter/
	go test -v structural/09_composite/
	go test -v structural/10_decorator/
	go test -v structural/11_facade/
	go test -v structural/12_flyweight/
	go test -v structural/13_proxy/
	
	# behavioral
	go test -v behavioral/14_chain/
	go test -v behavioral/15_command/
	go test -v behavioral/16_interpreter/
	go test -v behavioral/17_mediator/
	go test -v behavioral/18_memento/
	go test -v behavioral/19_observer/
	go test -v behavioral/20_state
	go test -v behavioral/21_strategy


