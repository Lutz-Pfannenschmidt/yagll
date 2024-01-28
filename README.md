# YAGLL - Yet Another Go Logging Libary
My personal take on a logging libary, primarily made for [ResponsePlan][0],
inspired by [Distillog][1].

## Features:
- [Adjustable log output](#adjustable-log-output)
- [Toggle every log level on and off](#toggle-every-log-level-on-and-off)
- [Adjustable colors](#adjustable-colors)
- [Adjustable template using text/template](#adjustable-template-using-texttemplate)
- [Default logger](#default-logger)

## Adjustable log output
The log output can be adjusted using the `SetOutput` function.
Any `os.File` can be used as an output.

## Toggle every log level on and off
Every log level can be toggled on and off using the `Toggle` function.

## Adjustable colors
The colors used for the log levels can be adjusted using the `SetColor` function.

## Adjustable template using text/template
The template used to format the log messages is fully adjustable using
`text/template`. The following variables are available:
- `{{ .Time }}` - The time the log message was created
- `{{ .Level }}` - The log level of the message
- `{{ .Message }}` - The message itself
- `{{ .Color }}` - The color of the log level
- `{{ .Reset }}` - The reset code to reset the color

## Default logger
A default logger is provided, which can be used to log messages without
creating a new logger. The default logger can be accessed at package level.

[0]: https://github.com/Lutz-Pfannenschmidt/ResponsePlan
[1]: https://github.com/amoghe/distillog