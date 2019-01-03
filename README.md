# Chunter :: SeeR
C[ourse]hunter: Let Chunter tell you exactly where someone drops that course you've been eyeing.

## Use
The API endpoint is `18.221.69.86:8080`. The `examples` folder has some example on how to interact with Chunter_SeeR JSON API.

All HTTP calls are POST (with request JSON in the body).


## Build and Run
`go chunter.go`

`./chunter`

## Local Server Examples (using HTTPie)
Install HTTPie: 

- macOS: `brew install httpie`

- Ubuntu/WSL: `apt-get install httpie`

Adding to mailing list:

` http -v localhost:8080 < examples/add_mail.json`

Adding course(s):

` http -v localhost:8080 < examples/add_course.json`

Getting stats:

`http -v localhost:8080 < examples/get_stats.json`
