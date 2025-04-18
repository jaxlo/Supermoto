# Supermoto
A guide for building fullstack websites in Go.

- [x] Build Docker config and deploy
- [x] Home page and introduction (Link to the thesis)
- [x] Quickstart/Hello World
- [ ] A basic set of sections for fullstack web dev
- [ ] SEO
- [ ] QOL features. Sticky nav, click to copy, Ect.
- [ ] Database (Implement this for comments)
- [ ] Hosting guide


## Running in dev and prod
Local testing:
``docker compose -f compose.yaml -f compose.dev.yaml up``

Prod:
``docker compose -f compose.yaml -f compose.prod.yaml up -d --build``

- `-d` runs it in detached mode
- `--build` rebuilds the entire thing


## Contributing
This is likely a project I will build out very slowly over time.
Feel free to open a PR with edits or new sections.
Just try to use the standard library as much as reasonably possible.


## License
This code is in the public domain.
You can use it for any purpose and there is no need for attribution.
No warranty is provided.
