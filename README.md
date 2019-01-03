# Chunter :: SeeR
C[ourse]hunter: Let Chunter tell you exactly where someone drops that course you've been eyeing.

## Use
The API endpoint is `18.221.69.86:8080`. Use the local examples below and replace IP:port.


## Build and Run
`go chunter.go`

`./chunter`

## Local Server Examples (using HTTPie)
`brew install httpie` || `apt-get install httpie`

Adding to mailing list:

`http GET localhost:8080/add/mail email==test.test@gmail.com server==mail.google.com`

Add a course:

`http GET localhost:8080/add/course subject==CS catalog_number==450`

