# terraform-provider-apidemo
Building a simple custom provider to gain a better understanding of the CRUD Terraform spec
Provider is built against a Go based Restful API server (https://github.com/alyarctiq/go-restapi-server)

Build
```
git clone
go mod tidy
go get -u -v

go bulid 
go build -o terraform-provider-apidemo
mv tf
```

Testing
```
This provider is built against a simple Go based Restful API server 
(https://github.com/alyarctiq/go-restapi-server)
Build and Run the above web-app. Listed on localhost:12345

$ curl http://localhost:12345/people/1
{"id":"1","firstname":"Ernest","lastname":"Hemingway","address":{"city":"Dublin","state":"CA"}}

$ curl http://localhost:12345/people/2
{"id":"2","firstname":"George","lastname":"Orwell"}

```

Validate there is no entry at position 3
```
$ curl http://localhost:12345/people/3
{}


Setup TF vars

cat main.tf 
resource "apidemo" "entry" {
  apiid        = "3"
  firstname    = "Mark"
  lastname     = "Twain"
}

```

Init
```
$ terraform init

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

Plan
```
$ terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.


------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + apidemo.entry
      id:        <computed>
      apiid:     "3"
      firstname: "Mark"
      lastname:  "Twain"


Plan: 1 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.

```

Apply
```
 terraform apply -auto-approve 
apidemo.entry: Creating...
  apiid:     "" => "3"
  firstname: "" => "Mark"
  lastname:  "" => "Twain"
apidemo.entry: Creation complete after 0s (ID: 3)

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

Validate
```
$curl http://localhost:12345/people/3

{"id":"3","firstname":"Mark","lastname":"Twain"}

Entry is Now There
```

Update
```
## Update TF Vars

$ cat main.tf 
resource "apidemo" "entry" {
  apiid        = "3"
  firstname    = "Charles"
  lastname     = "Dickens"
}

### Plan ###

$terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.

apidemo.entry: Refreshing state... (ID: 3)

------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  ~ apidemo.entry
      firstname: "Mark" => "Charles"
      lastname:  "Twain" => "Dickens"


Plan: 0 to add, 1 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.


### Apply ###
$ terraform apply -auto-approve
apidemo.entry: Refreshing state... (ID: 3)
apidemo.entry: Modifying... (ID: 3)
  firstname: "Mark" => "Charles"
  lastname:  "Twain" => "Dickens"
apidemo.entry: Modifications complete after 0s (ID: 3)

Apply complete! Resources: 0 added, 1 changed, 0 destroyed.

### Validate ###
$ curl http://localhost:12345/people/3
{"id":"3","firstname":"Charles","lastname":"Dickens"}

Update Successful!
```

Destroy
```
$ terraform destroy -auto-approve
apidemo.entry: Refreshing state... (ID: 3)
apidemo.entry: Destroying... (ID: 3)
apidemo.entry: Destruction complete after 0s

Destroy complete! Resources: 1 destroyed.

### Validate ###
$ curl http://localhost:12345/people/3
{}
```