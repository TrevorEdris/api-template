package app

const (
    // AppName is the name of this application.
    AppName = "api-template"

    // EnvPrefix is used to narrow the scope of environment variables being parsed to
    // only those that start with this prefix.
    //
    // Example:
    //    1. API_TEMPLATE_SOME_CONFIG_VALUE=true
    //    2. SOME_OTHER_CONFIG_VALUE=false
    //
    // Example 1 _would_ be parsed by this application, as it begins with the below `EnvPrefix`.
    // Example 2 _would not_ be parsed, as it does not begin with the `EnvPrefix`.
    EnvPrefix = "API_TEMPLATE_"
)
