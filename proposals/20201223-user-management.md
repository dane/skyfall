# User Management

Author: [Dane Harrigan](https://github.com/dane)
Created: December 23, 2020
Status: Draft

This proposal will define the internal APIs for user fetching, creating,
updates, deletes, verifications, and suspensions. 

## Background

In most services, a user account must exist first, followed by a verification
step for added security. User account updates, deletions, and suspensions can
occur at any time.

While not every service will expose user verifications or suspensions, Skyfall
will offer the features.

## Terminology

_TBD_

## Proposal

There are standard fields for a user account such as username, hashed password,
and timestamps to persist creation time, and last updated. These will be present
on the `User` protobuf message. A `properties` field will also be present to
capture all additional data.

The following will describe all APIs and incoming and outgoing protobuf
messages.

### CreateUser

The minimum required fields to create a user are a username, password, and a
password confirmation. All other data will be stored in properties. The
username can be a minimum of alpha-numeric characters or an email address,
depending on the configuration of the service. Passwords must be a minimum of 8
characters, and two of the following:
- Uppercase and lowercase characters
- At least one number
- At least one non-alpha-numeric character

```
message CreateUserRequest {
  string username = 1;
  string password = 2;
  string password_confirmation = 3;
  google.protobuf.Struct properties = 4;
}

message CreateUserResponse {
  User user = 1;
}

message User {
  string id = 1;
  string username = 2;
  string password = 3;
  string password_confirmation = 4;
  google.protobuf.Struct properties = 5;
  google.protobuf.Timestamp created_at = 6;
  google.protobuf.Timestamp updated_at = 7;
  google.protobuf.Timestamp verified_at = 8;
  google.protobuf.Timestamp suspended_at = 9;
}
```

### UpdateUser

The update operation will accept a field mask to allow single field
modifications. If the field mask is omitted, all fields will be assumed.
Timestamps (eg: `created_at` and `verified_at`) cannot be altered via the
update operation.

```
message UpdateUserRequest {
  User user = 1;
  google.protobuf.FieldMask fields = 2;
}

message UpdateUserResponse {
  User user = 1;
}
```

### DeleteUser

The delete operation is a "soft delete" to later support undeleting a user. The
`DeleteUserResponse` is empty, but intentionally not `google.protobuf.Empty` in
case data needs to be returned in the future.

```
message DeleteUserRequest {
  string id = 1;
}

message DeleteUserResponse {}
```

### VerifyUser

The verify operation is meant to support scanerios where a service requires a user
to verify their account before being able to user said service. All operations
described in this proposal will be allowed for unverified accounts.

```
message VerifyUserRequest {
  string id = 1;
}

message VerifyUserResponse {
  User user = 1;
}
```

### SuspendUser

Suspending an account renders the account unusable. This differs from the delete
operation because this does not free the account name.

```
message SuspendUserRequest {
  string id = 1;
}

message SuspendUserResponse {
  User user = 1;
}
```

### GetUser and GetUserByUsername

To align with many operations that target an account by ID, the `GetUser`
operation accepts an ID string. In some cases only a 

## Out of Scope

State any topics that have been considered, but were scoped out. Explain why
the topics were out of scope.
