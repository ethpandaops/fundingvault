#!/bin/bash

cd "$(dirname "$0")"
yarn install
npm run build
