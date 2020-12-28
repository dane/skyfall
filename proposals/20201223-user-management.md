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
  string username = 1;
  string password = 2;
  string password_confirmation = 3;
  google.protobuf.Struct properties = 4;
  google.protobuf.Timestamp created_at = 5;
  google.protobuf.Timestamp updated_at = 6;
  google.protobuf.Timestamp verified_at = 7;
  google.protobuf.Timestamp suspended_at = 8;
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

### Delete User

### Verify User

### Suspend User



## Out of Scope

State any topics that have been considered, but were scoped out. Explain why
the topics were out of scope.
