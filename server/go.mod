module example.com/server

go 1.22.5

replace example.com/recipes => ./pkg/recipes

require example.com/recipes v0.0.0-00010101000000-000000000000

require (
	github.com/gosimple/slug v1.14.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
)
