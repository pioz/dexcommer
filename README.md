# Dexcommer

Read lastest glucose values using Dexcom Follow credentials from command line (it's written in Go!).

## Installation

```bash
go get github.com/pioz/dexcommer
```

## Usage

```bash
$ DEXCOMMER_USERNAME=enrico DEXCOMMER_PASSWORD=passw0rd DEXCOMMER_APPLICATION_ID=1234 dexcommer
```

You can set the env variables in a `.env` file

Available environment variables are:

- DEXCOMMER_USERNAME
- DEXCOMMER_PASSWORD
- DEXCOMMER_APPLICATION_ID

Run `dexcommer -h` to display cli options.

## Lib

You can also read the latest glucose values from your Go program, here the [godoc](https://godoc.org/github.com/pioz/dexcommer).

## Questions or problems?

If you have any issues please add an [issue on
GitHub](https://github.com/pioz/dexcommer/issues) or fork the project and send a
pull request.

## Copyright

Copyright (c) 2020 [Enrico Pilotto (@pioz)](https://github.com/pioz). See
[LICENSE](https://github.com/pioz/dexcommer/blob/master/LICENSE) for details.
