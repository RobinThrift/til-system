# TIL System

[![Build Status](https://travis-ci.org/RobinThrift/til-system.svg?branch=master)](https://travis-ci.org/RobinThrift/til-system)

# Dependencies

All dependencies are committed in the `vendor` folder, so this is `go get`-able, and you don't need any extra tool to develop it,
however, it uses [gvt](https://github.com/FiloSottile/gvt) to manage and update it's dependencies, so it is advised to use it when
dealing with dependencies (but feel free to add/update the entry in the `vendor/manifest` file by hand, if you're so inclined).

To install `gvt` simply run: `go get -u github.com/FiloSottile/gvt`


# Asset generation

The server uses [go-bindata](https://github.com/jteeuwen/go-bindata) to embed the asset files inside of the binary. Run `make assets` to run
the code generation for embedding. Make sure to commit the generated file.

To install `go-bindata` simply run: `go get -u github.com/jteeuwen/go-bindata/...`


# Running

There are a few environment variables than can be used to configure the server:

- `TIL_PORT`: sets the port to listen on (default: `3000`)
- `TIL_SECRET`: sets the secret that needs to be provided by the clients (**required**)
- `TIL_REPO_URL`: specify a url or path to the git repository which will be cloned and pushed (**required**)
- `TIL_POST_DIR`: path inside the repo where the post files will be created (default: `content/til`)


# Posting Data

To create a new TIL entry `POST` data to `/add`:

```json
{
    "posted_date": UNIX_TIMESTAMP,
    "posted_from": SOME_KIND_OF_UA,
    "content": CONTENT_STRING
}
```

If the post was successful you will receive the same data back. Otherwise you will receive an error JSON response (and corresponding HTTP status code):

```json
{
    "code": CODE,
    "message": STRING
}
```
