import { parseCmdLine } from "./args.ts";
import { signIn } from "./signin.ts";

export async function main(args: string[]): Promise<number> {
  const command = parseCmdLine(args);
  if (command.error) {
    console.error(command.error);
    return 1;
  }
  await signIn({
    clientId: command.oidcClientId,
    clientSecret: command.clientSecret,
    port: command.port,
  });
  return 0;
}
