# jurassic-park-api
This is an API for managing the dinosaurs, cages, and species in a Jurassic Park.

I was not able to get to the testing aspect, which would have helped make sure that the API was more stable.

For concurrency, some of that is already baked in by using Go and Gin with Goroutines. The main issue becomes accessing the database.

The database would need a mechanism like locking to prevent inconsistentcies. 

Currently running on railway at jurassic-park-api-production.up.railway.app