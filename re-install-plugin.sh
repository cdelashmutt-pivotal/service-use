#!/bin/sh
cf uninstall-plugin ServiceUsePlugin
cf install-plugin -f service-use
