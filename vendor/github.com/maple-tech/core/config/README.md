# Core - Configuration

Globally handles loading and keeping the configuration options for the application. Contains pre-set options for database, logging, session, and web. Will load configuration files based on a given deployment target (environment). 

The deployment targets are the environment the service is intending to be run in. The options include `development`, `beta`, and `production`. Development is intended for local, and only local development testing. The beta target is intended for new feature, limited release versions. And production is the final stable version running for most all users.

At this time, the config package does not track or facilitate defaults. So all options need to be provided or else the "zero" value for that type will be used. I do not really intend to have a defaults functionality since that can get a bit confusing in-between versions.

## Usage

Like most packages, this needs to be initialized. And should be loaded **first** in the application lifecycle. Since most every other package depends on this. Using the function `config.Load()` with the deployment target, and the config path prefix, will load the configuration options from file. 

Configuration files use a prefix method instead of an actually declared file path. This means depending on which deployment environment you choose, the suffix and extension will automatically be added to the path. For instance per deployment target:

- **Development**: will append `.dev.json` to the config path. Turning the option `./config` into `./config.dev.json`
- **Beta**: will append `.beta.json` to the config path. Turning the option `./config` into `./config.beta.json`
- **Production**: will append `.json` to the config path. Turning the option `./config` into `./config.json`

Any package can retrieve the current config options using `config.Get()` with will return a pointer to the global config options that where applied. **Do not modify values in this pointer!**

You can also use `config.IsDevelopment()` to check if it was loaded as a development environment.