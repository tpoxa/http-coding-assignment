# Coding Assignment
The goal of this assignment is to build a simple HTTP service (with the appropriate unit tests) with two endpoints:

### POST /auth

Accepts JSON input in the format:

`{"username": "<user name>", "password": "<user password>"}`

and returns JWT OAUTH 2/OIDC token with the username as a subject. The username and the password don't have to be verified, but should not accept empty strings. The JWT token should expire in one hour.

It should return appropriate error status code if the JSON payload is not valid, or the username and password are not valid (are empty)


### POST /sum

Protected with a valid JWT token, generated by the **/auth** endpoint, provided as a Bearer Authorization header.

Accepts arbitrary JSON document as payload, which can contain a variety of things: **arrays** `[1,2,3,4]`, arbitrary **objects** `{"a":1, "b":2, "c":3}`, **numbers**, and **strings**. The endpoint should find all of the numbers throughout the document and add them together.

For example:

- **[1,2,3,4]** and **{"a":6,"b":4}** both have a sum of **10**.
- **[[[2]]]** and **{"a":{"b":4},"c":-2}** both have a sum of **2**.
- **{"a":[-1,1,"dark"]}** and **[-1,{"a":1, "b":"light"}]** both have a sum of **0**.
- **[]** and **{}** both have a sum of **0**.

This is not an **exhaustive** list of examples. Think about what other edge cases there might be.

The response should be the **hex digest of the SHA256 hash of the sum of all numbers in the document**. It should return the appropriate error status code if the JWT token or the JSON payload are not valid.

## Technical details

## Requirements
The only strict requirement we pose is that the app is written in **Go** using Go modules. We recommend the use of Go patterns wherever meaningful of course, and be prepared to show and explain them when presenting the solution. 

### Project Setup
Please setup the application to be build in a standard Go way or else provide a `start.sh` script.

## Bonus (optional)
A Dockerfile to build and run the service.

### External Libraries
We don't mind if you use whatever external libraries you like as long as they are with a non-restrictive for commercial use license.

### Estimation
You will see an open issue "Call for estimation". Please estimate by writing a comment when you think the task will be ready before you start. We don't set any hard deadlines.
