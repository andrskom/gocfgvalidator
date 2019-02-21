# gocfgvalidator

Simple pkg for recursive validation of config used reflection for it.

## Configuring 

Use option func `MustWithDeepOfRecursion(deep int)` if you want to change deep of validating.
By default set 7.

Use option func `WithStrictMode(enable bool)` for disable or enable strict mode.
By default strict mode is enabled.

### Strict mode

If enabled, u must implement interface `Validator` for all struct in your config.

## Development

Use make fail for linting and testing.

 