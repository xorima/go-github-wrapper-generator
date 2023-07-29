# go-github-wrapper-generator

A CLI tool to help generate an up to date wrapper for go-github.

## Why

Today it is hard to create interfaces for `go-github` due to how it is designed using fields instead of methods for accessing the various underlying parts of github.
So instead of getting a mock client that only needs to implement the `client.Apps.ListRepo` call you need to mock every call, or do more complex changes.

There are mocking libraries which can help with this, but sometimes you want more control than that, this generator will give you the option for that greater control.
