# go-setup-module
When I learn to separate our code into module 

- For production use, you’d publish the example.com/user module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/user code on your local file system

- To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is)

1. From the command prompt in the hello directory, run the following command:
```
go mod edit -replace veasna.com/user=../user
```
