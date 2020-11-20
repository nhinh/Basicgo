# Unit test for Basic go
1. Will test the functions
<br />in controller (CreateUser, CreateUserSelect, GetUser, GetUsers, UpdateUser, DeleteUser)
<br />in models (create, get, update, delete)
2. Test the functions in the controller (TestCreateUser, TestDeleteUser, TestGetUserByID, TestGetUsers, TestUpdateUser)
## Test procedure of a function: 
__Example__: TestCreateUser ()
1. Delete table
2. Create sample data, Including true and false records (same email, same username, incorrect email, username empty, email empty, pass empty)
3. Create a request
4. Call the method to be tested
5. Parse result
6. Compare the actual received status and the desired status.
7. Compare the actual received response and the desired response.
Similar to other test functions are similar

## How to run the test
### Run each function: 
> cd tests/controller<br />
> go test -v -run function_name
### Run all: 
> go test ./...