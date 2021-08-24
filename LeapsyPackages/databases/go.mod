module databases

go 1.16

replace leapsy.com/packages/configurations => ../configurations

replace leapsy.com/packages/logings => ../logings

replace leapsy.com/packages/network => ../network

replace leapsy.com/records => ../records

require (
	github.com/sirupsen/logrus v1.8.1
	go.mongodb.org/mongo-driver v1.6.0
	leapsy.com/packages/configurations v0.0.0-00010101000000-000000000000
	leapsy.com/packages/logings v0.0.0-00010101000000-000000000000
	leapsy.com/packages/network v0.0.0-00010101000000-000000000000
	leapsy.com/records v0.0.0-00010101000000-000000000000
)
