
import { Application } from "./deps.ts";

const app = new Application();

app.use(async ctx => {
	const name = ctx.request.url.pathname;
	console.log(`[LOG] Parsed ${name}`);
	ctx.response.status = 302;
	ctx.response.headers.set("Location", `https://sdbagel.com${name}`);
});

console.log("[EVT] Listening on http://localhost:8000");
await app.listen({ port: 8000 });