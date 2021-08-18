module main

go 1.14

replace leapsy.com/packages/logings => ../logings

replace leapsy.com/packages/configurations => ../configurations

require (
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/ini.v1 v1.62.0 // indirect
	leapsy.com/packages/configurations v0.0.0-00010101000000-000000000000
	leapsy.com/packages/logings v0.0.0-00010101000000-000000000000
)
