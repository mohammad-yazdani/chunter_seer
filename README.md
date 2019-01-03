# Chunter :: SeeR
C[ourse]hunter: Let Chunter tell you exactly where someone drops that course you've been eyeing.

## Build and Run
`go chunter.go`

`./chunter`

## Local Test (using HTTPie)
`brew install httpie` || `apt-get install httpie`

Adding to mailing list:

`http GET localhost:8080/add/mail email==test.test@gmail.com server==mail.google.com`

Add a course:

`http GET localhost:8080/add/course subject==CS catalog_number==450`

