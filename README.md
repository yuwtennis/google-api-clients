# google-api-clients

## Motivation
I use Google Client SDK in my job, however, sometimes I need concrete examples because [doc](https://developers.google.com/api-client-library) is not enough.  
This repository is to contribute to people like me.

## Contents

* [Available examples](#available-examples)
* [Tutorials](#tutorials)
  * [How to run examples](#how-to-run-examples)
  * [How to run tests](#how-to-run-tests)
* [Available environment variables](#available-environment-variables)

## Available examples

| Api | go  | python |
| --- |-----| ------ |
| admin/report | available | available |

## Tutorials

### How to run examples

Set [environment variables](#available-environment-variables) accordingly before running the example.

#### go

```markdown
cd go/
go run examples/PATH-TO-MAIN/main.go
```

#### python

```markdown
cd python/
pip3 intall 
```

### How to run tests

#### go

```markdown
cd go/
make unittest
```

#### python 

```markdown
cd python/
make unittest
```

## Available environment variables

### admin/report

| variable | description                                                                                                                                                               | default |
| -------- |---------------------------------------------------------------------------------------------------------------------------------------------------------------------------| ------- |
| GOOGLE_DEFAULT_CREDENTIALS | See [doc](https://cloud.google.com/docs/authentication/production#passing_variable)                                                                                       | no default |
| GOOGLE_SUBJECT_EMAIL | User email to impersonate . Required for [Google Workspace Domain-Wide Delegation of Authority](https://developers.google.com/admin-sdk/directory/v1/guides/delegation) . |