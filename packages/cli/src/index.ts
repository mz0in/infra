#!/usr/bin/env node

import * as chalk from 'chalk'
import * as boxen from 'boxen'
import * as updateNotifier from 'update-notifier'

import { apiKey } from './api'
import { program } from './commands/index'

import * as packageJSON from '../package.json'

if (!apiKey) {
  const errorBox = boxen.default(
    `Cannot find env var ${chalk.default.bold(
      'DEVBOOK_KEY',
    )}\n\nVisit https://dash.usedevbook.com/settings to get your API key then run this CLI with the env var set.`,
    {
      width: 70,
      float: 'center',
      padding: 0.5,
      margin: 1,
      borderStyle: 'round',
      borderColor: 'redBright',
    },
  )
  console.error(errorBox)
}

updateNotifier.default({ pkg: packageJSON }).notify()

program
  .version(packageJSON.version, undefined, 'Displays the version of Devbook CLI')
  .parse()