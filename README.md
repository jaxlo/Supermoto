# Supermoto
A guide for building fullstack websites in Go

- [ ] Build Docker config and deploy
- [ ] Home page and introduction (Link to the thesis)
- [ ] Quickstart
- [ ] Cookbook? Reference page?
- [ ] Database (Implement this for comments. Once you fully figure out migrations)
- [ ] Hosting guide

I like the ide of a reference page. Add links within the page

Local testing:
docker compose -f compose.yaml -f compose.dev.yaml up

Prod:
docker compose -f compose.yaml -f compose.prod.yaml up -d --build
-d runs it in detached mode
--build rebuilds the entire thing

