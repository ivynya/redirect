
import { Application } from "./deps.ts";
import { queryDatabase } from "./notion/queryDatabase.ts";

const app = new Application();

app.use(async ctx => {
	const name = ctx.request.url.pathname;

	const validRedirects = await queryDatabase();
	const redir = validRedirects.find(r =>
		r.Short[0].content === name || r.Short[0].content === name.slice(1));

	if (redir) {
		ctx.response.status = 302;
		ctx.response.headers.set("Location", redir.RedirectURL.url);
		console.log(`[LOG] Parsed ${name} > ${redir.RedirectURL.url}`);
	} else {
		ctx.response.status = 400;
		ctx.response.body = "The redirect you requested does not exist.";
		console.log(`[LOG] Invalid ${name} > 404`);
	}

	if (redir.CampaignID) {
		// Make a request to the companion Analytics API
	}
});

console.log("[EVT] Listening on http://localhost:8000");
await app.listen({ port: 8000 });