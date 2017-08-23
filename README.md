logger
---------

[![Build Status](https://travis-ci.org/corpix/logger.svg?branch=master)](https://travis-ci.org/corpix/logger)

Simple logging interface and implementations for various loggers.

## Supported log targets

- `logrus`, wrapper for [sirupsen/logrus](github.com/sirupsen/logrus) which implements `Logger`
- `nsq`, logs to Nsq with fallback to other `Logger`
- `prefixwrapper`, wraps any `Logger` interface messages with prefix
