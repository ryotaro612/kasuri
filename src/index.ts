#!/usr/bin/env -S node --enable-source-maps

import { argv } from "node:process";
import { main } from "./main.ts";

main(argv.slice(2)).then((exitCode) => {
  if (exitCode !== 0) {
    process.exit(exitCode);
  }
});
