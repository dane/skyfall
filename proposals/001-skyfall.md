# Skyfall Project

Author: [Dane Harrigan](https://github.com/dane)
Created: December 3, 2020
Status: Draft

An authentication and authorization service that is both language agnostic and extensible. 

## Background

Authentication (authn) and authorization (authz) are problems nearly every product has to address at some point. Standards exist in this space, OAuth 2.0, JSON Web Tokens (JWT), and multi-factor authentication (MFA) for example, but there is little code reuse between products and no reuse across languages. Best practices, such has password hashing algorithms, evolve every few years and migration strategries are necessary to keep user data secure. Building custom authn and authz solutions steals the focus from the actual product or problem at hand.

There are third parties, such as Okta, Auth0, and Google, that will provide authn and authz solutions, but engineering efforts are still allocated to integrate with them.  

## Terminology

- JWT - JSON Web Tokens. Visit https://jwt.io/introduction for a complete description.
- MFA - Multi-factor authentication, most commonly two-factor auth (2FA). When multiple forms of identification are required to authenticate with a service. Popular products are Google Authenticator, Authy, and some services use email or SMS to send additional authentication codes.
-   OAuth 2.0 - A standard for authorizing a service to interact with another service. Vist https://oauth.net/2/ for a complete description.
- Authn - Shorthand for authentication.
- Authz - Shorthand for authorization.

## Proposal



## Abandoned Ideas
