import {
  createServer,
  type IncomingMessage,
  type ServerResponse,
} from "node:http";
import { OidcClient } from "oidc-client-ts";
import open from "open";

/**
 *
 */
export async function signIn({
  clientId,
  clientSecret,
  port,
}: {
  clientId: string;
  clientSecret: string;
  port: number;
}) {
  const client = new OidcClient({
    authority: "https://accounts.google.com",
    client_id: clientId,
    client_secret: clientSecret,
    redirect_uri: `http://127.0.0.1:${port}/callback`,
  });

  const request = await client.createSigninRequest({});
  await open(request.url);

  const res: { data: undefined | string } = { data: undefined };
  const server = runServer(port, res);

  while (!res.data) {
    await sleep(500);
  }
  server.close();
  console.log(res.data);
  const re = await client.processSigninResponse(res.data);
  console.log(re);
}

function runServer(port: number, result: { data: undefined | string }) {
  const server = createServer(
    async (req: IncomingMessage, res: ServerResponse) => {
      console.log("request: ", req.url);
      //res.writeHead(302, { Location: req.url });
      res.statusCode = 200;
      result.data = `http://127.0.0.1:${port}${req.url}`;
      // res.write('');
      res.end();
    }
  );
  return server.listen(port, () => {
    // callback
  });
}

function sleep(ms: number) {
  return new Promise((res) => setTimeout(res, ms));
}
