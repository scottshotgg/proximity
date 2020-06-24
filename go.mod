module github.com/scottshotgg/proximity

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/scottshotgg/proximity/pkg/buffs v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0
)

replace github.com/scottshotgg/proximity/pkg/buffs => ./pkg/buffs
