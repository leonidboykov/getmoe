Boards section describes boards connections for `getmoe` to work with.

``` yaml
boards:
  # Use any name you want to, you will be able to use board names in downloads section
  <board_name>:
    provider: <provider_name>
    url: <board_url>
    password_salt:
    appkey_salt:
    posts_limit:
```

## Boards

### `provider`

Specifies booru API provider. Following providers are supported

  * `danbooru`
  * `gelbooru`
  * `moebooru`
  * `sankaku` (provides access to https://idol.sankakucomplex.com)
  * `sankaku_v2` (provides access to https://chan.sankakucomplex.com)

!!! note
    You may check [`provider`](https://github.com/leonidboykov/getmoe/tree/master/provider) package for all supported providers.

### `url`

URL of the board.

### `password_salt`

Password salt is used to hash your password before sending it to server. Use `%s` as placeholder for out password.

### `posts_limit`

This optional parameter allows to specify a maximum number of posts per request. By default all providers will try to set the largest possible value, but you may want to set it manually if your board allows a bigger value.

!!! note
    You may hit the maximum page limit if your `posts_limit` is too small and you have too generic tags for search.

### `credentials`

Some boards provides access to NSFW tags or allows to use more tags in search only for logged in users. You may add your credentials to board configuration.

``` yaml
login: username
password: super_secure
```

!!! note
    If you worry about storing your passwords as plain text in cofiguration files, you can hash your password manually and provide `hashed_password` setting instead of `password`.

For `gelbooru` sites you have to provide your user ID and API key **instead** of login and password:

``` yaml
user_id: 1337
apikey: your_gelbooru_apikey
```
