# Account Management

Author: [Dane Harrigan](https://github.com/dane)
Created: December 23, 2020
Status: Draft

This proposal will define the internal APIs for account fetching, creating,
updates, deletes, verifications, and suspensions. 

## Background

In most services, a account must exist first, followed by a verification
step for added security. Account updates, deletions, and suspensions can
occur at any time.

While not every service will expose account verifications or suspensions,
Skyfall will offer the features.

## Proposal

There are standard fields for an account such as name, hashed password, and
timestamps to persist creation time, and last updated. These will be present on
the `Account` protobuf message. A `properties` field will also be present to
capture all additional data.

The following will describe all APIs and incoming and outgoing protobuf
messages.

### CreateAccount

The minimum required fields to create an account are a name, password, and a
password confirmation. All other data will be stored in properties. The name can
be a minimum of 3 alpha-numeric characters or an email address, depending on the
configuration of the service. Passwords must be a minimum of 8 characters, and
two of the following:
- Uppercase and lowercase characters
- At least one number
- At least one non-alpha-numeric character

```
message CreateAccountRequest {
  string name = 1;
  string password = 2;
  string password_confirmation = 3;
  google.protobuf.Struct properties = 4;
}

message CreateAccountResponse {
  Account account = 1;
}

message Account {
  string id = 1;
  string name = 2;
  string password = 3;
  string password_confirmation = 4;
  google.protobuf.Struct properties = 5;
  google.protobuf.Timestamp created_at = 6;
  google.protobuf.Timestamp updated_at = 7;
  google.protobuf.Timestamp verified_at = 8;
  google.protobuf.Timestamp suspended_at = 9;
}
```

### UpdateAccount

The update operation will accept name and properties fields. Both fields are
optional, but partial updates to properties are not supported. 

```
message UpdateAccountRequest {
  string name = 1;
  google.protobuf.Struct properties = 2;
}

message UpdateAccountResponse {
  Account account = 1;
}
```

### DeleteAccount

The delete operation is a "soft delete" to later support undeleting an account.
The `DeleteAccountResponse` is empty, but intentionally not a
`google.protobuf.Empty` message in case data needs to be returned in the future.

```
message DeleteAccountRequest {
  string id = 1;
}

message DeleteAccountResponse {}
```

### VerifyAccount

The verify operation is meant to support scanerios where a service requires
verifying the account before being able to use said service. All operations
described in this proposal will be allowed for unverified accounts.

```
message VerifyAccountRequest {
  string id = 1;
}

message VerifyAccountResponse {
  Account account = 1;
}
```

### SuspendAccount

Suspending an account renders the account unusable. This differs from the delete
operation because this does not free the account name.

```
message SuspendAccountRequest {
  string id = 1;
}

message SuspendAccountResponse {
  Account account = 1;
}
```

### GetAccount and GetAccountByName

To align with many operations that target an account by ID, the `GetAccount`
operation accepts an ID string. In some cases only an account name is known. For
this scanerio the `GetAccountByName` operation can be used. These two operations
could be combined with the use of `oneof`, but after reviewing the generated
code of both, two separate operations seemed the most straightforward.

```
message GetAccountRequest {
  string id = 1;
}

message GetAccountResponse {
  Account account = 1;
}


message GetAccountByNameRequest {
  string name = 1;
}

message GetAccountByNameResponse {
  Account account = 1;
}

```

## Out of Scope

- AuthN/AuthZ - Auth related operations will be defined in future proposals
