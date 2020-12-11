# Skyfall Project

Author: [Dane Harrigan](https://github.com/dane)
Created: December 3, 2020
Status: Accepted

An authentication and authorization service that is both language agnostic and
extensible. 

## Background

Authentication (authn) and authorization (authz) are problems nearly every product
has to address at some point. Standards exist in this space, OAuth 2.0, JSON Web
Tokens (JWT), and multi-factor authentication (MFA) for example, but there is
little code reuse between products and no reuse across languages. Best practices,
such has password hashing algorithms, evolve every few years and migration
strategries are necessary to keep user data secure. Building custom authn and
authz solutions steals the focus from the actual product or problem at hand.

There are third parties, such as Okta, Auth0, and Google, that will provide
authn and authz solutions, but engineering efforts are still allocated to
integrate with them.  

## Terminology

- JWT - JSON Web Tokens. Visit https://jwt.io/introduction for a complete
  description.
- MFA - Multi-factor authentication, most commonly two-factor auth (2FA). When
  multiple forms of identification are required to authenticate with a service.
  Popular products are Google Authenticator, Authy, and some services use email
  or SMS to send additional authentication codes.
- OAuth 2.0 - A standard for authorizing a service to interact with another
  service. Vist https://oauth.net/2/ for a complete description.
- Authn - Shorthand for authentication.
- Authz - Shorthand for authorization.

## Proposal

Skyfall will consist of an internal API, publc JSON endpoints to fulfill OAuth
2.0 contracts, and HTML endpoints to allow user sign-in and OAuth authorization.
Modularity is a key feature of Skyfall. Endpoints, validations, rendering
engines, etc. will be easily replaceable.

### Internal API

The internal API will be gRPC. It will optionally offer a JSON interface to
support langauges that may prefer JSON over gRPC.

#### User Actions

- CreateUser - Create a new user
- UpdateUser - Update user data
- GetUser - Get a user
- DeleteUser - Soft delete a user
- UndeleteUser - Restore a user from a deleted state
- SuspendUser - Suspend a user
- UnsuspendUser - Restore a user from a suspended state

#### OAuth 2.0 Client Actions

- CreateOAuthClient - Create a new OAuth 2.0 client
- GetOAuthClient - Get an OAuth 2.0 client
- DeleteOAuthClient - Delete an OAuth 2.0 client
- UndeleteOAuthClient -Restore an OAuth 2.0 client from a deleted state
- SuspendOAuthClient - Suspend an OAuth 2.0 client
- UnsuspendOAuthClient - Restore an OAuth 2.0 client from a suspended state

#### Authn/Authz Actions

- Authenticate - Authenticate a user
- CreateAuthorizationGrant - Create an authorization grant for an OAuth 2.0
  client
- CreateAccessToken - Create an access token for an OAuth 2.0 client
- RefreshAccessToken - Refresh an access token for an OAuth 2.0 client
- VerifyAccessToken - Verify an OAuth 2.0 access token

### JSON Endpoints

To support the OAuth 2.0 contract the following endpoint(s) must be present:

- POST /oauth/token

### HTML Endpoints

- GET /sign-in - Allow a user to authenticate
- GET /sign-out - Allow a user to end their authenticated session
- GET /oauth/authorize - Allow a user to authorize another service to access
  their data

## Out of Scope

- MFA - While MFA was explained earlier, it is out of scope for this proposal.
  Once the internal API is feature complete as described above, I will start to
  work on MFA.
- User creation HTML endpoints - The act of creating new users, essentially a
  user sign-up page, should be handled by the operators of Skyfall. The user
  model is designed to hold arbitrary data so it is left to the operator on what
  should be provided during creation. 
