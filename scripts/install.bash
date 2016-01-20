#!/bin/bash

go build; cf uninstall-plugin appenvs; cf install-plugin appenvs
