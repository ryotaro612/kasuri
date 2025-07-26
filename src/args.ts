import { existsSync, readFileSync } from "node:fs";
import { type ParseArgsOptionsType, parseArgs, parseEnv } from "node:util";

/**
 *
 */
interface Command {
  port: number;
  oidcClientId: string;
  clientSecret: string;
  error: string | null;
  help: boolean;
}

/**
 *
 * @param args
 * @returns
 */
export function parseCmdLine(args: string[]): Command {
  const options = {
    port: {
      type: "string" as ParseArgsOptionsType,
      short: "p",
      default: "36332",
    },
    env: {
      type: "string" as ParseArgsOptionsType,
      short: "e",
      default: ".env",
    },
    help: {
      type: "boolean" as ParseArgsOptionsType,
      short: "h",
      default: false,
    },
  };

  const parsed = parseArgs({ args, options });
  const envPath = parsed.values.env as string;
  let env: NodeJS.Dict<string> = {};
  if (existsSync(envPath)) {
    env = readEnvFile(envPath);
  }

  let error = null;
  const oidcClientId =
    process.env.ASHITABA_OIDC_CLIENT_ID ?? env.ASHITABA_OIDC_CLIENT_ID;
  if (!oidcClientId) {
    error = "Environment variable ASHITABA_OIDC_CLIENT_ID must be defined.";
  }
  const oidClientSecret =
    process.env.ASHITABA_OIDC_CLIENT_SECRET ?? env.ASHITABA_OIDC_CLIENT_SECRET;
  if (!oidClientSecret) {
    error = "Environment variable ASHITABA_OIDC_CLIENT_SECRET must be defined.";
  }

  return {
    port: parseInt(parsed.values.port as string),
    oidcClientId: oidcClientId ?? "",
    clientSecret: oidClientSecret ?? "",
    error,
    help: parsed.values.help as boolean,
  };
}

/**
 *
 */
export function helpMessage(): string {
  return `
	Usage: ashitaba [options]

	Options:
	  -p, --port <port>          Port to listen on (default: 36332)
	  -e, --env <env>            Path to the environment file (default: .env)
	  -h, --help                 Show this help message
	`;
}

function readEnvFile(envFilePath: string): NodeJS.Dict<string> {
  const buffer = readFileSync(envFilePath);
  return parseEnv(buffer.toString());
}
