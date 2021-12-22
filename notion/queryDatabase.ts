
import { flattenResult } from "../deps.ts";

let lastUpdated = new Date(0);

export async function queryDatabase(): Promise<any[]> {
	if (lastUpdated.getTime() + 10000 > Date.now())
		return JSON.parse(await Deno.readTextFile("cache.json"));
	else lastUpdated = new Date();
	
	const id = Deno.env.get("NOTION_DB_ID");
	const res = await fetch(`https://api.notion.com/v1/databases/${id}/query`, {
		method: "POST",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Notion-Version": "2021-08-16"
		}
	});
	return flattenResult(await res.json());
}