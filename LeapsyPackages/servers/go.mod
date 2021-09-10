module servers

go 1.16

replace leapsy.com/packages/configurations => ../configurations

replace leapsy.com/packages/logings => ../logings

replace leapsy.com/packages/network => ../network

replace leapsy.com/records => ../records

replace leapsy.com/times => ../times

replace leapsy.com/databases => ../databases

replace leapsy.com/packages/emails => ../emails

replace leapsy.com/packages/gpses => ../gpses

replace leapsy.com/packages/jwts => ../jwts

replace leapsy.com/packages/jsons => ../jsons

replace leapsy.com/packages/paths => ../paths

replace leapsy.com/packages/tlss => ../tlss

require (
	github.com/gin-gonic/autotls v0.0.3
	github.com/gin-gonic/gin v1.7.3
	github.com/shogo82148/androidbinary v1.0.2
	github.com/sirupsen/logrus v1.8.1
	leapsy.com/databases v0.0.0-00010101000000-000000000000
	leapsy.com/packages/configurations v0.0.0-00010101000000-000000000000
	leapsy.com/packages/logings v0.0.0-00010101000000-000000000000
	leapsy.com/packages/network v0.0.0-00010101000000-000000000000
	leapsy.com/packages/paths v0.0.0-00010101000000-000000000000
	leapsy.com/records v0.0.0-00010101000000-000000000000
)
