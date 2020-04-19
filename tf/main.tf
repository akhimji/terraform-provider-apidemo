provider "apidemo" {
  address = "http://localhost/"
  port    = "12345"
  key     = "/people/"
}


resource "apidemo_entry" "entry" {
  apiid        = "3"
  firstname    = "Charles"
  lastname     = "Dickens"
}
