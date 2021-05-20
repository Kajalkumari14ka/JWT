# Implement JWT based token Authentication in Golang:

JWT: JSON Web Token
It contains API which is authenticated by JWT.

3 Endpoint:

-**/login**: To login. Once user logged in, JWT has been created for it. After it for every request JWT is passed from 
client, so server can authenticate the request.

-**/home**: it will authenticate and validate the token which is stored in cookie when he has logged in 
and if cookie does not have any token which means he has not logged in previously, it will give unauthorized exception.

-**/refresh**: If we have created token which is  available for 5 min suppose. After 5 min token has to be refreshed and
reused again. This API is to create refresh token and use in API.

