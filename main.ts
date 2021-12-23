
import { Application } from "./deps.ts";
import { queryDatabase } from "./notion/queryDatabase.ts";

const app = new Application();

app.use(async ctx => {
	const name = ctx.request.url.pathname;

	const validRedirects = await queryDatabase();
	const redir = validRedirects.find(r =>
		r.Short === name || r.Short === name.slice(1));

	if (redir) {
		ctx.response.status = 302;
		ctx.response.headers.set("Location", redir.RedirectURL);
		console.log(`[LOG] Parsed ${name} > ${redir.RedirectURL}`);
	} else {
		ctx.response.status = 400;
		ctx.response.body = "The redirect you requested does not exist.";
		console.log(`[LOG] Invalid ${name} > 404`);
		return;
	}

	if (redir.CampaignID) {
		// Make a request to the companion Analytics API
		const apiBase = Deno.env.get("ANALYTICS_API_HOST");
		await fetch(`https://${apiBase}/v1/campaign/${redir.CampaignID}`,
			{ method: "POST" });
		console.log(`[LOG] Analytics request for ${redir.RedirectURL}:${redir.CampaignID}`);
	}
});

console.log("[EVT] Listening on http://localhost:8000");
await app.listen({ port: 8000 });