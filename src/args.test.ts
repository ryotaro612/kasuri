import { describe, expect, test } from "vitest";
import { parseCmdLine } from "./args.ts";

describe("Command line arguments", async () => {
  test("The environment variable ASHITABA_OIDC_CLIENT_ID is set to OIDC Client ID", async () => {
    const res = parseCmdLine(["--env", "src/env"]);
    expect(res.oidcClientId).toBe("cid");
  });
  test("The environment variable ASHITABA_OIDC_CLIENT_SECRET is set to OIDC Client Secret", async () => {
    const res = parseCmdLine(["--env", "src/env"]);
    expect(res.clientSecret).toBe("sec");
  });

  test("--help is available.", async () => {
    const res = parseCmdLine(["--help"]);
    expect(res.help).toBeTruthy();
  });
});
