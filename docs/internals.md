# Internals

## Routes

| Path | Method | Visibility | Description |
|--|--|--|--|
| /sign-in | GET | Public | Serves a sign-in form to users
| /sign-in | POST | Public | Processes sign-in requests
| /sign-out | GET | Private | Terminates a user session

### User Session

A user session is a JWT. It has scopes to create and revoke OAuth tokens and a
default expiry of 30 days. The JWT is stored as a secure cookie named
`_session`.

### Claims

| Name | Type | Description |
|--|--|--|
| sub | String | User ID once signed in |
| scope | Array | Values `token:read` and `token:write` once signed in |
| err | String | Error messaging |
| aud | String | A unique identifier of the service |
| iat | Number | Issued timestamp in epoch format |
| nbf | Number | Identical value of `iat` |
| exp | Number | Expiration timestamp in epoch format |
