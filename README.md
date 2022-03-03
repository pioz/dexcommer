# Dexcommer

Read last glucose values using Dexcom Follow credentials from command line.

## Requirements

Go v1.17 or greater.

## Installation

```bash
go get github.com/pioz/dexcommer
```

## Usage

```bash
$ DEXCOMMER_USERNAME=enrico DEXCOMMER_PASSWORD=passw0rd dexcommer
```

You can set the env variables in a `.env` file

Available environment variables are:

- DEXCOMMER_USERNAME
- DEXCOMMER_PASSWORD
- DEXCOMMER_APPLICATION_ID

## Questions or problems?

If you have any issues please add an [issue on
GitHub](https://github.com/pioz/dexcommer/issues) or fork the project and send a
pull request.

## Copyright

Copyright (c) 2020 [Enrico Pilotto (@pioz)](https://github.com/pioz). See
[LICENSE](https://github.com/pioz/dexcommer/blob/master/LICENSE) for details.
