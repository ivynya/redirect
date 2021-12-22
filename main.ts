
import { Application } from "./deps.ts";
import { queryDatabase } from "./notion/queryDatabase.ts";

const app = new Application();

app.use(async ctx => {
	const name = ctx.request.url.pathname;
	console.log(`[LOG] Parsed ${name}`);
	console.log(await queryDatabase());
	ctx.response.status = 302;
	ctx.response.headers.set("Location", `https://sdbagel.com${name}`);
});

console.log("[EVT] Listening on http://localhost:8000");
await app.listen({ port: 8000 });